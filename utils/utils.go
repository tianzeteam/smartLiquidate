package utils

import (
	"fmt"
	"math/big"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"swap/liquidatecontract"

	"github.com/ethereum/go-ethereum/common"
	"github.com/miraclesu/uniswap-sdk-go/constants"
	"github.com/miraclesu/uniswap-sdk-go/entities"
)

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

func GetBestTradeExactIn(CollateralAsset string, BorrowAsset string, amountIn *big.Int, uniswapFactoryContract *liquidatecontract.Factory) []common.Address {

	aaveToken, _ := entities.NewToken(constants.Kovan, common.HexToAddress("0xb597cd8d3217ea6477232f9217fa70837ff667af"), 18, "AAVE", "")
	wethToken, _ := entities.NewToken(constants.Kovan, common.HexToAddress("0xd0a1e359811322d97991e03f863a0c30c2cf029c"), 18, "WETH", "")
	usdtToken, _ := entities.NewToken(constants.Kovan, common.HexToAddress("0x13512979ade267ab5100878e2e0f485b568328a4"), 18, "USDT", "")
	daiToken, _ := entities.NewToken(constants.Kovan, common.HexToAddress("0xff795577d9ac8bd7d90ee22b6c1703490b6512fd"), 18, "DAI", "")
	collateralToken, _ := entities.NewToken(constants.Kovan, common.HexToAddress(CollateralAsset), 18, "COLLATERAL", "")
	borrowToken, _ := entities.NewToken(constants.Kovan, common.HexToAddress(BorrowAsset), 18, "BORROW", "")

	tokenAmount_aave, _ := entities.NewTokenAmount(aaveToken, big.NewInt(10000))
	tokenAmount_weth, _ := entities.NewTokenAmount(wethToken, big.NewInt(10000))
	tokenAmount_usdt, _ := entities.NewTokenAmount(usdtToken, big.NewInt(10000))
	tokenAmount_dai, _ := entities.NewTokenAmount(daiToken, big.NewInt(10000))
	tokenAmount_collateral, _ := entities.NewTokenAmount(collateralToken, amountIn)
	tokenAmount_borrow, _ := entities.NewTokenAmount(borrowToken, big.NewInt(10000))
	aave_collateral_pair, _ := entities.NewPair(tokenAmount_aave, tokenAmount_collateral)
	dai_borrow_pair, _ := entities.NewPair(tokenAmount_dai, tokenAmount_borrow)
	aave_usdt_pair, _ := entities.NewPair(tokenAmount_aave, tokenAmount_usdt)
	aave_weth_pair, _ := entities.NewPair(tokenAmount_aave, tokenAmount_weth)
	weth_usdt_pair, _ := entities.NewPair(tokenAmount_weth, tokenAmount_usdt)
	aave_dai_pair, _ := entities.NewPair(tokenAmount_aave, tokenAmount_dai)
	dai_usdt_pair, _ := entities.NewPair(tokenAmount_dai, tokenAmount_usdt)

	pairs := []*entities.Pair{aave_collateral_pair, aave_usdt_pair, aave_weth_pair, weth_usdt_pair, aave_dai_pair, dai_usdt_pair,dai_borrow_pair}
	result, _ := entities.BestTradeExactIn(pairs, tokenAmount_collateral, borrowToken, &entities.BestTradeOptions{MaxNumResults: 5, MaxHops: 3}, nil, tokenAmount_collateral, nil)

	var tests = []struct {
		expect int
		output int
	}{
		{2, len(result)},
		{1, len(result[0].Route.Pairs)},
		{2, len(result[1].Route.Pairs)},
	}
	for i, test := range tests {
		if test.expect != test.output {
			fmt.Errorf("test #%d: expect[%+v], but got[%+v]", i, test.expect, test.output)
		}
	}
	return []common.Address{}
}
