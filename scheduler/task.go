package scheduler

import (
	"log"
	"math"
	"math/big"
	"swap/config"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/xormplus/xorm"

	"swap/liquidatecontract"
	models "swap/models"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"k8s.io/klog"
)

func Task(engine *xorm.Engine, config *config.Config, client *ethclient.Client) error{
	log.Println("--------------清算队列数据整理开始-------------")
	lendpoolContract, _ := liquidatecontract.NewLendPool(common.HexToAddress(config.Contract.LendpoolContract), client)

	liquidateUser := new(models.LiquidateUser)
	users, err := engine.Where("id >?", 1).Rows(liquidateUser)
	if err != nil {
		log.Fatal(err)
		return err
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
			if len(currLiquidateEntry) == 0 && currUserBorrow[0].UnderlyingAsset != currUserCollateral[0].UnderlyingAsset {
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
	return nil
}
