package bitcoin

import (
	"encoding/hex"
	"fmt"
)

/*
UTXO用一个map保存
transactionId => [outids]=>TXOutput
*/
type UTXO map[string]map[int]TXOutput

/*
从UTXO找到可用output
*/
func (utxo UTXO)FindSpendableOutputs(address string, amount int)  (int, map[string][]int){
	unspentOutputs := make(map[string][]int)
	accumulated := 0
	for txID,outmap := range utxo {
		for outkey, out := range outmap {
			if out.CanBeUnlockedWith(address) {
				accumulated += out.Value
				unspentOutputs[txID] = append(unspentOutputs[txID],outkey)
			}
		}
	}
	return accumulated,unspentOutputs
}

/*
每次产生一个交易，更新utxo
*/
func (utxo UTXO)Update(tx *Transaction) {
	if tx.IsCoinbase() == false {
		tid := hex.EncodeToString(tx.ID)
		//把output写入utxo
		unusedout := make(map[int]TXOutput)
		for outkey, out := range tx.Vout {
			unusedout[outkey] = out
		}
		utxo[tid] = unusedout
		//从utxo扣除已经使用的
		for _,vin := range tx.Vin {
			txID := hex.EncodeToString(vin.Txid)
			if outmap,ok := utxo[txID]; ok {
				if _,k := outmap[vin.Voutkey]; k {
					delete(utxo[txID], vin.Voutkey)
					if len(utxo[txID]) == 0 {
						delete(utxo, txID)
					}
				}
			}
		}
	}
}

func (utxo UTXO)Print() {
	for txID,outputkeys := range utxo {
		fmt.Printf("txID:%s\n",txID)
		fmt.Println(outputkeys)
	}
}