package liquidateswap

import (
	"context"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"swap/config"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/xormplus/xorm"
	"gopkg.in/yaml.v2"

	"swap/counter"
	"swap/liquidatecontract"
	models "swap/models"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"k8s.io/klog"
)

func FlashLoans(_counter counter.Counter, engine *xorm.Engine, liquidateQueue *models.LiquidateQueue, liquidateAndLoanContract *liquidatecontract.LiquidateLoan, client *ethclient.Client, priKey string, _assetToLiquidate common.Address, _flashAmt *big.Int, _collateral common.Address, _userToLiquidate common.Address, _amountOutMin *big.Int, _swapPath []common.Address, userId string, _gasPrice string, _gasLimit int,_gasCost int64) string {

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		klog.Error(err)
	}
	privateKey, err := crypto.HexToECDSA(priKey)
	if err != nil {
		klog.Error(err)
	}

	value := big.NewInt(0)
	var gasPriceFloat big.Float
	gasPriceFloat.SetString(_gasPrice) //200 GWei
	gasPrice, _ := gasPriceFloat.Int(&big.Int{})
	gasLimit := uint64(_gasLimit)

	var non big.Int
	nonce := _counter.Add(1)
	//klog.Info("nonce----------------->", nonce,liquidateQueue.Id)
	non.SetUint64(nonce)
	auth, _ := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	auth.Nonce = &non
	auth.Value = value
	auth.GasLimit = gasLimit
	auth.GasPrice = gasPrice

	gasCost := big.NewInt(_gasCost);

	tx, err := liquidateAndLoanContract.ExecuteFlashLoans(auth, _assetToLiquidate, _flashAmt, _collateral, _userToLiquidate, _amountOutMin, _swapPath,gasCost)

	if err != nil {
		klog.Error(err)
		return "fail"
	}
	klog.Info("tx: https://kovan.etherscan.io/tx/", tx.Hash(), "  nonce::", nonce)
	liquidateQueue.Status = "close"
	liquidateQueue.LiquidateTime = time.Now().Unix()
	//klog.Info(" liquidateQueue.Id  ", liquidateQueue.Id)
	engine.ID(liquidateQueue.Id).Update(liquidateQueue)

	liquidateResult := new(models.LiquidateResult)
	liquidateResult.UserId = userId
	liquidateResult.ReceiptAsset = liquidateQueue.BorrowAsset
	liquidateResult.TxHash = tx.Hash().Hex()
	liquidateResult.CreateTime = time.Now().Unix()
	liquidateResult.Status = "pending"
	engine.InsertOne(liquidateResult)

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
