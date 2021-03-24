package bitcoin

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

const(
	bonus = 10
	targetBits = 4
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
	pow := NewPoW(newBlock,targetBits)
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

/*
通过对区块进行迭代找到所有未花费输出
transactionid=>uoutputs
 */

func (bc *BlockChain)FindUTXO() UTXO{
	utxo := make(UTXO)
	spentTXOs := make(map[string][]int)//transactionId=>[output的下标]

	for _,block := range bc.blocks {
		for _,tx := range block.Transactions {
			txID := hex.EncodeToString(tx.ID)
			unsedOutput := make([]int,0)
		Outputs:
			//遍历output,找出没有用的output
			for outIdx, _ := range tx.Vout {
				// Was the output spent?
				if spentTXOs[txID] != nil {
					for _, spentOut := range spentTXOs[txID] {
						if spentOut == outIdx {
							continue Outputs
						}
					}
				}
				//this output was not used
				unsedOutput = append(unsedOutput, outIdx)
			}
			utxo[txID] = unsedOutput

			//遍历input,找出已经使用的output
			if tx.IsCoinbase() == false {
				for _, in := range tx.Vin {
					//find spend output
					inTxID := hex.EncodeToString(in.Txid)
					spentTXOs[inTxID] = append(spentTXOs[inTxID], in.Voutkey)
					//delete has put in utxo
					if outputkeys,ok := utxo[inTxID]; ok{
						for k,outIdx:= range outputkeys {
							if outIdx == in.Voutkey {
								utxo[inTxID] = append(utxo[inTxID][:k], utxo[inTxID][k+1:]...) //从utxo中删除这个outIdx
								if len(utxo[inTxID]) == 0 {
									delete(utxo, inTxID)
								}
							}
						}
					}
				}
			}
		}
	}
	return utxo
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

/*查询没有使用的交易输出，类似与查询某个地址的余额。
区块链没有账号统计，所以需要对整条链遍历，输入表示使用，输出表示得到。
对所有transaction的输出遍历，把属于这个address的transaction写入到unspentTXs;
对所有transaction的输入遍历，把属于这个address已经花的output写入到spentTXOs;
*/

func (bc *BlockChain) FindUnspentTransactions(address string) []Transaction{
	var unspentTXs []Transaction
	spentTXOs := make(map[string][]int)//transactionId=>[output的下标]

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
					if out.CanBeUnlockedWith(address) { //是否属于这个地址的输出
						unspentTXs = append(unspentTXs, *tx)
					}
				}

				if tx.IsCoinbase() == false {
					for _, in := range tx.Vin {
						//find spend output
						if in.CanUnlockOutputWith(address) {
							inTxID := hex.EncodeToString(in.Txid)
							spentTXOs[inTxID] = append(spentTXOs[inTxID], in.Voutkey)
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
		fmt.Printf("createTime:%d\n",block.Timestamp)
		fmt.Println()
	}
}
