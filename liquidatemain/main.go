package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"math/big"
	"swap/config"
	"swap/liquidateswap"
	"swap/scheduler"
	"swap/utils"
	"time"

	"github.com/go-co-op/gocron"

	_ "github.com/go-sql-driver/mysql"
	"github.com/xormplus/xorm"
	"gopkg.in/yaml.v2"

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
	nonce, err := liquidateswap.GetNonce(client, config.Account.AccountPriKey)
	if err != nil {
		//klog.Error(err)
		klog.Fatal("访问rpc节点网络问题,清算进程退出")
		return
	}

	for {
		liquidateQueue := new(models.LiquidateQueue)
		liquidateEntrys, err := engine.Where("id >?", 1).And("status = ?", "waiting").Rows(liquidateQueue)
		if err != nil {
			log.Fatal(err)
			return
		}
		defer liquidateEntrys.Close()
		for liquidateEntrys.Next() {
			err = liquidateEntrys.Scan(liquidateQueue)
			userId := liquidateQueue.UserId

			if err != nil {
				klog.Error(err)
			}

			HeathFactorData, err := lendpoolContract.GetUserAccountData(opts, common.HexToAddress(userId))
			if err != nil {
				klog.Error(err)
				continue
			}

			userHealthFactor := big.NewFloat(0).SetInt(HeathFactorData.HealthFactor)
			userTotalCollateralETH := big.NewFloat(0).SetInt(HeathFactorData.TotalCollateralETH)
			userTotalDebtETH := big.NewFloat(0).SetInt(HeathFactorData.TotalDebtETH)
			decimals := big.NewFloat(math.Pow(10, float64(18))) //精度
			userStandardHealthFactor := new(big.Float).Quo(userHealthFactor, decimals)
			userStandardTotalDebtEth := new(big.Float).Quo(userTotalDebtETH, decimals)
			userStandardTotalCollateralEth := new(big.Float).Quo(userTotalCollateralETH, decimals)
			//fmt.Println("用户债务健康值:::",standardHealthFactor)
			MAX_HEALTH_THRESHOLD := config.Liquidate.MAX_HEALTH_THRESHOLD
			MAX_DEBT_THRESHOLD := config.Liquidate.MAX_DEBT_THRESHOLD
			MAX_COLLATERAL_THRESHOLD := config.Liquidate.MAX_COLLATERAL_THRESHOLD
			currentUserHealthValue, _ := userStandardHealthFactor.Float64()
			currentUserTotalDebtEthFloat, _ := userStandardTotalDebtEth.Float64()
			currentUserTotalCollateralEth, _ := userStandardTotalCollateralEth.Float64()
			fmt.Println("TotalCollateralEthFloat:::", currentUserTotalCollateralEth)
			fmt.Println("totalDebtEthFloat:::", currentUserTotalDebtEthFloat)
			fmt.Println("当前用户健康值:::", currentUserHealthValue)
			utils.GetBestTradeExactIn(liquidateQueue.CollateralAsset,liquidateQueue.BorrowAsset,big.NewInt(20000),uniswapFactoryContract)
			if currentUserHealthValue < MAX_HEALTH_THRESHOLD && currentUserTotalDebtEthFloat > MAX_DEBT_THRESHOLD && currentUserTotalCollateralEth > MAX_COLLATERAL_THRESHOLD {
				fmt.Println("用户id:::", userId)

				collateralAsset := common.HexToAddress(liquidateQueue.CollateralAsset)
				borrowAsset := common.HexToAddress(liquidateQueue.BorrowAsset)
				flashLoanAmount, _ := new(big.Int).SetString(liquidateQueue.BorrowAmount, 10)
				flashLoanAmount = flashLoanAmount.Mul(flashLoanAmount, big.NewInt(1)).Div(flashLoanAmount, big.NewInt(2)) //* 0.5
				amountOutMin := big.NewInt(0)
				liquidateAddress := common.HexToAddress(userId)
				swapPath := []common.Address{
					collateralAsset, //用奖励抵押物资产 兑换 借贷资产
					borrowAsset,
				}
				pair, err := uniswapFactoryContract.GetPair(opts, collateralAsset, borrowAsset)
				if err != nil {
					klog.Fatalln("清算进程中获取货币对异常")
					continue
				}
				fmt.Println("pair---->", pair)
				if pair.Hex() != "0x0000000000000000000000000000000000000000" {
					TxHash := liquidateswap.FlashLoans(&nonce, engine, liquidateQueue, liquidateAndLoanContract, client, config.Account.AccountPriKey, borrowAsset, flashLoanAmount, collateralAsset, liquidateAddress, amountOutMin, swapPath)

					if TxHash != "fail" {
						liquidateResult := new(models.LiquidateResult)
						liquidateResult.UserId = userId
						liquidateResult.ReceiptAsset = liquidateQueue.BorrowAsset
						liquidateResult.TxHash = TxHash
						liquidateResult.CreateTime = time.Now().Unix()
						liquidateResult.Status = "pending"
						engine.InsertOne(liquidateResult)
					}
				}
			}

		}

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
