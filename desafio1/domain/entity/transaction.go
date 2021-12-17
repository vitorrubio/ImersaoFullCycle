package entity

type Transaction struct {
	ID        string
	AccountID string
	Amount    float64
}

func NewTransaction() *Transaction {
	return &Transaction{}
}
