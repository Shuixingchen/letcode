package bitcoin

type Transaction struct {
	From string
	To string
	Amount float64
}

func CreateTransaction(from string, to string, amount float64) *Transaction{
	return &Transaction{
		From:from,
		To:to,
		Amount:amount,
	}
}