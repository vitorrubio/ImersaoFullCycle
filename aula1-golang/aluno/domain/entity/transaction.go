package entity

import (
	"errors"
)

const (
	REJECTED = "rejected"
	APPROVED = "approved"
)

type Transaction struct {
	ID           string
	AccountID    string
	Amount       float64
	creditCard   CreditCard
	Status       string
	ErrorMessage string
}

func NewTransaction() *Transaction {
	return &Transaction{}
}

func (t *Transaction) IsValid() error {

	if t.Amount > 1000 {
		return errors.New("you dont have limit for this transaction")
	}

	if t.Amount < 1 {
		return errors.New("the amount must be greater than 1")
	}

	validCC := t.creditCard.IsValid()
	if validCC != nil {
		return validCC
	}

	return nil
}

func (t *Transaction) SetCreditCard(creditCard CreditCard) {
	t.creditCard = creditCard
}
