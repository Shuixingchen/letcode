package tokens

import (
	"letcode/contract/artificial/erc20"
	"letcode/crawler/models"
	"math/big"

	log "github.com/sirupsen/logrus"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	ec *ethclient.Client
)

func init() {
	var err error
	ec, err = ethclient.Dial("https://mainnet.infura.io/v3/40b043c639b44d72966d3535d523a4b3")
	if err != nil {
		log.Fatal(err)
	}
}

// 与erc20合约交互
func QueryERC20(addr string) *models.Token {
	contractAddr := common.HexToAddress(addr)
	var token = new(models.Token)
	var err error
	// 加载智能合约
	tc, err := erc20.NewErc20(contractAddr, ec)
	if err != nil {
		log.Fatal(err)
	}

	token.Addr = addr
	token.Name, err = tc.Name(nil)
	if err != nil {
		log.WithFields(log.Fields{"err": err}).Error("unable to get the erc 20 token name")
	}
	token.Symbol, err = tc.Symbol(nil)
	if err != nil {
		log.WithFields(log.Fields{"err": err}).Error("unable to get the symbol of the erc20 token")
	}
	decimals, err := tc.Decimals(nil)
	if err != nil {
		log.WithFields(log.Fields{"err": err}).Error("unable to get the decimals of the erc 20 token")
	}
	token.Decimals = int64(decimals)
	supply, err := tc.TotalSupply(nil)
	if err != nil || supply == nil {
		log.WithFields(log.Fields{"err": err}).Error("unable to get the total supply of the erc 20 token")
	}
	if supply == nil {
		supply = new(big.Int)
	}
	token.InitTotalSupply = supply
	return token
}
