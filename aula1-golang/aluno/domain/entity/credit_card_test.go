package entity

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreditCardNumber(t *testing.T) {
	_, err := NewCreditCard("1234567890123456", "Jose da Silva", 12, 2024, 123)
	assert.Equal(t, "invalid credit card number", err.Error())

	//_, err = NewCreditCard("4000000000000000", "Jose da Silva", 12, 2024, 123)
	//assert.Equal(t, "invalid credit card number", err.Error())

	_, err = NewCreditCard("4193523830170205", "Jose da Silva", 12, 2024, 123)
	assert.Nil(t, err)
}

func TestCreditCardExpirationMonth(t *testing.T) {
	_, err := NewCreditCard("4193523830170205", "Jose da Silva", 12, 2024, 123)
	assert.Nil(t, err)

	_, err = NewCreditCard("4193523830170205", "Jose da Silva", 13, 2024, 123)
	assert.Equal(t, "invalid credit card expiration month", err.Error())

	_, err = NewCreditCard("4193523830170205", "Jose da Silva", 0, 2024, 123)
	assert.Equal(t, "invalid credit card expiration month", err.Error())
}

func TestCreditCardExpirationYear(t *testing.T) {
	_, err := NewCreditCard("4193523830170205", "Jose da Silva", 12, 2024, 123)
	assert.Nil(t, err)

	_, err = NewCreditCard("4193523830170205", "Jose da Silva", 12, time.Now().Year(), 123)
	assert.Nil(t, err)

	_, err = NewCreditCard("4193523830170205", "Jose da Silva", 12, 2020, 123)
	assert.Equal(t, "invalid credit card expiration year", err.Error())

}

func TestCreditCardExpired(t *testing.T) {
	_, err := NewCreditCard("4193523830170205", "Jose da Silva", 12, 2024, 123)
	assert.Nil(t, err)

	today := time.Now()
	expiredDate := today.AddDate(0, -1, 0)
	_, err = NewCreditCard("4193523830170205", "Jose da Silva", int(expiredDate.Month()), expiredDate.Year(), 123)
	assert.Error(t, err)
	assert.Equal(t, "expired credit card", err.Error())

}
