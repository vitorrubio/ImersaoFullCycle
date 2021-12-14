package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransactionWithAmountGreaterThan1000(t *testing.T) {
	transaction := NewTransaction()
	transaction.ID = "1"
	transaction.AccountID = "1"
	transaction.Amount = 2000
	err := transaction.IsValid()

	assert.Error(t, err)
	assert.Equal(t, "Amount cannot be greater than 1000", err.Error())

}

//"Amount must be greater than 1"

func TestTransactionWithAmountLesserThan1(t *testing.T) {
	transaction := NewTransaction()
	transaction.ID = "1"
	transaction.AccountID = "1"
	transaction.Amount = 0.5
	err := transaction.IsValid()

	assert.Error(t, err)
	assert.Equal(t, "Amount must be greater than 1", err.Error())

}
