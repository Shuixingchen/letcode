package handler

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	addressHex = "0xe725D38CC421dF145fEFf6eB9Ec31602f95D8097"                       // 地址
	privateHex = "19935d89cb5c67657c64a6383d601e30f04eb179a0369227403e5343bba22107" // 地址对应的私钥
)

// 使用go构造一个tx, 交易触发合约的setA(uint256)方法
//tx.Data 包含字符串，

func SendTx() {
	// 创建交易
	sigTransaction, err := singTx()
	jsonSigTx, _ := sigTransaction.MarshalJSON()
	fmt.Println("sigTransaction: ", string(jsonSigTx))

	// 四、发送交易
	err = ec.SendTransaction(context.Background(), sigTransaction)
	if err != nil {
		fmt.Println("ethClient.SendTransaction failed: ", err.Error())
		return
	}
	fmt.Println("send transaction success,tx: ", sigTransaction.Hash().Hex())
}

func singTx() (*types.Transaction, error) {
	// 一、ABI编码请求参数
	methodId := crypto.Keccak256([]byte("setA(uint256)"))[:4] // 只要前4个字节
	paramValue := math.U256Bytes(new(big.Int).Set(big.NewInt(123)))
	input := append(methodId, paramValue...)
	fmt.Println("input: ", common.Bytes2Hex(input))

	// 二、构造交易对象
	nonce, _ := ec.NonceAt(context.Background(), common.HexToAddress(addressHex), nil)
	gasPrice, _ := ec.SuggestGasPrice(context.Background())
	value := big.NewInt(0)
	gasLimit := uint64(3000000)
	rawTx := types.NewTransaction(nonce, common.HexToAddress(addressHex), value, gasLimit, gasPrice, input)
	jsonRawTx, _ := rawTx.MarshalJSON()
	fmt.Println("rawTx: ", string(jsonRawTx))

	// 三、交易签名
	signer := types.NewEIP155Signer(big.NewInt(1))
	key, err := crypto.HexToECDSA(privateHex)
	if err != nil {
		fmt.Println("crypto.HexToECDSA failed: ", err.Error())
		return nil, err
	}
	sigTransaction, err := types.SignTx(rawTx, signer, key)
	if err != nil {
		fmt.Println("types.SignTx failed: ", err.Error())
		return nil, err
	}
	return sigTransaction, nil
}

// 验证交易，比较签名恢复得到的地址与tx.sender
func ParseAddressFromSigTx() {
	sigTransaction, err := singTx()
	if err != nil {
		fmt.Println("types.SignTx failed: ", err.Error())
	}
	// 使用相同的签名器，eip155后修改了签名器
	signer := types.NewEIP155Signer(big.NewInt(1))
	senderAddr, err := signer.Sender(sigTransaction)

	fmt.Println("sender: ", senderAddr.Hex())
	fmt.Println("addressHex: ", addressHex)
}
