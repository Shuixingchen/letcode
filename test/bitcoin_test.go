package test

import (
	"letcode/bitcoin"
	"testing"
)

func TestCreateBitcoin(t *testing.T) {
	//创建一个区块链，并且生成一个创世块
	bc := bitcoin.CreateBlockChain("from")

	//产生交易，并且验证交易
	txone := bitcoin.NewUTXOTransaction("from","to", 1,bc)
	transactions := []*bitcoin.Transaction{txone}

	//完成工作量证明，并且创建区块，加入到链中
	bc.AddBlock(transactions)

	bc.Print()
}

//func TestRsa(t *testing.T) {
//	pri,pub := bitcoin.GenRsaKey()
//	data := "aadfadf";
//	sign := bitcoin.RsaSignWithSha256(data,pri)
//	res := bitcoin.RsaVerySignWithSha256(data,sign,pub)
//
//	fmt.Print(res)
//}