package main

import "sync"

type Block struct {
	Height uint64
	lock   sync.Mutex
}

func main() {
	var b Block
	BadFunc(&b)
	print(b.Height)
}

func BadFunc(b *Block) {
	res := MakeData()
	b.Height = res.Height
}

func MakeData() *Block {
	var nb Block
	nb.Height = 11
	return &nb
}
