package core

import (
	"fmt"
	"math"
	"math/big"
	"swap/config"
	"swap/liquidateswap"
	"swap/utils"

	_ "github.com/go-sql-driver/mysql"
	"github.com/xormplus/xorm"

	"swap/liquidatecontract"
	models "swap/models"
	"swap/counter"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"k8s.io/klog"
)

func Process( _counter counter.Counter, engine *xorm.Engine, liquidateQueue *models.LiquidateQueue, liquidateAndLoanContract *liquidatecontract.LiquidateLoan, client *ethclient.Client, lendpoolContract *liquidatecontract.LendPool, opts *bind.CallOpts, config *config.Config, uniswapFactoryContract *liquidatecontract.Factory) {
	userId := liquidateQueue.UserId
	HeathFactorData, err := lendpoolContract.GetUserAccountData(opts, common.HexToAddress(userId))
	if err != nil {
		klog.Error(err)
		return
	}
	userHealthFactor := big.NewFloat(0).SetInt(HeathFactorData.HealthFactor)
	userTotalCollateralETH := big.NewFloat(0).SetInt(HeathFactorData.TotalCollateralETH)
	userTotalDebtETH := big.NewFloat(0).SetInt(HeathFactorData.TotalDebtETH)
	decimals := big.NewFloat(math.Pow(10, float64(18))) //精度
	userStandardHealthFactor := new(big.Float).Quo(userHealthFactor, decimals)
	userStandardTotalDebtEth := new(big.Float).Quo(userTotalDebtETH, decimals)
	userStandardTotalCollateralEth := new(big.Float).Quo(userTotalCollateralETH, decimals)
	MAX_HEALTH_THRESHOLD := config.Liquidate.MAX_HEALTH_THRESHOLD
	MAX_DEBT_THRESHOLD := config.Liquidate.MAX_DEBT_THRESHOLD
	MAX_COLLATERAL_THRESHOLD := config.Liquidate.MAX_COLLATERAL_THRESHOLD
	currentUserHealthValue, _ := userStandardHealthFactor.Float64()
	currentUserTotalDebtEthFloat, _ := userStandardTotalDebtEth.Float64()
	currentUserTotalCollateralEth, _ := userStandardTotalCollateralEth.Float64()

	fmt.Println("当前用户健康值:::", currentUserHealthValue)

	if currentUserHealthValue < MAX_HEALTH_THRESHOLD && currentUserTotalDebtEthFloat > MAX_DEBT_THRESHOLD && currentUserTotalCollateralEth > MAX_COLLATERAL_THRESHOLD {

		fmt.Println("用户id:::", userId)
		collateralAsset := common.HexToAddress(liquidateQueue.CollateralAsset)
		borrowAsset := common.HexToAddress(liquidateQueue.BorrowAsset)
		flashLoanAmount, _ := new(big.Int).SetString(liquidateQueue.BorrowAmount, 10)
		flashLoanAmount = flashLoanAmount.Mul(flashLoanAmount, big.NewInt(1)).Div(flashLoanAmount, big.NewInt(2)) //* 0.5
		amountOutMin := big.NewInt(0)
		liquidateAddress := common.HexToAddress(userId)

		swapPath := utils.GetBestTradeExactIn(collateralAsset, borrowAsset, big.NewInt(20000), uniswapFactoryContract, config)
		if swapPath == nil {
			return
		}
		klog.Infoln(" 对用户  "+userId+"  开始执行清算 queueID ", liquidateQueue.Id)
		liquidateswap.FlashLoans(_counter, engine, liquidateQueue, liquidateAndLoanContract, client, config.Account.AccountPriKey, borrowAsset, flashLoanAmount, collateralAsset, liquidateAddress, amountOutMin, swapPath, userId, config.Account.GasPrice, config.Account.GasLimit)

	}
}
