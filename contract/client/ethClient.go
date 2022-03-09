package client

import (
	"github.com/btccom/gokit/explorer/ethereum/ethclient"
	log "github.com/sirupsen/logrus"
)

func CreateEthClient() *ethclient.Client {
	var err error
	ec, err := ethclient.NewClient("wss://ropsten.infura.io/ws/v3/40b043c639b44d72966d3535d523a4b3")
	// ec, err := ethclient.NewClient("wss://mainnet.infura.io/ws/v3/40b043c639b44d72966d3535d523a4b3")
	if err != nil {
		log.Fatal(err)
	}
	return ec
}
