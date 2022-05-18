package handler

import (
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/btccom/gokit/explorer/ethereum/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	log "github.com/sirupsen/logrus"
)

/*
  bytes32 label = keccak256(bytes(name));
 uint256 tokenId = uint256(label);
*/
func GetENSTokenID(name string) {
	hash := crypto.Keccak256Hash([]byte(name))
	fmt.Printf("name=%s; tokenId=%d", name, hash.Big())
}

func ParseENSRegister(txHash string) {
	txJson := []byte(`{"jsonrpc":"2.0","method":"eth_getTransactionByHash","params":["` + txHash + `"],"id":1}`)
	txRes := Fetch(txJson)
	var tx Transaction
	if err := json.Unmarshal([]byte(txRes), &tx); err != nil {
		log.WithFields(log.Fields{"method": "eth_getTransactionByHash", "params": txHash}).Panic(err)
	}
	params, err := ParseEnsParams(tx.Input)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(params)
}

func ParseEnsParams(input string) (*types.ENSRegisterParams, error) {
	var ensParams types.ENSRegisterParams
	params, err := types.ParseENSTx(input)
	if err != nil {
		return nil, err
	}
	ensParams.Name = params["name"].(string)
	ensParams.Owner = params["owner"].(common.Address)
	ensParams.Resolver = params["resolver"].(common.Address)
	ensParams.Duration = params["duration"].(*big.Int)
	ensParams.Secret = params["secret"].([32]byte)
	return &ensParams, nil
}
