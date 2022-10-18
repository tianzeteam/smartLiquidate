package models

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type AaveAsset struct {
	CreateTime      int64  `xorm:"not null INT(11)"`
	UpdateTime      int64  `xorm:"not null INT(11)"`
	UserId          string `xorm:"not null VARCHAR(255)"`
	UnderlyingAsset string `xorm:"not null VARCHAR(255)"`
	Status          string `xorm:"not null ENUM('close','open')"`
	Id              int    `xorm:"not null pk autoincr INT(11)"`
	Symbol          string `xorm:"not null VARCHAR(255)"`
	AssetType       string `xorm:"not null ENUM('borrow','collateral')"`
	Amount          string `xorm:"not null VARCHAR(255)"`
}

type LiquidateUser struct {
	Id     int    `xorm:"not null pk autoincr INT(11)"`
	UserId string `xorm:"not null VARCHAR(255)"`
}

type LiquidateResult struct {
	Id           int    `xorm:"not null pk autoincr INT(11)"`
	UserId       string `xorm:"not null VARCHAR(255)"`
	ReceiptAsset string `xorm:"not null VARCHAR(255)"`
	TxHash       string `xorm:"not null VARCHAR(255)"`
	CreateTime   int64  `xorm:"not null INT(11)"`
	Status       string `xorm:"not null ENUM('pending','fail','success')"`
}

type LiquidateQueue struct {
	Id              int    `xorm:"not null pk autoincr INT(11)"`
	UserId          string `xorm:"not null VARCHAR(255)"`
	CollateralAsset string `xorm:"not null VARCHAR(255)"`
	BorrowAsset     string `xorm:"not null VARCHAR(255)"`
	BorrowAmount    string `xorm:"not null VARCHAR(255)"`
	CreateTime      int64  `xorm:"not null INT(11)"`
	LiquidateTime   int64  `xorm:"not null INT(11)"`
	Status          string `xorm:"not null ENUM('close','waiting')"`
}

type USER struct {
	ID                    string `json:"id"`
	BorrowedReservesCount int    `json:"borrowedReservesCount"`
	CollateralReserve     []struct {
		CurrentATokenBalance string `json:"currentATokenBalance"`
		Reserve              struct {
			UsageAsCollateralEnabled    bool   `json:"usageAsCollateralEnabled"`
			ReserveLiquidationThreshold string `json:"reserveLiquidationThreshold"`
			ReserveLiquidationBonus     string `json:"reserveLiquidationBonus"`
			BorrowingEnabled            bool   `json:"borrowingEnabled"`
			UtilizationRate             string `json:"utilizationRate"`
			Symbol                      string `json:"symbol"`
			UnderlyingAsset             string `json:"underlyingAsset"`
			Price                       struct {
				PriceInEth string `json:"priceInEth"`
			} `json:"price"`
			Decimals int `json:"decimals"`
		} `json:"reserve"`
	} `json:"collateralReserve"`
	BorrowReserve []struct {
		CurrentTotalDebt string `json:"currentTotalDebt"`
		Reserve          struct {
			UsageAsCollateralEnabled    bool   `json:"usageAsCollateralEnabled"`
			ReserveLiquidationThreshold string `json:"reserveLiquidationThreshold"`
			BorrowingEnabled            bool   `json:"borrowingEnabled"`
			UtilizationRate             string `json:"utilizationRate"`
			Symbol                      string `json:"symbol"`
			UnderlyingAsset             string `json:"underlyingAsset"`
			Price                       struct {
				PriceInEth string `json:"priceInEth"`
			} `json:"price"`
			Decimals int `json:"decimals"`
		} `json:"reserve"`
	} `json:"borrowReserve"`
}

type TokenMeta struct {
	Symbol       string
	TokenAddress common.Address
	Decimals     *big.Int
}
