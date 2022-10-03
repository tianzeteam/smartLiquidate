package repository

import (
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"strconv"
	models "swap/models"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/xormplus/xorm"
	"gopkg.in/jmcvetta/napping.v3"
)

var url string = "https://api.thegraph.com/subgraphs/name/aave/protocol-v2"

type graphqlRequest struct {
	Query     string                 `json:"query"`
	Variables map[string]interface{} `json:"variables"`
}

type RespUserData struct {
	DATA struct {
		USERS []models.USER `json:"users"`
	} `json:"data"`
}

func FetchLiquidateData(engine *xorm.Engine) {
	// Query
	query := `query GET_LOANS {
	  users(first:1000, skip:%s, orderBy: id, orderDirection: desc, where: {borrowedReservesCount_gt: 0}) {
		id
		borrowedReservesCount
		collateralReserve:reserves(where: {currentATokenBalance_gt: 0}) {
		  currentATokenBalance
		  reserve{
			usageAsCollateralEnabled
			reserveLiquidationThreshold
			reserveLiquidationBonus
			borrowingEnabled
			utilizationRate
			symbol
			underlyingAsset
			price {
			  priceInEth
			}
			decimals
		  }
		}
		borrowReserve: reserves(where: {currentTotalDebt_gt: 0}) {
		  currentTotalDebt
		  reserve{
			usageAsCollateralEnabled
			reserveLiquidationThreshold
			borrowingEnabled
			utilizationRate
			symbol
			underlyingAsset
			price {
			  priceInEth
			}
			decimals
		  }
		}
	  }
	}`

	header := &http.Header{}
	header.Set("Content-Type", "application/json")

	for i := 0; i < 6; i++ {
		temp := query
		skipNum := i * 1000
		temp = fmt.Sprintf(temp, strconv.Itoa(skipNum))
		session := napping.Session{}
		session.Header = header

		mashalled, _ := json.Marshal(graphqlRequest{Query: temp})
		var data map[string]json.RawMessage
		err := json.Unmarshal(mashalled, &data)
		if err != nil {
			fmt.Println(err)
		}

		resp, err := session.Post(url, &data, nil, nil)
		if err != nil {
			log.Fatal(err)
		}

		updateData(engine, resp.RawText())
		time.Sleep(time.Duration(10) * time.Second)
	}
}

func updateData(engine *xorm.Engine, userData string) {

	var RespData RespUserData
	json.Unmarshal([]byte(userData), &RespData)

	for _, user := range RespData.DATA.USERS {
		//fmt.Println("-------fetch user data one by one  --->", user.ID)
		parseUserData(engine, user.ID)
		//parseBorrowReserveData(engine, user)
		//parseCollateralReserveData(engine, user)
	}

}
func parseCollateralReserveData(engine *xorm.Engine, user models.USER) {
	var max_collateralValue = big.NewInt(0)
	var max_collateralSymbol string
	var max_collateralUnderlyingAsset string
	var max_collateralPrincipal string
	userId := user.ID
	collateralReserves := user.CollateralReserve
	for _, collateralReserve := range collateralReserves {

		_currentTotalAToken := collateralReserve.CurrentATokenBalance
		currentTotalAToken, _ := new(big.Int).SetString(_currentTotalAToken, 10)
		reserve := collateralReserve.Reserve
		price, _ := new(big.Int).SetString(reserve.Price.PriceInEth, 10)

		current_collateralValue := big.NewInt(0)
		var mulExp, base, decimals = big.NewInt(1), big.NewInt(10), big.NewInt(int64(reserve.Decimals))
		mulExp.Exp(base, decimals, nil)
		current_collateralValue = current_collateralValue.Mul(currentTotalAToken, price).Div(current_collateralValue, mulExp)
		if current_collateralValue.Cmp(max_collateralValue) > 0 {
			*max_collateralValue = *current_collateralValue
			max_collateralSymbol = reserve.Symbol
			max_collateralUnderlyingAsset = reserve.UnderlyingAsset
			max_collateralPrincipal = _currentTotalAToken
		}

	}
	currAssets := make([]models.AaveAsset, 0)
	errs := engine.Where("user_id=?", userId).And("asset_type = ?", "collateral").Find(&currAssets)
	if errs != nil {
		log.Fatalln(errs)
	}
	if len(currAssets) == 0 {
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
func parseBorrowReserveData(engine *xorm.Engine, user models.USER) {
	var max_BorrowedValue = big.NewInt(0)
	var max_borrowedSymbol string
	var max_borrowedUnderlyingAsset string
	var max_borrowedPrincipal = "0"
	userId := user.ID
	borrowReserves := user.BorrowReserve
	for _, borrowReserve := range borrowReserves {

		_currentTotalDebt := borrowReserve.CurrentTotalDebt
		currentTotalDebt, _ := new(big.Int).SetString(_currentTotalDebt, 10)
		reserve := borrowReserve.Reserve

		price, _ := new(big.Int).SetString(reserve.Price.PriceInEth, 10)
		current_BorrowedValue := big.NewInt(0)
		var mulExp, base, decimals = big.NewInt(1), big.NewInt(10), big.NewInt(int64(borrowReserve.Reserve.Decimals))
		mulExp.Exp(base, decimals, nil)

		current_BorrowedValue = current_BorrowedValue.Mul(currentTotalDebt, price).Div(current_BorrowedValue, mulExp)
		if current_BorrowedValue.Cmp(max_BorrowedValue) > 0 {
			*max_BorrowedValue = *current_BorrowedValue
			max_borrowedSymbol = reserve.Symbol
			max_borrowedUnderlyingAsset = reserve.UnderlyingAsset
			max_borrowedPrincipal = _currentTotalDebt
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
	if len(currAssets) == 0 {
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
		currAssets[0].UserId = userId
		currAssets[0].AssetType = "borrow"
		currAssets[0].Amount = max_borrowedPrincipal
		currAssets[0].Status = "open"
		currAssets[0].UnderlyingAsset = max_borrowedUnderlyingAsset
		currAssets[0].Symbol = max_borrowedSymbol
		currAssets[0].UpdateTime = time.Now().Unix()
		_, err = session.ID(currAssets[0].Id).Update(currAssets[0])
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

func parseUserData(engine *xorm.Engine, userId string) {
	user := &models.LiquidateUser{}
	count, _ := engine.Where("user_id=?", userId).Count(user)
	if count == 0 {
		user.UserId = userId
		engine.InsertOne(user)
	}
}
