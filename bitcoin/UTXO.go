package bitcoin

import "fmt"

/*
UTXO用一个map保存
transactionId => []TXOutput
*/
type UTXO map[string][]int

func (utxo UTXO)Print() {
	for txID,outputkeys := range utxo {
		fmt.Printf("txID:%s\n",txID)
		fmt.Println(outputkeys)
	}
}