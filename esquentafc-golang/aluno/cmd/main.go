package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/vitorrubio/esquentaFC-golang/imersao-esquenta-go/adapter/repository"
	"github.com/vitorrubio/esquentaFC-golang/imersao-esquenta-go/usecase/process_transaction"
)

func main() {
	fmt.Println("Hello, World!")

	db, err := sql.Open("sqlite3", "test.db")

	if err != nil {
		log.Fatal(err)
	}

	repo := repository.NewTransactionRepositoryDb(db)
	usecase := process_transaction.NewProcessTransaction(repo)

	input := process_transaction.TransactionDtoInput{
		ID:        "1",
		AccountID: "1",
		Amount:    0,
	}

	output, err := usecase.Execute(input)

	if err != nil {
		fmt.Println(err.Error())
	}

	outputJson, _ := json.Marshal(output)

	fmt.Println(string(outputJson))

}
