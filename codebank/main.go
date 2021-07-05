package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"github.com/hrsniper/imersao-fullstack-fullcycle-3/domain"
	"github.com/hrsniper/imersao-fullstack-fullcycle-3/infrastructure/repository"
	"github.com/hrsniper/imersao-fullstack-fullcycle-3/usecase"
)

func main() {
	db := setupDb()

	defer db.Close()

	creditCard := domain.NewCreditCard()
	creditCard.Number = "1234"
	creditCard.Name = "Wesley"
	creditCard.ExpirationYear = 2021
	creditCard.ExpirationMonth = 7
	creditCard.CVV = 123
	creditCard.Limit = 1000
	creditCard.Balance = 0

	repo := repository.NewTransactionRepositoryDb(db)

	err := repo.CreateCreditCard(*creditCard)

	if err != nil {
		fmt.Println(err)
	}
}

func setupTransactionUseCase(db *sql.DB) usecase.UseCaseTransaction {
	transactionRepository := repository.NewTransactionRepositoryDb(db)

	useCase := usecase.NewUseCaseTransaction(transactionRepository)

	return useCase
}

func setupDb() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		"db",
		"5432",
		"postgres",
		"root",
		"codebank",
	)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatal("error connection to database")
	}

	return db
}
