package process_transaction

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	mock_repository "github.com/vitorrubio/imersao_fullcycle_desafio1/domain/repository/mock"
)

func TestProcessTransaction_Execute(t *testing.T) {
	input := TransactionDtoInput{
		ID:        "1",
		AccountID: "1",
		Amount:    200,
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repositoryMoc := mock_repository.NewMockITransactionRepository(ctrl)
	repositoryMoc.EXPECT().Insert(input.ID, input.AccountID, input.Amount).
		Return(nil)

	usecase := NewProcessTransaction(repositoryMoc)

	err := usecase.Execute(input)
	assert.Nil(t, err)
}
