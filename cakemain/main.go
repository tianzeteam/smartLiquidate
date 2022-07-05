package main

import (
	"context"
	"math"
	"math/big"
	"os"
	"time"

	"swap/cakecontract"
	"swap/cakeswap"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"k8s.io/klog"
)

const (
	BscMainNetHttpUrl = "https://bsc.mytokenpocket.vip"
	AccountPriKey     = "a030096bcadebf2a915c3f39bc8476f2d2089aec8bb626f9b689f2824d69df5f"
	AccountAddr       = "0x14b8e58C8639b3991df607B79042dF0150C601dd"

	PancakeSwapContract = "0x10ED43C718714eb63d5aA57B78B54704E256024E"
	BakeTokenContract   = "0xE02dF9e3e622DeBdD69fb838bB799E3F168902c5"
	CakeTokenContract   = "0x0E09FaBB73Bd3Ade0a17ECC321fD13a19e81cE82"
	BusdTokenContract   = "0xe9e7CEA3DedcA5984780Bafc599bD69ADd087D56"
	DoggyTokenContract  = "0x74926b3d118a63f6958922d3dc05eb9c6e6e00c6"
)

func main() {
	client, err := ethclient.Dial(BscMainNetHttpUrl)
	if err != nil {
		klog.Fatal(err)
	}

	BuyPath := []common.Address{
		common.HexToAddress(BakeTokenContract), //用bake兑换doggy
		common.HexToAddress(CakeTokenContract), //cake
		common.HexToAddress(BusdTokenContract), //busd
		//common.HexToAddress(DoggyTokenContract), //doggy
	}

	opts := &bind.CallOpts{
		From:    common.HexToAddress(AccountAddr),
		Context: context.Background(),
	}
	b, err := cakecontract.NewCakecontract(common.HexToAddress(PancakeSwapContract), client)
	if err != nil {
		klog.Error(err)
	}

	var StartAmountFloat big.Float
	StartAmountFloat.SetString("0.1")                     //用1个买
	tenDecimal := big.NewFloat(math.Pow(10, float64(18))) //精度
	StartAmount, _ := new(big.Float).Mul(tenDecimal, &StartAmountFloat).Int(&big.Int{})

	AmountOutFromContract, err := b.GetAmountsOut(opts, StartAmount, BuyPath)
	if err != nil {
		klog.Error(err)
	}
	klog.Info("AmountOutFromContract: ", AmountOutFromContract)

	// 20%滑点
	NewAmountOut := AmountOutFromContract[len(AmountOutFromContract)-1]
	klog.Info("NewAmountOut:::", NewAmountOut)
	ammountOutMin := NewAmountOut.Mul(NewAmountOut, big.NewInt(8)).Div(NewAmountOut, big.NewInt(10))
	
	klog.Info("ammountOutMin:::", ammountOutMin)
	tx := cakeswap.Swap(client, AccountPriKey, BuyPath, StartAmount, ammountOutMin)
	klog.Info("jump into for: ", "")
	for {
		_, pending, transactionErr := client.TransactionByHash(context.Background(), common.HexToHash(tx))
		klog.Info("pending: ", pending)
		if transactionErr != nil {
			klog.Error(transactionErr)
			continue
		}
		if pending {
			time.Sleep(time.Millisecond * 500)
			continue
		}
		r, err := client.TransactionReceipt(context.Background(), common.HexToHash(tx))
		klog.Info("err: ", err)
		if err != nil {
			klog.Error(err)
			continue
		}
		klog.Info("sale tx status: ", r.Status)
		os.Exit(0)
	}
}
