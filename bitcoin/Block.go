package bitcoin

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

type Block struct {
	Hash []byte
	PreHash []byte
	Transactions []*Transaction
	Timestamp int64
	targetBits int
	Noce int
}

type BlockChain struct {
	blocks []*Block
}

func CreateBlock(data []*Transaction, preHash []byte) *Block{
	newBlock := &Block{
		PreHash:preHash,
		Transactions:data,
		Timestamp:time.Now().Unix(),
	}
	//工作量证明
	pow := NewPoW(newBlock,24)
	noce,hash := pow.Run()
	newBlock.Hash = hash
	newBlock.Noce = noce
	return newBlock
}

func (b *Block) HashTransactions() []byte {
	var txHashes [][]byte
	var txHash [32]byte

	for _, tx := range b.Transactions {
		txHashes = append(txHashes, tx.ID)
	}
	txHash = sha256.Sum256(bytes.Join(txHashes, []byte{}))

	return txHash[:]
}

//create
func CreateBlockChain(address string) *BlockChain{
	coinBase := NewCoinbaseTX(address, "")
	genesis := CreateBlock([]*Transaction{coinBase}, []byte(""))
	return &BlockChain{[]*Block{genesis}}
}

func (b *BlockChain)AddBlock(data []*Transaction){
	preBlock := b.blocks[len(b.blocks)-1]
	block := CreateBlock(data, preBlock.Hash)
	b.blocks = append(b.blocks,block)
}

func (bc *BlockChain) FindSpendableOutputs(address string, amount int) (int, map[string][]int) {
	unspentOutputs := make(map[string][]int)
	unspentTXs := bc.FindUnspentTransactions(address)
	accumulated := 0

Work:
	for _, tx := range unspentTXs {
		txID := hex.EncodeToString(tx.ID)

		for outIdx, out := range tx.Vout {
			if out.CanBeUnlockedWith(address) && accumulated < amount {
				accumulated += out.Value
				unspentOutputs[txID] = append(unspentOutputs[txID], outIdx)

				if accumulated >= amount {
					break Work
				}
			}
		}
	}

	return accumulated, unspentOutputs
}

func (bc *BlockChain) FindUnspentTransactions(address string) []Transaction{
	var unspentTXs []Transaction
	spentTXOs := make(map[string][]int)

	for _,block := range bc.blocks {
		for _,tx := range block.Transactions {
			txID := hex.EncodeToString(tx.ID)
			Outputs:
				for outIdx, out := range tx.Vout {
					// Was the output spent?
					if spentTXOs[txID] != nil {
						for _, spentOut := range spentTXOs[txID] {
							if spentOut == outIdx {
								continue Outputs
							}
						}
					}
					if out.CanBeUnlockedWith(address) {
						unspentTXs = append(unspentTXs, *tx)
					}
				}

				if tx.IsCoinbase() == false {
					for _, in := range tx.Vin {
						if in.CanUnlockOutputWith(address) {
							inTxID := hex.EncodeToString(in.Txid)
							spentTXOs[inTxID] = append(spentTXOs[inTxID], in.Vout)
						}
					}
				}
			}
	}
	return unspentTXs
}



func (b *BlockChain)Print(){
	for _,block := range b.blocks {
		fmt.Printf("pre hash:%x\n",block.PreHash)
		fmt.Printf("data:%s\n",block.HashTransactions())
		fmt.Printf("hash:%x\n",block.Hash)
		fmt.Println()
	}
}
