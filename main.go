package main

import (
	"fmt"
	"letcode/bitcoin/db"
)

func main(){
	db := db.CreateDb()
	fmt.Println(db)
	//bc := bitcoin.CreateBlockChain()
	//bc.AddBlock([]byte("add block1"))
	//bc.AddBlock([]byte("add block2"))
	//bc.Print()
}
