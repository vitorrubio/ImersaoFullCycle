package factory

import (
	"database/sql"

	repo "github.com/vitorrubio/imersao_fullcycle_desafio1/adapter/repository"
	"github.com/vitorrubio/imersao_fullcycle_desafio1/domain/repository"
)

type RepositoryDatabaseFactory struct {
	DB *sql.DB
}

func NewRepositoryDatabaseFactory(db *sql.DB) *RepositoryDatabaseFactory {
	return &RepositoryDatabaseFactory{DB: db}
}

func (r RepositoryDatabaseFactory) CreateTransactionRepository() repository.ITransactionRepository {
	return repo.NewTransactionRepositoryDb(r.DB)
}
