package bitcoin

type MTree struct {
	Hash string
	Transactions []*Transaction
}

func CreateMTree(transactions []*Transaction) *MTree{

	return &MTree{
		Hash:GetSHA256HashCode(transactions),
		Transactions:transactions,
	}
}