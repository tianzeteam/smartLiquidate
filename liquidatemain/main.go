package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"math/big"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"swap/config"
	"time"

	"github.com/go-co-op/gocron"

	_ "github.com/go-sql-driver/mysql"
	"github.com/xormplus/xorm"
	"gopkg.in/yaml.v2"

	"swap/liquidatecontract"
	models "swap/models"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"k8s.io/klog"
)

var engine *xorm.Engine

func task(engine *xorm.Engine, config *config.Config, client *ethclient.Client) {
	log.Println("--------------清算队列数据整理开始-------------")
	lendpoolContract, _ := liquidatecontract.NewLendPool(common.HexToAddress(config.Contract.LendpoolContract), client)

	liquidateUser := new(models.LiquidateUser)
	users, err := engine.Where("id >?", 1).Rows(liquidateUser)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer users.Close()
	for users.Next() {
		err = users.Scan(liquidateUser)
		userId := liquidateUser.UserId

		if err != nil {
			klog.Error(err)
		}

		HeathFactorData, err := lendpoolContract.GetUserAccountData(nil, common.HexToAddress(userId))
		if err != nil {
			klog.Error(err)
			continue
		}

		userHealthFactor := big.NewFloat(0).SetInt(HeathFactorData.HealthFactor)
		userTotalDebtETH := big.NewFloat(0).SetInt(HeathFactorData.TotalDebtETH)
		decimals := big.NewFloat(math.Pow(10, float64(18))) //精度
		userStandardHealthFactor := new(big.Float).Quo(userHealthFactor, decimals)
		userStandardTotalDebtEth := new(big.Float).Quo(userTotalDebtETH, decimals)

		MAX_HEALTH_THRESHOLD := float64(1.5)
		MAX_DEBT_THRESHOLD := float64(0.1)
		currentUserHealthValue, _ := userStandardHealthFactor.Float64()
		currentUserTotalDebtEthFloat, _ := userStandardTotalDebtEth.Float64()
		//fmt.Println("debtThreshold:::", MAX_DEBT_THRESHOLD)
		//fmt.Println("totalDebtEthFloat:::", currentUserTotalDebtEthFloat)
		//fmt.Println("当前用户健康值:::", currentUserHealthValue)
		if currentUserHealthValue < MAX_HEALTH_THRESHOLD && currentUserTotalDebtEthFloat > MAX_DEBT_THRESHOLD {

			currUserCollateral := make([]models.AaveAsset, 0)
			err := engine.Where("user_id = ?", userId).And("asset_type = ?", "collateral").Find(&currUserCollateral)
			if err != nil {
				klog.Error(err)
			}
			if len(currUserCollateral) == 0 {
				klog.Error("当前用户抵押信息不完整:::::::", userId)
				continue
			}
			currUserBorrow := make([]models.AaveAsset, 0)

			error := engine.Where("user_id = ?", userId).And("asset_type = ?", "borrow").Find(&currUserBorrow)
			if error != nil {
				klog.Error(error)

			}
			if len(currUserBorrow) == 0 {
				klog.Error("当前用户借款信息不完整:::::::", userId)
				continue
			}

			currLiquidateEntry := make([]models.LiquidateQueue, 0)

			errs := engine.Where("user_id = ?", userId).Find(&currLiquidateEntry)
			if errs != nil {
				klog.Error(errs)

			}
			if len(currLiquidateEntry) == 0 {
				liquidateQueue := new(models.LiquidateQueue)
				liquidateQueue.UserId = userId
				liquidateQueue.BorrowAmount = currUserBorrow[0].Amount
				liquidateQueue.BorrowAsset = currUserBorrow[0].UnderlyingAsset
				liquidateQueue.CollateralAsset = currUserCollateral[0].UnderlyingAsset
				liquidateQueue.CreateTime = time.Now().Unix()
				liquidateQueue.Status = "waiting"
				engine.InsertOne(liquidateQueue)
			}

		}

	}
	log.Println("--------------清算队列数据整理结束-------------")
}

func main() {
	config := InitConfig()
	engine, _ := xorm.NewEngine("mysql", config.App.DatabaseUrl)
	client, err := ethclient.Dial(config.Account.BscMainNetHttpUrl)
	if err != nil {
		klog.Fatal(err)
	}
	goScheduler := gocron.NewScheduler(time.UTC) // 使用UTC时区

	goScheduler.Every(120).Seconds().WaitForSchedule().Do(task, engine, config, client)
	goScheduler.StartAsync()

	opts := &bind.CallOpts{
		From:    common.HexToAddress(config.Account.AccountAddr),
		Context: context.Background(),
	}

	lendpoolContract, _ := liquidatecontract.NewLendPool(common.HexToAddress(config.Contract.LendpoolContract), client)
	liquidateAndLoanContract, err := liquidatecontract.NewLiquidateLoan(common.HexToAddress(config.Contract.LiquidateLoanContract), client)
	if err != nil {
		klog.Fatal(err)
	}
	liquidateQueue := new(models.LiquidateQueue)

	for {

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
			userTotalDebtETH := big.NewFloat(0).SetInt(HeathFactorData.TotalDebtETH)
			decimals := big.NewFloat(math.Pow(10, float64(18))) //精度
			userStandardHealthFactor := new(big.Float).Quo(userHealthFactor, decimals)
			userStandardTotalDebtEth := new(big.Float).Quo(userTotalDebtETH, decimals)
			//fmt.Println("用户债务健康值:::",standardHealthFactor)
			MAX_HEALTH_THRESHOLD := float64(1)
			MAX_DEBT_THRESHOLD := float64(0.1)
			currentUserHealthValue, _ := userStandardHealthFactor.Float64()
			currentUserTotalDebtEthFloat, _ := userStandardTotalDebtEth.Float64()
			fmt.Println("debtThreshold:::", MAX_DEBT_THRESHOLD)
			fmt.Println("totalDebtEthFloat:::", currentUserTotalDebtEthFloat)
			fmt.Println("当前用户健康值:::", currentUserHealthValue)
			if currentUserHealthValue < MAX_HEALTH_THRESHOLD && currentUserTotalDebtEthFloat > MAX_DEBT_THRESHOLD {
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

				flashLoansV2(liquidateAndLoanContract, client, config.Account.AccountPriKey, borrowAsset, flashLoanAmount, collateralAsset, liquidateAddress, amountOutMin, swapPath)

			}

		}

		time.Sleep(time.Duration(5) * time.Second)
	}
}

func flashLoansV2(liquidateAndLoanContract *liquidatecontract.LiquidateLoan, client *ethclient.Client, priKey string, _assetToLiquidate common.Address, _flashAmt *big.Int, _collateral common.Address, _userToLiquidate common.Address, _amountOutMin *big.Int, _swapPath []common.Address) string {
	return "......"
}

func flashLoans(liquidateAndLoanContract *liquidatecontract.LiquidateLoan, client *ethclient.Client, priKey string, _assetToLiquidate common.Address, _flashAmt *big.Int, _collateral common.Address, _userToLiquidate common.Address, _amountOutMin *big.Int, _swapPath []common.Address) string {

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		klog.Error(err)
	}
	privateKey, err := crypto.HexToECDSA(priKey)
	if err != nil {
		klog.Error(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	//
	if !ok {
		klog.Error("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		klog.Error(err)
	}

	value := big.NewInt(0)
	var gasPriceFloat big.Float
	gasPriceFloat.SetString("23") //200 GWei

	tenDecimal := big.NewFloat(math.Pow(10, float64(9)))
	gasPrice, _ := new(big.Float).Mul(tenDecimal, &gasPriceFloat).Int(&big.Int{})

	gasLimit := uint64(3000101)

	var non big.Int
	non.SetUint64(nonce)
	auth, _ := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	auth.Nonce = &non
	auth.Value = value
	auth.GasLimit = gasLimit
	auth.GasPrice = gasPrice

	tx, err := liquidateAndLoanContract.ExecuteFlashLoans(auth, _assetToLiquidate, _flashAmt, _collateral, _userToLiquidate, _amountOutMin, _swapPath)
	if err != nil {
		klog.Error(err)
	}
	klog.Info("liquidate tx: https://kovan.etherscan.io/tx/", tx.Hash())
	return tx.Hash().Hex()
}

func GetAppPath() string {
	host_name := os.Getenv("HOST_NAME")
	fmt.Println(host_name)
	if host_name == "" {
		return ".."
	}
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	index := strings.LastIndex(path, string(os.PathSeparator))

	return path[:index]
}

func InitConfig() *config.Config {

	var _config *config.Config

	yamlFile, err := ioutil.ReadFile(GetAppPath() + "/config/config.yml")
	if err != nil {
		fmt.Println(err.Error())
	}

	err = yaml.Unmarshal(yamlFile, &_config)
	if err != nil {
		fmt.Println(err.Error())
	}
	return _config
}