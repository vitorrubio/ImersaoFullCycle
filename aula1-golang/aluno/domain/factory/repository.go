package factory

import "github.com/vitorrubio/imersao5/gateway/domain/repository"

type RepositoryFactory interface {
	CreateTransactionRepository() repository.TransactionRepository
}
