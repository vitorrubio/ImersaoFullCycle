package process_transaction

import (
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	mock_broker "github.com/vitorrubio/imersao5/gateway/adapter/broker/kafka/mock"
	"github.com/vitorrubio/imersao5/gateway/domain/entity"
	mock_repository "github.com/vitorrubio/imersao5/gateway/domain/repository/mock"
)

func TestProcessTransaction_ExecuteInvalidCreditCard(t *testing.T) {
	input := TransactionDtoInput{
		ID:                        "1",
		AccountID:                 "1",
		CreditCardNumber:          "1234567890123456",
		CreditCardName:            "John Doe",
		CreditCardExpirationMonth: 12,
		CreditCardExpirationYear:  time.Now().Year(),
		CreditCardCVV:             123,
		Amount:                    200,
	}

	expectedOutput := TransactionDtoOutput{
		ID:           "1",
		Status:       entity.REJECTED,
		ErrorMessage: "invalid credit card number",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repositoryMoc := mock_repository.NewMockTransactionRepository(ctrl)
	repositoryMoc.EXPECT().Insert(input.ID, input.AccountID, input.Amount, expectedOutput.Status, expectedOutput.ErrorMessage).
		Return(nil)

	producerMock := mock_broker.NewMockProducerInterface(ctrl)
	producerMock.EXPECT().Publish(expectedOutput, []byte(input.ID), "transactions_result")

	usecase := NewProcessTransaction(repositoryMoc, producerMock, "transactions_result")

	output, err := usecase.Execute(input)
	assert.Error(t, err)
	assert.Equal(t, expectedOutput, output)

}

func TestProcessTransaction_ExecuteRejectedTransaction(t *testing.T) {
	input := TransactionDtoInput{
		ID:                        "1",
		AccountID:                 "1",
		CreditCardNumber:          "4193523830170205",
		CreditCardName:            "John Doe",
		CreditCardExpirationMonth: 12,
		CreditCardExpirationYear:  time.Now().Year(),
		CreditCardCVV:             123,
		Amount:                    1200,
	}

	expectedOutput := TransactionDtoOutput{
		ID:           "1",
		Status:       entity.REJECTED,
		ErrorMessage: "you dont have limit for this transaction",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repositoryMoc := mock_repository.NewMockTransactionRepository(ctrl)
	repositoryMoc.EXPECT().Insert(input.ID, input.AccountID, input.Amount, expectedOutput.Status, expectedOutput.ErrorMessage).
		Return(nil)

	producerMock := mock_broker.NewMockProducerInterface(ctrl)
	producerMock.EXPECT().Publish(expectedOutput, []byte(input.ID), "transactions_result")

	usecase := NewProcessTransaction(repositoryMoc, producerMock, "transactions_result")

	output, err := usecase.Execute(input)
	assert.Error(t, err)
	assert.Equal(t, expectedOutput, output)

}

func TestProcessTransaction_ExecuteApprovedTransaction(t *testing.T) {
	input := TransactionDtoInput{
		ID:                        "1",
		AccountID:                 "1",
		CreditCardNumber:          "4193523830170205",
		CreditCardName:            "John Doe",
		CreditCardExpirationMonth: 12,
		CreditCardExpirationYear:  time.Now().Year(),
		CreditCardCVV:             123,
		Amount:                    900,
	}

	expectedOutput := TransactionDtoOutput{
		ID:           "1",
		Status:       entity.APPROVED,
		ErrorMessage: "",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repositoryMoc := mock_repository.NewMockTransactionRepository(ctrl)
	repositoryMoc.EXPECT().Insert(input.ID, input.AccountID, input.Amount, expectedOutput.Status, expectedOutput.ErrorMessage).
		Return(nil)

	producerMock := mock_broker.NewMockProducerInterface(ctrl)
	producerMock.EXPECT().Publish(expectedOutput, []byte(input.ID), "transactions_result")

	usecase := NewProcessTransaction(repositoryMoc, producerMock, "transactions_result")

	output, err := usecase.Execute(input)
	assert.Nil(t, err)
	assert.Equal(t, expectedOutput, output)

}
