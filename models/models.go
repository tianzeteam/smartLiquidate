package models

type AaveAsset struct {
	CreateTime      int    `xorm:"not null INT(11)"`
	UpdateTime      int    `xorm:"not null INT(11)"`
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
