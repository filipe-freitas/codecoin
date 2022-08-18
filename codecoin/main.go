package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/filipe-freitas/codecoin/infrastructure/grpc/server"
	"github.com/filipe-freitas/codecoin/infrastructure/kafka"
	"github.com/filipe-freitas/codecoin/infrastructure/repository"
	usecase "github.com/filipe-freitas/codecoin/usecase"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading the .env file")
	}
}

func main() {

	DB := setupDB()
	defer DB.Close()

	producer := setupKafkaProducer()
	processTransactionUseCase := setupTransactionUseCase(DB, producer)

	serveGRPC(processTransactionUseCase)
}

func setupDB() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("host"),
		os.Getenv("port"),
		os.Getenv("user"),
		os.Getenv("password"),
		os.Getenv("dbname"))

	DB, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println(err)
		log.Fatal("Error connecting to the database")
	}

	return DB
}

func setupKafkaProducer() kafka.KafkaProducer {
	producer := kafka.NewKafkaProducer()
	producer.SetupProducer(os.Getenv("KafkaBootstrapServers"))
	return producer
}

func setupTransactionUseCase(DB *sql.DB, producer kafka.KafkaProducer) usecase.UseCaseTransaction {
	transactionRepository := repository.NewTransactionRepositoryDB(DB)
	useCase := usecase.NewUseCaseTransaction(transactionRepository)
	useCase.KafkaProducer = producer
	return useCase
}

func serveGRPC(processTransactionUseCase usecase.UseCaseTransaction) {
	grpcServer := server.NewGRPCServer()
	grpcServer.ProcessTransactionUseCase = processTransactionUseCase
	grpcServer.Serve()
}
