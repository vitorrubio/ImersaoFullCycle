package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransaction_IsValid(t *testing.T) {

	transaction := NewTransaction()

	cc, invalidCC := NewCreditCard("4193523830170205", "Jose da Silva", 12, 2024, 123)
	transaction.SetCreditCard(*cc)

	transaction.ID = "1"
	transaction.AccountID = "1"
	transaction.Amount = 900
	assert.Nil(t, transaction.IsValid())
	assert.Nil(t, invalidCC)
}

func TestTransaction_IsNotValidWithAmountGreaterThan1000(t *testing.T) {
	transaction := NewTransaction()

	cc, invalidCC := NewCreditCard("4193523830170205", "Jose da Silva", 12, 2024, 123)
	transaction.SetCreditCard(*cc)

	transaction.ID = "1"
	transaction.AccountID = "1"
	transaction.Amount = 1001
	err := transaction.IsValid()
	assert.Error(t, err)
	assert.Equal(t, "you dont have limit for this transaction", err.Error())
	assert.Nil(t, invalidCC)
}

func TestTransaction_IsNotValidWithAmountLesserThan1(t *testing.T) {
	transaction := NewTransaction()

	cc, invalidCC := NewCreditCard("4193523830170205", "Jose da Silva", 12, 2024, 123)
	transaction.SetCreditCard(*cc)

	transaction.ID = "1"
	transaction.AccountID = "1"
	transaction.Amount = 0.5
	err := transaction.IsValid()
	assert.Error(t, err)
	assert.Equal(t, "the amount must be greater than 1", err.Error())
	assert.Nil(t, invalidCC)
}

func TestTransaction_InvalidCc(t *testing.T) {

	transaction := NewTransaction()

	//na verdade não precisaria disso, de invalidCC != nil já pode terminar o teste aqui
	cc, _ := NewCreditCard("1234567890123456", "Jose da Silva", 12, 2024, 123)

	if cc != nil {
		transaction.SetCreditCard(*cc) //se não fizer isso dá erro aqui, não pode pegar o valor de um ponteiro nil
	}

	//pode ser feito o teste sem adicionar cc também

	transaction.ID = "1"
	transaction.AccountID = "1"
	transaction.Amount = 900
	err := transaction.IsValid()
	assert.Error(t, err)
	assert.Equal(t, "invalid credit card number", err.Error())

}
