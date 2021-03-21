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
	Txid      []byte //上个交易的transactionid
	Voutkey      int //上个交易的输出下标
	ScriptSig string //简化验证，直接存付款人的address
}
type TXOutput struct {
	Value        int
	ScriptPubKey string //简化验证，直接存收款人的address
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

//
func (in *TXInput) CanUnlockOutputWith(unlockingData string) bool {
	return in.ScriptSig == unlockingData
}

//解锁这个输出，只有解锁成功，这个输出才能使，暂时用地址来解锁。传进来的地址和输出保存的地址一致，代表成功。
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