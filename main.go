package main

import (
	"letcode/bitcoin"
)

func main(){
	//res := bitcoin.GetSHA256HashCode([]byte("aaa"))
	//pri,pub := bitcoin.GenRsaKey()
	//data := "aadfadf";
	//sign := bitcoin.RsaSignWithSha256(data,pri)
	//res := bitcoin.RsaVerySignWithSha256(data,sign,pub)
	bc := bitcoin.CreateBlockChain()
	bc.AddBlock("add block1")
	bc.AddBlock("add block2")
	bc.Print()
}
