package test

import (
	"fmt"
	"letcode/bitcoin"
	"testing"
)

func TestCreateBitcoin(t *testing.T) {
	//生成一堆公私钥
	pri,pub := bitcoin.GenRsaKey()

	//创建一个区块链，并且生成一个创世块
	bc := bitcoin.CreateBlockChain("from")

	//产生交易，并且验证交易
	txone,_ := bitcoin.NewUTXOTransaction("to",1, pub,pri,bc)
	transactions := []*bitcoin.Transaction{txone}

	//完成工作量证明，并且创建区块，加入到链中
	bc.AddBlock(transactions)

	bc.Print()
}

func TestUTXO(t *testing.T){
	pri,pub := bitcoin.GenRsaKey()
	bc := bitcoin.CreateBlockChain("from")
	txone,_ := bitcoin.NewUTXOTransaction("to",1, pub,pri,bc)
	txtwo,_ := bitcoin.NewUTXOTransaction("to",1, pub,pri,bc)
	ts := []*bitcoin.Transaction{txone,txtwo}
	bc.AddBlock(ts)
	bc.FindUTXO().Print()
}



func TestRsa(t *testing.T) {
	pri,pub := bitcoin.GenRsaKey()
	data := "aadfadf";
	sign := bitcoin.RsaSignWithSha256(data,pri)
	res := bitcoin.RsaVerySignWithSha256(data,sign,pub)
	fmt.Print(res)
}