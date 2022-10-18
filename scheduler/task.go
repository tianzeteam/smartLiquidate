package scheduler

import (
	"log"
	"math"
	"math/big"
	"swap/config"
	"swap/oraclecontract"
	"time"

	"swap/helpcontract"

	_ "github.com/go-sql-driver/mysql"
	"github.com/xormplus/xorm"

	"swap/liquidatecontract"
	models "swap/models"
	"swap/repository"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"k8s.io/klog"
)

/* func UpdateAssetTask(engine *xorm.Engine, liquidateContract *liquidatecontract.Liquidatecontract) {
	log.Println("--------------更新清算数据任务开始-------------")
	liquidateUser := new(models.LiquidateUser)
	users, err := engine.Where("id >?", 1).Rows(liquidateUser)
	if err != nil {
		log.Fatal(err)
	}
	defer users.Close()
	for users.Next() {
		err = users.Scan(liquidateUser)
		userId := liquidateUser.UserId
		if err != nil {
			klog.Error(err)
			continue
		}
		collateralReserves, borrowReserves, error := liquidateContract.GetLastUserAssetData(nil, common.HexToAddress(userId))
		if error != nil {
			klog.Errorln("userid ",userId)
			klog.Error(error)
			continue
		}
		//klog.Info("collateralReserves ", collateralReserves)
		//klog.Info("borrowReserves ", borrowReserves)
		_updateCollateral(engine, collateralReserves, userId)
		_updateBorrow(engine, borrowReserves, userId)

	}

} */

func UpdateAssetTask(engine *xorm.Engine, helpcontract *helpcontract.Helpcontract,priceOracle *oraclecontract.Oraclecontract,tokenDatas []models.TokenMeta) {
	log.Println("--------------更新清算数据任务开始-------------")
	liquidateUser := new(models.LiquidateUser)
	users, err := engine.Where("id >?", 1).Rows(liquidateUser)
	if err != nil {
		log.Fatal(err)
	}
	defer users.Close()

	for users.Next() {
		err = users.Scan(liquidateUser)
		userId := liquidateUser.UserId
		if err != nil {
			klog.Error(err)
			continue
		}
        var startTime = time.Now().UnixMilli()
		var collateralReserves []liquidatecontract.LiquidateLoanCollateralReserve   
        var borrowReserves []liquidatecontract.LiquidateLoanBorrowReserve
		nextId:
		for _,tokenData := range tokenDatas {
			tokenAddress :=tokenData.TokenAddress
			tokenNames := tokenData.Symbol
			tokenDecimals :=tokenData.Decimals
			assetReserves, error := helpcontract.GetUserReserveData(nil, tokenAddress,common.HexToAddress(userId))
			if error != nil {
				klog.Error(error)
				break  nextId
			}

            _currentATokenBalance := assetReserves.CurrentATokenBalance
            _currentStableDebt := assetReserves.CurrentStableDebt
            _currentVariableDebt := assetReserves.CurrentVariableDebt
           _usageAsCollateralEnabled := assetReserves.UsageAsCollateralEnabled
                 
           if _currentATokenBalance.Cmp(big.NewInt(0)) >0 {
			_priceInEth ,error := priceOracle.GetAssetPrice(nil,tokenAddress)
			if error !=nil {
				klog.Error(error)
				break nextId
			}
		     newReserve := liquidatecontract.LiquidateLoanCollateralReserve{CurrentATokenBalance:_currentATokenBalance,Symbol:tokenNames,UnderlyingAsset: tokenAddress,Decimals:tokenDecimals,UsageAsCollateralEnabled:_usageAsCollateralEnabled,PriceInEth:_priceInEth,NotNull:true}
			 collateralReserves = append(collateralReserves,newReserve)
		   }

		   if _currentStableDebt.Cmp( big.NewInt(0) ) >0 || _currentVariableDebt.Cmp(big.NewInt(0))>0 {
			_priceInEth ,error := priceOracle.GetAssetPrice(nil,tokenAddress)
			if error !=nil {
				klog.Error(error)
				break nextId
			}
			 _currentDebt  := _currentStableDebt.Add(_currentStableDebt,_currentVariableDebt)
			 newReserve := liquidatecontract.LiquidateLoanBorrowReserve{CurrentTotalDebt:_currentDebt,Symbol:tokenNames,UnderlyingAsset: tokenAddress,Decimals:tokenDecimals,UsageAsCollateralEnabled:_usageAsCollateralEnabled,PriceInEth:_priceInEth,NotNull:true}
			 borrowReserves = append(borrowReserves,newReserve)
		   }
		}
		var endTime = time.Now().UnixMilli()
		klog.Info("lostTime  ", endTime -startTime)
		klog.Info("collateralReserves ", collateralReserves)
		klog.Info("borrowReserves ", borrowReserves)
		_updateCollateral(engine, collateralReserves, userId)
		_updateBorrow(engine, borrowReserves, userId)

	}

}

func _updateBorrow(engine *xorm.Engine, borrowReserves []liquidatecontract.LiquidateLoanBorrowReserve, userId string) {
	var max_BorrowedValue = big.NewInt(0)
	var max_borrowedSymbol string
	var max_borrowedUnderlyingAsset string
	var max_borrowedPrincipal = "0"

	for _, borrowReserve := range borrowReserves {
		if borrowReserve.NotNull {
			currentTotalDebt := borrowReserve.CurrentTotalDebt
			priceInEth := borrowReserve.PriceInEth
			symbol := borrowReserve.Symbol
			underlyingAsset := borrowReserve.UnderlyingAsset
			decimals := borrowReserve.Decimals
			current_BorrowedValue := big.NewInt(0)
			var mulExp, base = big.NewInt(1), big.NewInt(10)
			mulExp.Exp(base, decimals, nil)

			current_BorrowedValue = current_BorrowedValue.Mul(currentTotalDebt, priceInEth).Div(current_BorrowedValue, mulExp)
			if current_BorrowedValue.Cmp(max_BorrowedValue) >= 0 {
				*max_BorrowedValue = *current_BorrowedValue
				max_borrowedSymbol = symbol
				max_borrowedUnderlyingAsset = underlyingAsset.String()
				max_borrowedPrincipal = currentTotalDebt.String()
			}
		}

	}

	currAssets := make([]models.AaveAsset, 0)
	session := engine.NewSession()
	defer session.Close()
	err := session.Begin()
	if err != nil {
		session.Rollback()
		return
	}
	err = session.Where("user_id=?", userId).And("asset_type = ?", "borrow").Find(&currAssets)
	if err != nil {
		log.Fatalln(err)
	}
	if len(currAssets) == 0 && len(max_borrowedUnderlyingAsset) != 0 {
		aaveAsset := &models.AaveAsset{}
		aaveAsset.UserId = userId
		aaveAsset.AssetType = "borrow"
		aaveAsset.Amount = max_borrowedPrincipal
		aaveAsset.Status = "open"
		aaveAsset.UnderlyingAsset = max_borrowedUnderlyingAsset
		aaveAsset.Symbol = max_borrowedSymbol
		aaveAsset.CreateTime = time.Now().Unix()
		_, err = session.InsertOne(aaveAsset)
	} else {
		if len(max_borrowedUnderlyingAsset) != 0 {
			currAssets[0].UserId = userId
			currAssets[0].AssetType = "borrow"
			currAssets[0].Amount = max_borrowedPrincipal
			currAssets[0].Status = "open"
			currAssets[0].UnderlyingAsset = max_borrowedUnderlyingAsset
			currAssets[0].Symbol = max_borrowedSymbol
			currAssets[0].UpdateTime = time.Now().Unix()
			_, err = session.ID(currAssets[0].Id).Update(currAssets[0])
		}
	}
	if err != nil {
		session.Rollback()
		return
	}
	// add Commit() after all actions
	err = session.Commit()
	if err != nil {
		return
	}
}
func _updateCollateral(engine *xorm.Engine, collateralReserves []liquidatecontract.LiquidateLoanCollateralReserve, userId string) {
	var max_collateralValue = big.NewInt(0)
	var max_collateralSymbol string
	var max_collateralUnderlyingAsset string
	var max_collateralPrincipal string

	for _, collateralReserve := range collateralReserves {
		if collateralReserve.NotNull {
			_currentTotalAToken := collateralReserve.CurrentATokenBalance
			_priceInEth := collateralReserve.PriceInEth
			_symbol := collateralReserve.Symbol
			_decimals := collateralReserve.Decimals
			_underlyingAsset := collateralReserve.UnderlyingAsset
			current_collateralValue := big.NewInt(0)
			var mulExp, base = big.NewInt(1), big.NewInt(10)
			mulExp.Exp(base, _decimals, nil)
			current_collateralValue = current_collateralValue.Mul(_currentTotalAToken, _priceInEth).Div(current_collateralValue, mulExp)
			if current_collateralValue.Cmp(max_collateralValue) >= 0 {
				*max_collateralValue = *current_collateralValue
				max_collateralSymbol = _symbol
				max_collateralUnderlyingAsset = _underlyingAsset.String()
				max_collateralPrincipal = _currentTotalAToken.String()
			}
		}

	}
	currAssets := make([]models.AaveAsset, 0)
	errs := engine.Where("user_id=?", userId).And("asset_type = ?", "collateral").Find(&currAssets)
	if errs != nil {
		log.Fatalln(errs)
	}
	if len(currAssets) == 0 && len(max_collateralUnderlyingAsset) != 0 {
		aaveAsset := &models.AaveAsset{}
		aaveAsset.UserId = userId
		aaveAsset.AssetType = "collateral"
		aaveAsset.Amount = max_collateralPrincipal
		aaveAsset.Status = "open"
		aaveAsset.UnderlyingAsset = max_collateralUnderlyingAsset
		aaveAsset.Symbol = max_collateralSymbol
		aaveAsset.CreateTime = time.Now().Unix()
		engine.InsertOne(aaveAsset)
	} else {
		if len(max_collateralUnderlyingAsset) != 0 {
			currAssets[0].UserId = userId
			currAssets[0].AssetType = "collateral"
			currAssets[0].Amount = max_collateralPrincipal
			currAssets[0].Status = "open"
			currAssets[0].UnderlyingAsset = max_collateralUnderlyingAsset
			currAssets[0].Symbol = max_collateralSymbol
			currAssets[0].UpdateTime = time.Now().Unix()
			engine.ID(currAssets[0].Id).Update(currAssets[0])
		}
	}
}
func FetchTask(engine *xorm.Engine) error {
	log.Println("--------------fetch 清算数据任务开始----------------")
	repository.FetchLiquidateData(engine)
	return nil
}
func Task(engine *xorm.Engine, config *config.Config, client *ethclient.Client) error {
	log.Println("--------------清算队列数据整理任务开始-------------")
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
		userCollateralETH := big.NewFloat(0).SetInt(HeathFactorData.TotalCollateralETH)
		decimals := big.NewFloat(math.Pow(10, float64(18))) //精度
		userStandardHealthFactor := new(big.Float).Quo(userHealthFactor, decimals)
		userCollateralEthWithDecimals := new(big.Float).Quo(userCollateralETH, decimals)

		PENDING_HEALTH_THRESHOLD := config.Liquidate.PENDING_HEALTH_THRESHOLD
		MAX_COLLATERAL_THRESHOLD := config.Liquidate.MAX_COLLATERAL_THRESHOLD
		currentUserHealthValue, _ := userStandardHealthFactor.Float64()
		userCollateralEthWithDecimalsFloat, _ := userCollateralEthWithDecimals.Float64()

		if currentUserHealthValue < PENDING_HEALTH_THRESHOLD && userCollateralEthWithDecimalsFloat > MAX_COLLATERAL_THRESHOLD {

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
