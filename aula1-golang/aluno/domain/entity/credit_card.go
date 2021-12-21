package entity

import (
	"errors"
	"regexp"
	"time"
)

type CreditCard struct {
	number          string
	name            string
	expirationMonth int
	expirationYear  int
	cvv             int
}

func NewCreditCard(number string,
	name string,
	expirationMonth int,
	expirationYear int,
	cvv int) (*CreditCard, error) {

	cc := &CreditCard{
		number:          number,
		name:            name,
		expirationMonth: expirationMonth,
		expirationYear:  expirationYear,
		cvv:             cvv,
	}

	err := cc.IsValid()
	if err != nil {
		return nil, err
	}

	return cc, nil
}

func (cc *CreditCard) IsValid() error {

	err := cc.validateNumber()
	if err != nil {
		return err
	}

	err = cc.validateMonth()
	if err != nil {
		return err
	}

	err = cc.validateYear()
	if err != nil {
		return err
	}

	err = cc.validateDate()
	if err != nil {
		return err
	}

	return nil
}

///invalid credit card number
func (cc *CreditCard) validateNumber() error {
	//`^[0-9]{16}$`
	re := regexp.MustCompile(`^(?:4[0-9]{12}(?:[0-9]{3})?|[25][1-7][0-9]{14}|6(?:011|5[0-9][0-9])[0-9]{12}|3[47][0-9]{13}|3(?:0[0-5]|[68][0-9])[0-9]{11}|(?:2131|1800|35\d{3})\d{11})$`)

	if !re.MatchString(cc.number) {
		return errors.New("invalid credit card number")
	}
	return nil
}

///invalid credit card expiration month
func (cc *CreditCard) validateMonth() error {

	if cc.expirationMonth <= 0 || cc.expirationMonth > 12 {
		return errors.New("invalid credit card expiration month")
	}

	return nil
}

///invalid credit card expiration year
func (cc *CreditCard) validateYear() error {
	currentYear := time.Now().Year()
	if cc.expirationYear < currentYear {
		return errors.New("invalid credit card expiration year")
	}

	return nil
}

///expired credit card
func (cc *CreditCard) validateDate() error {
	currentYear := time.Now().Year()
	currentMonth := time.Now().Month()

	if cc.expirationYear == currentYear && cc.expirationMonth < int(currentMonth) {
		return errors.New("expired credit card")
	}

	return nil
}
