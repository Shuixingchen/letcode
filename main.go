package main

import (
	"fmt"
	"letcode/bitcoin"
)


func main(){

	pri,pub := bitcoin.GenRsaKey()
	data := "aadfadf";
	sign := bitcoin.RsaSignWithSha256(data,pri)
	res := bitcoin.RsaVerySignWithSha256(data,sign,pub)

	fmt.Print(res)
}
