package bitcoin

import (
	"encoding/hex"
	"errors"
	"fmt"
)


type Transaction struct {
	ID   []byte
	Vin  []TXInput
	Vout []TXOutput
}
type ScriptSig struct {
	Address string //付款人地址
	PubKey []byte //付款人公钥
	Sig string //本次交易签名
}

type TXInput struct{
	Txid      []byte //上个交易的transactionid
	Voutkey      int //上个交易的输出下标
	ScriptSig ScriptSig
}
type TXOutput struct {
	Value        int
	ScriptPubKey string //收款人的address
}

//创建一个铸币交易,不需要输入，输出把奖励写到地址即可
func NewCoinbaseTX(to, data string) *Transaction {
	if data == "" {
		data = fmt.Sprintf("Reward to '%s'", to)
	}
	txin := TXInput{[]byte{}, -1, nil}
	txout := TXOutput{bonus, to}
	tx := Transaction{nil, []TXInput{txin}, []TXOutput{txout}}
	tx.SetID()
	return &tx
}

/*
验证input是否合法

*/
func (in *TXInput) CanUnlockOutputWith(unlockingData string) bool {
	return in.ScriptSig.Sig == unlockingData
}

//解锁这个输出，只有解锁成功，这个输出才能使，暂时用地址来解锁。传进来的地址和输出保存的地址一致，代表成功。
func (out *TXOutput) CanBeUnlockedWith(unlockingData string) bool {
	return out.ScriptPubKey == unlockingData
}

func NewUTXOTransaction(to string, amount int, pubKey []byte, prvKey []byte, bc *BlockChain) (*Transaction, error) {
	var outputs []TXOutput
	var inputs []TXInput
	from := PubKeyToAddress(pubKey)
	acc, validOutputs := bc.UTXO.FindSpendableOutputs(from, amount)

	if acc < amount {
		return nil, errors.New("ERROR: Not enough funds from "+from)
	}
	// Build a list of inputs
	for txid, outmap := range validOutputs {
		txID, _ := hex.DecodeString(txid)

		for outkey, _ := range outmap {
			//对这个input进行签名，确定是from这个人操作的
			scriptsig := ScriptSig{
				Address: from,
				PubKey:  pubKey,
				Sig:     RsaSignWithSha256(txid, prvKey),
			}
			input := TXInput{txID, outkey, scriptsig}
			inputs = append(inputs, input)
		}
	}

	// Build a list of outputs
	outputs = append(outputs, TXOutput{amount, to})
	if acc > amount {
		outputs = append(outputs, TXOutput{acc - amount, from}) // a change
	}

	tx := &Transaction{nil, inputs, outputs}
	tx.SetID()
	bc.UTXO.Update(tx)
	return tx, nil
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