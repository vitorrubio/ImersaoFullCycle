package factory

import (
	"github.com/vitorrubio/imersao_fullcycle_desafio1/domain/repository"
)

type RepositoryFactory interface {
	CreateTransactionRepository() repository.ITransactionRepository
}
