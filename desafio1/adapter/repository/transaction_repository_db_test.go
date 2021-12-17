package repository

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/vitorrubio/imersao_fullcycle_desafio1/adapter/repository/fixture"
)

func TestTransactionRepositoryDb_Insert(t *testing.T) {
	migrationsDir := os.DirFS("fixture/sql")
	db := fixture.Up(migrationsDir)
	defer fixture.Down(db, migrationsDir)

	repository := NewTransactionRepositoryDb(db)
	err := repository.Insert("1", "1", 2)
	assert.Nil(t, err)
}
