package handler

import (
	"fmt"
	"log"
	"math/big"
	"strings"
	"time"

	core "github.com/daoleno/uniswap-sdk-core/entities"
	coreEntities "github.com/daoleno/uniswap-sdk-core/entities"
	"github.com/daoleno/uniswapv3-sdk/constants"
	"github.com/daoleno/uniswapv3-sdk/entities"
	"github.com/daoleno/uniswapv3-sdk/examples/contract"
	"github.com/daoleno/uniswapv3-sdk/examples/helper"
	"github.com/daoleno/uniswapv3-sdk/periphery"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// 使用uniswapv3 sdk与uniswap交互
// go get github.com/daoleno/uniswapv3-sdk

var (
	USDC     = core.NewToken(1, common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"), 6, "USDC", "USD Coin")
	DAI      = core.NewToken(1, common.HexToAddress("0x6B175474E89094C44Da98b954EedeAC495271d0F"), 18, "DAI", "Dai Stablecoin")
	OneEther = big.NewInt(1e18)
)

func Uniswap() {
	Quoter()
	// Swap()
	// mintOrAdd(nil)
}

// 与Quoter合约交互
func Quoter() {
	client, err := ethclient.Dial(helper.PolygonRPC)
	if err != nil {
		panic(err)
	}
	quoterContract, err := contract.NewUniswapv3Quoter(common.HexToAddress(helper.ContractV3Quoter), client)
	if err != nil {
		panic(err)
	}

	token0 := common.HexToAddress(helper.WMaticAddr)
	token1 := common.HexToAddress(helper.AmpAddr)
	fee := big.NewInt(3000)
	amountIn := helper.FloatStringToBigInt("1.00", 18)
	sqrtPriceLimitX96 := big.NewInt(0)

	var out []interface{}
	rawCaller := &contract.Uniswapv3QuoterRaw{Contract: quoterContract}
	err = rawCaller.Call(nil, &out, "quoteExactInputSingle", token0, token1,
		fee, amountIn, sqrtPriceLimitX96)
	if err != nil {
		log.Fatal(err)
	}
	// quoterContract.QuoteExactInputSingle()
	fmt.Println(out...)
}

// swap,底层是调用SwapRouter
func Swap() {
	client := ec

	wallet := helper.InitWallet("19935d89cb5c67657c64a6383d601e30f04eb179a0369227403e5343bba22107")
	if wallet == nil {
		log.Fatal("init wallet failed")
	}

	pool, err := helper.ConstructV3Pool(client, helper.WMATIC, helper.AMP, uint64(constants.FeeMedium))
	if err != nil {
		log.Fatal(err)
	}

	// 最大滑点0.01%
	slippageTolerance := coreEntities.NewPercent(big.NewInt(1), big.NewInt(1000))
	//after 5 minutes
	d := time.Now().Add(time.Minute * time.Duration(15)).Unix()
	deadline := big.NewInt(d)

	// single trade input
	// single-hop exact input
	r, err := entities.NewRoute([]*entities.Pool{pool}, helper.WMATIC, helper.AMP)
	if err != nil {
		log.Fatal(err)
	}

	swapValue := helper.FloatStringToBigInt("0.1", 18)
	trade, err := entities.FromRoute(r, coreEntities.FromRawAmount(helper.WMATIC, swapValue), coreEntities.ExactInput)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%v %v\n", trade.Swaps[0].InputAmount.Quotient(), trade.Swaps[0].OutputAmount.Wrapped().Quotient())
	params, err := periphery.SwapCallParameters([]*entities.Trade{trade}, &periphery.SwapOptions{
		SlippageTolerance: slippageTolerance,
		Recipient:         wallet.PublicKey,
		Deadline:          deadline,
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("calldata = 0x%x\n", params.Value.String())

	tx, err := helper.TryTX(client, common.HexToAddress(helper.ContractV3SwapRouterV1),
		swapValue, params.Calldata, wallet)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(tx.Hash().String())
}

// 新增流动性
func mintOrAdd(tokenID *big.Int) {
	wallet := helper.InitWallet("19935d89cb5c67657c64a6383d601e30f04eb179a0369227403e5343bba22107")
	if wallet == nil {
		log.Fatal("init wallet failed")
	}

	pool, err := helper.ConstructV3Pool(ec, helper.WMATIC, helper.AMP, uint64(constants.FeeMedium))
	if err != nil {
		log.Fatal("create pool ", err)
	}

	//0.1 MATIC
	amount0 := helper.IntWithDecimal(1, 17)
	amount1 := helper.FloatStringToBigInt("5", 18)
	pos, err := entities.FromAmounts(pool, -43260, 29400, amount0, amount1, false)
	if err != nil {
		log.Fatal(err)
	}

	onePercent := coreEntities.NewPercent(big.NewInt(1), big.NewInt(100))
	log.Println(pos.MintAmountsWithSlippage(onePercent))

	d := time.Now().Add(time.Minute * time.Duration(15)).Unix()
	deadline := big.NewInt(d)

	var opts *periphery.AddLiquidityOptions
	if tokenID == nil {
		//mint a new liquidity position
		opts = &periphery.AddLiquidityOptions{
			CommonAddLiquidityOptions: &periphery.CommonAddLiquidityOptions{
				SlippageTolerance: onePercent,
				Deadline:          deadline,
			},
			MintSpecificOptions: &periphery.MintSpecificOptions{
				Recipient:  wallet.PublicKey,
				CreatePool: true,
			},
		}
	} else {
		//add liquidity to an existing pool
		opts = &periphery.AddLiquidityOptions{
			IncreaseSpecificOptions: &periphery.IncreaseSpecificOptions{
				TokenID: tokenID,
			},
			CommonAddLiquidityOptions: &periphery.CommonAddLiquidityOptions{
				SlippageTolerance: onePercent,
				Deadline:          deadline,
			},
		}

	}
	params, err := periphery.AddCallParameters(pos, opts)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("0x%x  value=%s\n", params.Calldata, params.Value.String())

	//matic is a native token, so we need to set the actually value to transfer
	tx, err := helper.TryTX(ec, common.HexToAddress(helper.ContractV3NFTPositionManager),
		amount0, params.Calldata, wallet)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(tx.Hash().String())
}

func remove(client *ethclient.Client, wallet *helper.Wallet, tokenID *big.Int) {
	//our pool is the fee medium pool
	pool, err := helper.ConstructV3Pool(client, helper.WMATIC, helper.AMP, uint64(constants.FeeMedium))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("liquidity= ", pool.Liquidity)

	posManager, err := contract.NewUniswapv3NFTPositionManager(common.HexToAddress(helper.ContractV3NFTPositionManager), client)
	if err != nil {
		log.Fatal(err)
	}
	contractPos, err := posManager.Positions(nil, tokenID)
	if err != nil {
		log.Fatal(err)
	}
	percent25 := coreEntities.NewPercent(big.NewInt(1), big.NewInt(25))
	fullPercent := coreEntities.NewPercent(contractPos.Liquidity, big.NewInt(1))
	removingLiquidity := fullPercent.Multiply(percent25)

	pos, err := entities.NewPosition(pool, removingLiquidity.Quotient(),
		int(contractPos.TickLower.Int64()),
		int(contractPos.TickUpper.Int64()),
	)
	if err != nil {
		log.Fatal(err)
	}

	d := time.Now().Add(time.Minute * time.Duration(15)).Unix()
	deadline := big.NewInt(d)
	opts := &periphery.RemoveLiquidityOptions{
		TokenID:             tokenID,
		LiquidityPercentage: percent25,
		SlippageTolerance:   coreEntities.NewPercent(big.NewInt(1), big.NewInt(100)), //%1  ,
		Deadline:            deadline,
		CollectOptions: &periphery.CollectOptions{
			ExpectedCurrencyOwed0: coreEntities.FromRawAmount(helper.AMP, big.NewInt(0)),
			ExpectedCurrencyOwed1: coreEntities.FromRawAmount(helper.WMATIC, big.NewInt(0)),
			Recipient:             wallet.PublicKey,
		},
	}
	params, err := periphery.RemoveCallParameters(pos, opts)
	if err != nil {
		log.Fatal(err)
	}

	//matic is a native token, so we need to set the actually value to transfer
	tx, err := helper.TryTX(client, common.HexToAddress(helper.ContractV3NFTPositionManager),
		big.NewInt(0), params.Calldata, wallet)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(tx.Hash().String())
}

func burn(client *ethclient.Client, wallet *helper.Wallet, tokenID *big.Int) {
	ABI, _ := abi.JSON(strings.NewReader(contract.Uniswapv3NFTPositionManagerABI))
	out, err := ABI.Pack("burn", tokenID)
	if err != nil {
		log.Fatal(err)
	}

	//matic is a native token, so we need to set the actually value to transfer
	tx, err := helper.TryTX(client, common.HexToAddress(helper.ContractV3NFTPositionManager),
		big.NewInt(0), out, wallet)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(tx.Hash().String())
}
