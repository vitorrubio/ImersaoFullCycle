package entity

import "errors"

type Transaction struct {
	ID            string
	AccountID     string
	Amount        float64
	Status        string
	ErrrorMessage string
}

func NewTransaction() *Transaction {
	return &Transaction{}
}

func (t *Transaction) IsValid() error {
	if t.Amount > 1000 {
		return errors.New("Amount cannot be greater than 1000")
	}

	if t.Amount < 1 {
		return errors.New("Amount must be greater than 1")
	}

	return nil
}
