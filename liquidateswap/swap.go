package liquidateswap

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"io/ioutil"
	"log"
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

	"swap/liquidatecontract"
	models "swap/models"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"k8s.io/klog"
)

func GetNonce(client *ethclient.Client, priKey string) (uint64, error) {
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
		return nonce, err
	}
	klog.Info("nonce----------------->", nonce)
	return nonce, nil
}

func FlashLoans(nonce *uint64, engine *xorm.Engine, liquidateQueue *models.LiquidateQueue, liquidateAndLoanContract *liquidatecontract.LiquidateLoan, client *ethclient.Client, priKey string, _assetToLiquidate common.Address, _flashAmt *big.Int, _collateral common.Address, _userToLiquidate common.Address, _amountOutMin *big.Int, _swapPath []common.Address) string {

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		klog.Error(err)
	}
	privateKey, err := crypto.HexToECDSA(priKey)
	if err != nil {
		klog.Error(err)
	}

	value := big.NewInt(0)
	//var gasPriceFloat big.Float
	//gasPriceFloat.SetString("23") //200 GWei

	//tenDecimal := big.NewFloat(math.Pow(10, float64(9)))
	//gasPrice, _ := new(big.Float).Mul(tenDecimal, &gasPriceFloat).Int(&big.Int{})
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
		return "fail"
	}

	gasLimit := uint64(3000101)

	var non big.Int
	//klog.Info("nonce----------------->", *nonce)
	non.SetUint64(*nonce)
	auth, _ := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	auth.Nonce = &non
	auth.Value = value
	auth.GasLimit = gasLimit
	auth.GasPrice = gasPrice

	tx, err := liquidateAndLoanContract.ExecuteFlashLoans(auth, _assetToLiquidate, _flashAmt, _collateral, _userToLiquidate, _amountOutMin, _swapPath)
	*nonce++
	if err != nil {
		klog.Error(err)
		return "fail"
	}
	klog.Info("liquidate tx: https://kovan.etherscan.io/tx/", tx.Hash())
	liquidateQueue.Status = "close"
	liquidateQueue.LiquidateTime = time.Now().Unix()
	engine.ID(liquidateQueue.Id).Update(liquidateQueue)
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
