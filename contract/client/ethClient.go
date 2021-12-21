package client

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rpc"
)

type EthClient struct {
	Rpc *rpc.Client
}

// CodeAt returns code at a given address.
// Take note: this function is used by TokenContract only. Shall not be used directly.
func (ec *EthClient) CodeAt(ctx context.Context, addr common.Address, block *big.Int) ([]byte, error) {
	var result hexutil.Bytes
	return result, nil
}

// GetCodeLatest returns the latest code at given address.
func (ec *EthClient) GetCodeLatest(ctx context.Context, addr string) (string, error) {
	var result hexutil.Bytes

	return result.String(), nil
}

// GetCodeLatest returns the latest code at given address.
func (ec *EthClient) Close(ctx context.Context) error {
	return ec.Close(ctx)
}

// func (ec *EthClient) CallContract(ctx context.Context, msg ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {
// 	var hex hexutil.Bytes
// 	return hex, nil
// }

// CallContract implements the contractcaller interface
func (ec *EthClient) CallContract(ctx context.Context, msg ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {
	var hex hexutil.Bytes
	// err := ec.CallContext(ctx, &hex, "eth_call", toCallArg(msg), toBlockNumArg(blockNumber))

	return hex, nil
}
