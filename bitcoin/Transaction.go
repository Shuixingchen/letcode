package bitcoin

type Transaction struct {
	From string
	To string
	Amount float64
}

//输入
type TxIn struct {

}

func CreateTransaction(from string, to string, amount float64) *Transaction{
	return &Transaction{
		From:from,
		To:to,
		Amount:amount,
	}
}