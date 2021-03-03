package bitcoin

import (
	"fmt"
	"time"
)

type Block struct {
	Hash []byte
	PreHash []byte
	Data []byte
	Timestamp int64
	targetBits int
	Noce int
}

type BlockChain struct {
	blocks []*Block
}

func CreateBlock(data []byte, preHash []byte) *Block{
	newBlock := &Block{
		PreHash:preHash,
		Data:data,
		Timestamp:time.Now().Unix(),
	}
	//工作量证明
	pow := NewPoW(newBlock,24)
	noce,hash := pow.Run()
	newBlock.Hash = hash
	newBlock.Noce = noce
	return newBlock
}

func CreateBlockChain() *BlockChain{
	genesis := CreateBlock([]byte("genesis block"), []byte(""))
	return &BlockChain{[]*Block{genesis}}
}

func (b *BlockChain)AddBlock(data []byte){
	preBlock := b.blocks[len(b.blocks)-1]
	block := CreateBlock(data, preBlock.Hash)
	b.blocks = append(b.blocks,block)
}
func (b *BlockChain)Print(){
	for _,block := range b.blocks {
		fmt.Printf("pre hash:%x\n",block.PreHash)
		fmt.Printf("data:%s\n",block.Data)
		fmt.Printf("hash:%x\n",block.Hash)
		fmt.Println()
	}
}
