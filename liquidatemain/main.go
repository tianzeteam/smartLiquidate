package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"

	"swap/config"
	"swap/core"
	"swap/scheduler"
	"swap/utils"
	"time"

	"github.com/go-co-op/gocron"

	_ "github.com/go-sql-driver/mysql"
	"github.com/xormplus/xorm"
	"gopkg.in/yaml.v2"

	"swap/counter"
	"swap/liquidatecontract"
	models "swap/models"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"k8s.io/klog"
)

func main() {

	config := InitConfig()
	engine, _ := xorm.NewEngine("mysql", config.App.DatabaseUrl)
	taskEngine, _ := xorm.NewEngine("mysql", config.App.DatabaseUrl)
	client, err := ethclient.Dial(config.Account.RpcHttpUrl)

	if err != nil {
		klog.Fatal(err)

	}
	goScheduler := gocron.NewScheduler(time.UTC) // 使用UTC时区

	goScheduler.Every(60).Seconds().WaitForSchedule().Do(scheduler.Task, taskEngine, config, client)
	goScheduler.StartAsync()

	opts := &bind.CallOpts{
		From:    common.HexToAddress(config.Account.AccountAddr),
		Context: context.Background(),
	}

	lendpoolContract, _ := liquidatecontract.NewLendPool(common.HexToAddress(config.Contract.LendpoolContract), client)
	liquidateAndLoanContract, _ := liquidatecontract.NewLiquidateLoan(common.HexToAddress(config.Contract.LiquidateLoanContract), client)
	uniswapFactoryContract, err := liquidatecontract.NewFactory(common.HexToAddress(config.Contract.UniswapV2Factory), client)
	if err != nil {
		klog.Fatal(err)
	}
	nonce, err := utils.GetNonce(client, config.Account.AccountPriKey)
	counter := counter.NewMutexCounter(nonce)
	klog.Infoln(" counter ", counter.Read())
	if err != nil {
		//klog.Error(err)
		klog.Fatal("访问rpc节点网络问题,清算进程退出")
		return
	}

	asyn_channel := make(chan models.LiquidateQueue, 120) // capacity size > rows

	for {
		Runing := config.App.Runing
		liquidateQueue := new(models.LiquidateQueue) // 实例化 LiquidateQueue struct ，指向pointer
		count, _ := engine.Where("id >?", 1).And("status = ?", "waiting").Count(liquidateQueue)
		var skim int64
		if count % int64(120) == 0 {
			skim = count / int64(120)
		} else {
			skim = count / int64(120) + 1
		}
		fmt.Println(" count ", count)
		fmt.Println(" skim ", skim)
		for i := 0; i < int(skim); i++ {
			offset := i * 120
			liquidateEntrys, err := engine.Limit(120, offset).Where("id >?", 1).And("status = ?", "waiting").Rows(liquidateQueue)
			if err != nil {
				log.Fatal(err)
				return
			}
			defer liquidateEntrys.Close()
			for liquidateEntrys.Next() {
				err = liquidateEntrys.Scan(liquidateQueue)
				asyn_channel <- *liquidateQueue
				if err != nil {
					klog.Error(err)
				}
				// Goroutine

				if Runing {
					go func() {
						ConsumerQueue := <-asyn_channel
						core.Process(counter, engine, &ConsumerQueue, liquidateAndLoanContract, client, lendpoolContract, opts, config, uniswapFactoryContract)
					}()
				}

			}
		}
		klog.Info("进入 sleep 过程")

		time.Sleep(time.Duration(5) * time.Second)

	}
}

func InitConfig() *config.Config {

	var _config *config.Config

	yamlFile, err := ioutil.ReadFile(utils.GetAppPath() + "/config/config.yml")
	if err != nil {
		fmt.Println(err.Error())
	}

	err = yaml.Unmarshal(yamlFile, &_config)
	if err != nil {
		fmt.Println(err.Error())
	}
	return _config
}
