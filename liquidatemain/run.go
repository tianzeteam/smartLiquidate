package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"math/big"
	"swap/config"
	"time"

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

func main() {
	config := InitConfig()
	client, err := ethclient.Dial(config.Account.BscMainNetHttpUrl)
	if err != nil {
		klog.Fatal(err)
	}

	opts := &bind.CallOpts{
		From:    common.HexToAddress(config.Account.AccountAddr),
		Context: context.Background(),
	}

	engine, _ := xorm.NewEngine("mysql", config.App.DatabaseUrl)

	lendpoolContract, _ := liquidatecontract.NewLendPool(common.HexToAddress(config.Contract.LendpoolContract), client)
	liquidateAndLoanContract, err := liquidatecontract.NewLiquidateLoan(common.HexToAddress(config.Contract.LiquidateLoanContract), client)
	if err != nil {
		klog.Fatal(err)
	}
	liquidateUser := new(models.LiquidateUser)

	for {

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

			HeathFactorData, err := lendpoolContract.GetUserAccountData(opts, common.HexToAddress(userId))
			if err != nil {
				klog.Error(err)
				continue
			}
			//klog.Info("HeathFactorData: ", &HeathFactorData)
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

				if err != nil {
					klog.Error(err)
				}
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
				collateralAsset := common.HexToAddress(currUserCollateral[0].UnderlyingAsset)
				borrowAsset := common.HexToAddress(currUserBorrow[0].UnderlyingAsset)
				flashLoanAmount, _ := new(big.Int).SetString(currUserBorrow[0].Amount, 10)
				flashLoanAmount = flashLoanAmount.Mul(flashLoanAmount, big.NewInt(1)).Div(flashLoanAmount, big.NewInt(2)) //* 0.5
				amountOutMin := big.NewInt(0)
				liquidateAddress := common.HexToAddress(userId)
				swapPath := []common.Address{
					collateralAsset, //用奖励抵押物资产 兑换 借贷资产
					borrowAsset,
				}

				flashLoans(liquidateAndLoanContract, client, config.Account.AccountPriKey, borrowAsset, flashLoanAmount, collateralAsset, liquidateAddress, amountOutMin, swapPath)

			}

		}

		time.Sleep(time.Duration(5) * time.Second)
	}
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

func InitConfig() *config.Config {

	var _config *config.Config

	yamlFile, err := ioutil.ReadFile("../config/config.yml")
	if err != nil {
		fmt.Println(err.Error())
	}

	err = yaml.Unmarshal(yamlFile, &_config)
	if err != nil {
		fmt.Println(err.Error())
	}
	return _config
}
