package process_transaction

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/golang/mock/gomock"

	mock_entity "github.com/vitorrubio/esquentaFC-golang/imersao-esquenta-go/entity/mock"
)

func TestPrpocessTransactionWhenItIsValid(t *testing.T) {

	input := TransactionDtoInput{
		ID:        "1",
		AccountID: "1",
		Amount:    200,
	}

	expected := TransactionDtoOutput{
		ID:            "1",
		Status:        "approved",
		ErrrorMessage: "",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repositoryMock := mock_entity.NewMockTransactionRepository(ctrl)
	repositoryMock.EXPECT().Insert(input.ID, input.AccountID, input.Amount, "approved", "").Return(nil)

	usecase := NewProcessTransaction(repositoryMock)
	output, err := usecase.Execute(input)
	assert.Nil(t, err)
	assert.Equal(t, expected, output)

}
