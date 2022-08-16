package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/filipe-freitas/codecoin/domain"
	"github.com/filipe-freitas/codecoin/infrastructure/repository"
	usecase "github.com/filipe-freitas/codecoin/usecase"
	_ "github.com/lib/pq"
)

func main() {

	DB := setupDB()
	defer DB.Close()

	cc := domain.NewCreditCard()
	cc.Number = "1234"
	cc.Name = "Filipe"
	cc.ExpirationMonth = 12
	cc.ExpirationYear = 2022
	cc.CVV = 123
	cc.Limit = 1273.725
	cc.Balance = 100

	repo := repository.NewTransactionRepositoryDB(DB)
	err := repo.CreateCreditCard(*cc)
	if err != nil {
		fmt.Println(err)
	}
}

func setupDB() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		"db",
		"5432",
		"postgres",
		"root",
		"codecoin")

	DB, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println(err)
		log.Fatal("Error connecting to the database")
	}

	return DB
}

func setupTransactionUseCase(DB *sql.DB) usecase.UseCaseTransaction {
	transactionRepository := repository.NewTransactionRepositoryDB(DB)
	useCase := usecase.NewUseCaseTransaction(transactionRepository)
	return useCase
}
