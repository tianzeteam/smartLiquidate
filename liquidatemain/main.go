package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"swap/config"
	"swap/core"
	"swap/counter"
	"swap/deploy"
	"swap/liquidatecontract"
	models "swap/models"
	"swap/scheduler"
	"swap/utils"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-co-op/gocron"
	_ "github.com/go-sql-driver/mysql"
	"github.com/xormplus/xorm"
	"gopkg.in/yaml.v2"
	"k8s.io/klog"
)

func main() {

	config := InitConfig()
	engine, _ := xorm.NewEngine("mysql", config.App.DatabaseUrl)
	taskEngine, _ := xorm.NewEngine("mysql", config.App.DatabaseUrl)
	//fetchEngine, _ := xorm.NewEngine("mysql", config.App.DatabaseUrl)
	client, err := ethclient.Dial(config.Account.RpcHttpUrl)

	if err != nil {
		klog.Fatal("访问rpc节点网络问题,清算进程退出")
		return
	}

	opts := &bind.CallOpts{
		From:    common.HexToAddress(config.Account.AccountAddr),
		Context: context.Background(),
	}

	lendpoolContract, _ := liquidatecontract.NewLendPool(common.HexToAddress(config.Contract.LendpoolContract), client)
	liquidateAndLoanContract, _ := liquidatecontract.NewLiquidatecontract(common.HexToAddress(config.Contract.LiquidateLoanContract), client)
	uniswapFactoryContract, err := liquidatecontract.NewFactory(common.HexToAddress(config.Contract.UniswapV2Factory), client)
	//priceOracle, _ := oraclecontract.NewOraclecontract(common.HexToAddress("0xA50ba011c48153De246E5192C8f9258A2ba79Ca9"), client)
	//helpcontract, _ := helpcontract.NewHelpcontract(common.HexToAddress("0x057835Ad21a177dbdd3090bB1CAE03EaCF78Fc6d"), client)
	if err != nil {
		klog.Fatal(err)
	}
	deploy.GetTokenIndex(liquidateAndLoanContract)
	//tokensMeta := deploy.GetAllTokens()
	goScheduler := gocron.NewScheduler(time.UTC) // 使用UTC时区
	//goScheduler.Every(20).Seconds().WaitForSchedule().Do(scheduler.UpdateAssetTask, fetchEngine, helpcontract, priceOracle, tokensMeta)
	goScheduler.Every(12000).Seconds().Do(scheduler.LiquidateEventTask, taskEngine, config, client)
	//goScheduler.Every(60).Seconds().WaitForSchedule().Do(scheduler.FetchTask, fetchEngine)
	//goScheduler.Every(60).Seconds().WaitForSchedule().Do(scheduler.Task, taskEngine, config, client)
	goScheduler.StartAsync()

	nonce, err := utils.GetNonce(client, config.Account.AccountPriKey)
	if err != nil {
		klog.Fatal("获取account nonce值错误,清算进程退出")
		return
	}
	counter := counter.NewMutexCounter(nonce)
	klog.Infoln(" counter ", counter.Read())

	asyn_channel := make(chan models.LiquidateQueue, 12000) // capacity size > rows

	for {

		Runing := config.App.Runing
		liquidateQueue := new(models.LiquidateQueue) // 实例化 LiquidateQueue struct ，指向pointer
		count, _ := engine.Where("id >?", 1).And("status = ?", "waiting").Count(liquidateQueue)
		var skim int64
		if count%int64(120) == 0 {
			skim = count / int64(120)
		} else {
			skim = count/int64(120) + 1
		}
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
	//fmt.Println("config path ::", _config)
	return _config
}
