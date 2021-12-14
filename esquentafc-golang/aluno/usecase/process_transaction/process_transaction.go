package process_transaction

import (
	"github.com/vitorrubio/esquentaFC-golang/imersao-esquenta-go/entity"
)

type ProcessTransaction struct {
	Repository entity.TransactionRepository
}

func NewProcessTransaction(repository entity.TransactionRepository) *ProcessTransaction {
	return &ProcessTransaction{
		Repository: repository,
	}
}

func (p *ProcessTransaction) Execute(input TransactionDtoInput) (TransactionDtoOutput, error) {

	transaction := entity.NewTransaction()
	transaction.Amount = input.Amount
	transaction.AccountID = input.AccountID
	transaction.ID = input.ID
	invalidtransaction := transaction.IsValid()

	if invalidtransaction != nil {
		err := p.Repository.Insert(transaction.ID, transaction.AccountID, transaction.Amount, "rejected", invalidtransaction.Error())

		if err != nil {
			return TransactionDtoOutput{}, err
		}

		output := TransactionDtoOutput{
			ID:            transaction.ID,
			Status:        "rejected",
			ErrrorMessage: invalidtransaction.Error(),
		}

		return output, nil
	}

	err := p.Repository.Insert(transaction.ID, transaction.AccountID, transaction.Amount, "approved", "")

	if err != nil {
		return TransactionDtoOutput{}, err
	}

	output := TransactionDtoOutput{
		ID:            transaction.ID,
		Status:        "approved",
		ErrrorMessage: "",
	}

	return output, nil

}
