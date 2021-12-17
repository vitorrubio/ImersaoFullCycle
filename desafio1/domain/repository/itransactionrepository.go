package repository

type ITransactionRepository interface {
	Insert(id string, account string, amount float64) error
}
