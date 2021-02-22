package bitcoin

import "time"

type Block struct {
	Hash string
	PreHash string
	Data *MTree
	time int64
	target int64
	noce int64
}

func CreateBlock(transactions []*Transaction, noce int64, preHash string) *Block{
	newMtree := CreateMTree(transactions)
	newBlock := Block{
		PreHash:preHash,
		Data:newMtree,
		time:time.Now().Unix(),
		target:1111,
		noce:noce,
	}
	newBlock.Hash = GetSHA256HashCode(newBlock)
	return &newBlock
}
