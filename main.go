package main

import (
	"fmt"
	"letcode/bitcoin"
)

func main(){
	res := bitcoin.GetSHA256HashCode([]byte("aaa"))
	fmt.Println(res)
}
