package bitcoin

import (
	"fmt"
	"time"
)

type Block struct {
	Hash string
	PreHash string
	Data string
	time int64

}

type BlockChain struct {
	blocks []*Block
}

func CreateBlock(data string, preHash string) *Block{
	newBlock := Block{
		PreHash:preHash,
		Data:data,
		time:time.Now().Unix(),
	}
	newBlock.Hash = GetSHA256HashCode(newBlock)
	return &newBlock
}

func CreateBlockChain() *BlockChain{
	genesis := CreateBlock("genesis block", "")
	return &BlockChain{[]*Block{genesis}}
}

func (b *BlockChain)AddBlock(data string){
	preBlock := b.blocks[len(b.blocks)-1]
	block := CreateBlock(data, preBlock.Hash)
	b.blocks = append(b.blocks,block)
}
func (b *BlockChain)Print(){
	for _,block := range b.blocks {
		fmt.Printf("pre hash:%s\n",block.PreHash)
		fmt.Printf("data:%s\n",block.Data)
		fmt.Printf("hash:%s\n",block.Hash)
		fmt.Println()
	}
}
