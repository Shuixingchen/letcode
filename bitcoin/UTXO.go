package bitcoin

import "fmt"

/*
UTXO用一个map保存
transactionId => [outids]=>TXOutput
*/
type UTXO map[string]map[int]TXOutput

/*
新增一个block,需要更新UTXO
*/
func (utxo UTXO)Update(block *Block){
	transactions := block.Transactions
	for _,tx := range transactions {
		if tx.IsCoinbase() == false {
			for _,vin := range tx.Vin {

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