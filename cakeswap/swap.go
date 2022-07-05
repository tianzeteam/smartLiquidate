package cakeswap

import (
	"context"
	"crypto/ecdsa"
	"math"
	"math/big"
	"time"

	"swap/cakecontract"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"k8s.io/klog"
)

const (
	PancakeSwapContract = "0x10ED43C718714eb63d5aA57B78B54704E256024E"
)

func Swap(client *ethclient.Client, priKey string, path []common.Address, amountIn, amountOutMin *big.Int) string {
	b, err := cakecontract.NewCakecontract(common.HexToAddress(PancakeSwapContract), client)
	if err != nil {
		klog.Error(err)
	}

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
	gasPriceFloat.SetString("5") //200 GWei

	tenDecimal := big.NewFloat(math.Pow(10, float64(9)))
	gasPrice, _ := new(big.Float).Mul(tenDecimal, &gasPriceFloat).Int(&big.Int{})

	gasLimit := uint64(310101)

	var non big.Int
	non.SetUint64(nonce)
	auth, _ := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	auth.Nonce = &non
	auth.Value = value
	auth.GasLimit = gasLimit
	auth.GasPrice = gasPrice

	toAddr := fromAddress
	time.Local = time.FixedZone("CST", 0)
	nowInt64 := time.Now().Unix()
	deadline := big.NewInt(nowInt64 + 5*60)

	tx, err := b.SwapExactTokensForTokens(auth, amountIn, amountOutMin, path, toAddr, deadline)
	if err != nil {
		klog.Error(err)
	}
	klog.Info("swap tx: https://bscscan.com/tx/", tx.Hash())
	return tx.Hash().Hex()
}
