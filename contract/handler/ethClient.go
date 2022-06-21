package handler

import (
	"log"

	gokitclient "github.com/btccom/gokit/explorer/ethereum/ethclient"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	ec  *ethclient.Client
	ecw *ethclient.Client
	gec *gokitclient.Client
)

func init() {
	var err error
	ec, err = ethclient.Dial("https://polygon-mumbai.g.alchemy.com/v2/_4xDtlTKWmynPDVaX1JfRvysRif0wZ85")
	if err != nil {
		log.Fatal(err)
	}
	ecw, err = ethclient.Dial("wss://polygon-mumbai.g.alchemy.com/v2/_4xDtlTKWmynPDVaX1JfRvysRif0wZ85")
	if err != nil {
		log.Fatal(err)
	}

	gec, err = gokitclient.NewClient("wss://mainnet.infura.io/ws/v3/40b043c639b44d72966d3535d523a4b3")
	if err != nil {
		log.Fatal(err)
	}
}
