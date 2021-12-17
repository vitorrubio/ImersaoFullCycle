package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/vitorrubio/imersao_fullcycle_desafio1/adapter/factory"
	"github.com/vitorrubio/imersao_fullcycle_desafio1/usecase/process_transaction"
)

func main() {
	fmt.Println("Desafio 1 - Golang - 07/12")

	fmt.Println("Source: https://github.com/vitorrubio/imersao_fullcycle_desafio1.git")

	fmt.Println("****************************************************************************************")

	fmt.Println("")

	db, err := sql.Open("sqlite3", "banco.db")

	if err != nil {
		log.Fatal(err)
	}

	repositoryFactory := factory.NewRepositoryDatabaseFactory(db)

	repository := repositoryFactory.CreateTransactionRepository()

	usecase := process_transaction.NewProcessTransaction(repository)

	input := process_transaction.TransactionDtoInput{
		ID:        "1",
		AccountID: "1",
		Amount:    200,
	}

	err = usecase.Execute(input)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println("Registro incluido com sucesso!")

}
