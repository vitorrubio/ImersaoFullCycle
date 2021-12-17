package process_transaction

import (
	"github.com/vitorrubio/imersao_fullcycle_desafio1/domain/entity"
	"github.com/vitorrubio/imersao_fullcycle_desafio1/domain/repository"
)

type ProcessTransaction struct {
	repository repository.ITransactionRepository
}

func NewProcessTransaction(repository repository.ITransactionRepository) *ProcessTransaction {
	return &ProcessTransaction{repository: repository}
}

func (p *ProcessTransaction) Execute(input TransactionDtoInput) error {

	transaction := entity.NewTransaction()
	transaction.ID = input.ID
	transaction.AccountID = input.AccountID
	transaction.Amount = input.Amount

	err := p.repository.Insert(
		transaction.ID,
		transaction.AccountID,
		transaction.Amount)

	return err
}
