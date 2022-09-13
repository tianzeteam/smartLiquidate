package utils

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"swap/config"
	"swap/liquidatecontract"

	"github.com/ethereum/go-ethereum/common"
	"github.com/miraclesu/uniswap-sdk-go/constants"
	"github.com/miraclesu/uniswap-sdk-go/entities"

	//"github.com/miraclesu/uniswap-sdk-go/constants"
	//"github.com/miraclesu/uniswap-sdk-go/entities"
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
	return nonce, nil
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

func GetBestTradeExactIn(CollateralAsset common.Address, BorrowAsset common.Address, amountIn *big.Int, uniswapFactoryContract *liquidatecontract.Factory, config *config.Config) []common.Address {

	pair, err := uniswapFactoryContract.GetPair(nil, CollateralAsset, BorrowAsset)
	if err != nil {
		klog.Fatalln("清算进程中获取货币对异常")
		return nil
	}
	var swapPath []common.Address
	if pair.Hex() != "0x0000000000000000000000000000000000000000" {
		swapPath = []common.Address{
			CollateralAsset, //用奖励抵押物资产 兑换 借贷资产
			BorrowAsset,
		}
		return swapPath
	}

	if strings.ToLower(CollateralAsset.Hex()) == "0xb597cd8d3217ea6477232f9217fa70837ff667af" || strings.ToLower(CollateralAsset.Hex()) == "0xd0a1e359811322d97991e03f863a0c30c2cf029c" || strings.ToLower(CollateralAsset.Hex()) == "0x13512979ade267ab5100878e2e0f485b568328a4" || strings.ToLower(CollateralAsset.Hex()) == "0xff795577d9ac8bd7d90ee22b6c1703490b6512fd" || strings.ToLower(BorrowAsset.Hex()) == "0xb597cd8d3217ea6477232f9217fa70837ff667af" || strings.ToLower(BorrowAsset.Hex()) == "0xd0a1e359811322d97991e03f863a0c30c2cf029c" || strings.ToLower(BorrowAsset.Hex()) == "0x13512979ade267ab5100878e2e0f485b568328a4" || strings.ToLower(BorrowAsset.Hex()) == "0xff795577d9ac8bd7d90ee22b6c1703490b6512fd" {
		swapPath = []common.Address{
			CollateralAsset, //用奖励抵押物资产 兑换 借贷资产
			BorrowAsset,
		}
		return swapPath
	}

	aaveToken, _ := entities.NewToken(constants.ChainID(config.ChainId.Network), common.HexToAddress(config.BaseToken.Aave), 18, "AAVE", "")
	wethToken, _ := entities.NewToken(constants.ChainID(config.ChainId.Network), common.HexToAddress(config.BaseToken.Weth), 18, "WETH", "")
	usdtToken, _ := entities.NewToken(constants.ChainID(config.ChainId.Network), common.HexToAddress(config.BaseToken.Usdt), 18, "USDT", "")
	daiToken, _ := entities.NewToken(constants.ChainID(config.ChainId.Network), common.HexToAddress(config.BaseToken.Dai), 18, "DAI", "")
	collateralToken, _ := entities.NewToken(constants.ChainID(config.ChainId.Network), CollateralAsset, 18, "COLLATERAL", "")
	borrowToken, _ := entities.NewToken(constants.ChainID(config.ChainId.Network), BorrowAsset, 18, "BORROW", "")

	tokenAmount_aave, _ := entities.NewTokenAmount(aaveToken, big.NewInt(10000))
	tokenAmount_weth, _ := entities.NewTokenAmount(wethToken, big.NewInt(10000))
	tokenAmount_usdt, _ := entities.NewTokenAmount(usdtToken, big.NewInt(10000))
	tokenAmount_dai, _ := entities.NewTokenAmount(daiToken, big.NewInt(10000))
	tokenAmount_collateral, _ := entities.NewTokenAmount(collateralToken, amountIn)
	tokenAmount_borrow, _ := entities.NewTokenAmount(borrowToken, big.NewInt(10000))
	x0_collateral_pair, _ := entities.NewPair(tokenAmount_dai, tokenAmount_collateral)
	x1_collateral_pair, _ := entities.NewPair(tokenAmount_usdt, tokenAmount_collateral)
	x2_collateral_pair, _ := entities.NewPair(tokenAmount_aave, tokenAmount_collateral)
	x3_collateral_pair, _ := entities.NewPair(tokenAmount_weth, tokenAmount_collateral)
	dai_borrow_pair, _ := entities.NewPair(tokenAmount_dai, tokenAmount_borrow)
	aave_usdt_pair, _ := entities.NewPair(tokenAmount_aave, tokenAmount_usdt)
	aave_weth_pair, _ := entities.NewPair(tokenAmount_aave, tokenAmount_weth)
	weth_usdt_pair, _ := entities.NewPair(tokenAmount_weth, tokenAmount_usdt)
	aave_dai_pair, _ := entities.NewPair(tokenAmount_aave, tokenAmount_dai)
	dai_usdt_pair, _ := entities.NewPair(tokenAmount_dai, tokenAmount_usdt)

	pairs := []*entities.Pair{x0_collateral_pair, x1_collateral_pair, aave_usdt_pair, x2_collateral_pair, x3_collateral_pair, aave_weth_pair, aave_dai_pair, weth_usdt_pair, dai_borrow_pair, dai_usdt_pair}
	results, _ := entities.BestTradeExactIn(pairs, tokenAmount_collateral, borrowToken, &entities.BestTradeOptions{MaxNumResults: 4, MaxHops: 3}, nil, tokenAmount_collateral, nil)

	for _, result := range results {
		tokens := result.Route.Path

		if len(tokens) > 2 {
			for _, token := range tokens {
				swapPath = append(swapPath, token.Address)

			}
			fmt.Println(len(swapPath))
			return swapPath
		}
	}
	return []common.Address{}
}
