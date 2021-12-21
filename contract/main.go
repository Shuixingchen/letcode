package main

import (
	"context"
	"letcode/contract/abi"
	"letcode/contract/client"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rpc"
)

var addr = "aaa"

func main() {
	var client client.EthClient
	var err error
	client.Rpc, err = rpc.Dial("https://mainnet.infura.io/v3/40b043c639b44d72966d3535d523a4b3")
	if err != nil {
		panic(err)
	}
	defer client.Close(context.Background())
	contractAddress := common.HexToAddress(addr)

	tokenCall, err := abi.NewToken(contractAddress, client)
	tokenCall.Name(nil)
	tokenCall.Symbol(nil)
	tokenCall.Decimals(nil)
	tokenCall.TotalSupply(nil)
}
