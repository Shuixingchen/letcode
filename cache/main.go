package main

import (
	"fmt"
	cache "letcode/cache/redis"
	"math/big"
)
/*
使用json字符串缓存数据，json
*/
type Account struct {
	Address string `json:"address"`
	Name    string `json:"name"`
	EthBal  HexBig `json:"ethBalance"`
	// TokenBal    map[string]HexBig `json:"-"`
	TxSize      int         `json:"txSize"`
	ERC20TxSize int         `json:"tokenTxSize"`
	ItTxSize    int         `json:"internalTxSize"`
	EventSize   int         `json:"eventSize"`
	MinedBlocks int         `json:"minedBlocks"`
	MinedUncles int         `json:"minedUncles"`
	Type        AccountType `json:"accountType"`
}

type AccountType = int

type HexBig struct {
	*big.Int
}

type Block struct {
	heigh int
}

func redis(){
	redisClient,err := cache.NewRedisClient()
	if err!= nil{
		fmt.Println(err)
	}
	var list []*Account
	for i:=1;i<10;i++{
		b := new(Account)
		list = append(list,b)
	}
	var newList []*Account
	err = redisClient.AddDataToCache("keyee", list)
	if err!=nil{
		fmt.Println("setdata",err)
	}
	err = redisClient.GetDatafromCache("keyee",&newList)
	fmt.Println("getdata:", err)
}


func main(){
	redis()
}