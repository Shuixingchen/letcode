package bitcoin

import (
	"encoding/hex"
	"fmt"
	"log"
)

const(
	bonus = 10
)

type Transaction struct {
	ID   []byte
	Vin  []TXInput
	Vout []TXOutput
}

type TXInput struct{
	Txid      []byte
	Vout      int
	ScriptSig string
}
type TXOutput struct {
	Value        int
	ScriptPubKey string
}

//创建一个铸币交易,不需要输入，输出把奖励写到地址即可
func NewCoinbaseTX(to, data string) *Transaction {
	if data == "" {
		data = fmt.Sprintf("Reward to '%s'", to)
	}
	txin := TXInput{[]byte{}, -1, data}
	txout := TXOutput{bonus, to}
	tx := Transaction{nil, []TXInput{txin}, []TXOutput{txout}}
	tx.SetID()
	return &tx
}

func (in *TXInput) CanUnlockOutputWith(unlockingData string) bool {
	return in.ScriptSig == unlockingData
}
func (out *TXOutput) CanBeUnlockedWith(unlockingData string) bool {
	return out.ScriptPubKey == unlockingData
}

func NewUTXOTransaction(from, to string, amount int, bc *BlockChain) *Transaction {
	var outputs []TXOutput
	var inputs []TXInput

	acc, validOutputs := bc.FindSpendableOutputs(from, amount)

	if acc < amount {
		log.Panic("ERROR: Not enough funds")
	}
	// Build a list of inputs
	for txid, outs := range validOutputs {
		txID, _ := hex.DecodeString(txid)

		for _, out := range outs {
			input := TXInput{txID, out, from}
			inputs = append(inputs, input)
		}
	}

	// Build a list of outputs
	outputs = append(outputs, TXOutput{amount, to})
	if acc > amount {
		outputs = append(outputs, TXOutput{acc - amount, from}) // a change
	}

	tx := Transaction{nil, inputs, outputs}
	tx.SetID()

	return &tx
}

func (tx *Transaction) SetID() {
	tx.ID = []byte(GetSHA256HashCode(tx))
}

func (tx *Transaction) IsCoinbase() bool{
	if len(tx.Vin) == 0 {
		return true
	} else{
		return false
	}
}