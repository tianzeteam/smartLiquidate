package main

import (
	"context"
	"math"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"k8s.io/klog"
	"swap/bakecontract"
	"swap/bakeswap"
)

const (
	BscMainNetHttpUrl = "https://bsc-dataseed4.binance.org"
	AccountPriKey     = "85204df7d8354cdaa080ee06aea8dd25fe6a6031d22d73f8e12b883e8d9eb686"
	AccountAddr       = "0x03F5d1F7C36477f7641399BFCf2179A06A297679"

	BakeSwapContract   = "0xCDe540d7eAFE93aC5fE6233Bee57E1270D3E330F"
	BakeTokenContract  = "0xE02dF9e3e622DeBdD69fb838bB799E3F168902c5"
	DoggyTokenContract = "0x74926b3d118a63f6958922d3dc05eb9c6e6e00c6"
)

func main() {
	client, err := ethclient.Dial(BscMainNetHttpUrl)
	if err != nil {
		klog.Fatal(err)
	}

	BuyPath := []common.Address{
		common.HexToAddress(BakeTokenContract), //用bake兑换doggy
		common.HexToAddress(DoggyTokenContract),
	}

	opts := &bind.CallOpts{
		From:    common.HexToAddress(AccountAddr),
		Context: context.Background(),
	}
	b, err := bakecontract.NewBakecontract(common.HexToAddress(BakeSwapContract), client)
	if err != nil {
		klog.Error(err)
	}

	var StartAmountFloat big.Float
	StartAmountFloat.SetString("1")                       //用1个买
	tenDecimal := big.NewFloat(math.Pow(10, float64(18))) //精度
	StartAmount, _ := new(big.Float).Mul(tenDecimal, &StartAmountFloat).Int(&big.Int{})

	AmountOutFromContract, err := b.GetAmountsOut(opts, StartAmount, BuyPath)
	if err != nil {
		klog.Error(err)
	}
	klog.Info("AmountOutFromContract: ", AmountOutFromContract)

	// 20%滑点
	NewAmountOut := AmountOutFromContract[len(AmountOutFromContract)-1]
	ammountOutMin := NewAmountOut.Mul(NewAmountOut, big.NewInt(8)).Div(NewAmountOut, big.NewInt(10))

	tx := bakeswap.Swap(client, AccountPriKey, BuyPath, StartAmount, ammountOutMin)

	for true {
		_, pending, err := client.TransactionByHash(context.Background(), common.HexToHash(tx))
		if err != nil {
			klog.Error(err)
			continue
		}
		if pending {
			time.Sleep(time.Millisecond * 500)
			continue
		}
		r, err := client.TransactionReceipt(context.Background(), common.HexToHash(tx))
		if err != nil {
			klog.Error(err)
			continue
		}
		klog.Info("sale tx status: ", r.Status)
		os.Exit(0)
	}
}
