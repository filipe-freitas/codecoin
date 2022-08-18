package usecase

import (
	"encoding/json"
	"os"
	"time"

	"github.com/filipe-freitas/codecoin/domain"
	"github.com/filipe-freitas/codecoin/dto"
	"github.com/filipe-freitas/codecoin/infrastructure/kafka"
)

type UseCaseTransaction struct {
	TransactionRepository domain.TransactionRepository
	KafkaProducer         kafka.KafkaProducer
}

func NewUseCaseTransaction(transactionRepository domain.TransactionRepository) UseCaseTransaction {
	return UseCaseTransaction{TransactionRepository: transactionRepository}
}

func (u UseCaseTransaction) ProcessTransaction(transactionDTO dto.Transaction) (domain.Transaction, error) {
	creditCard := u.hydrateCreditCard(transactionDTO)
	ccBalanceAndLimit, err := u.TransactionRepository.GetCreditCard(*creditCard)
	if err != nil {
		return domain.Transaction{}, err
	}
	creditCard.ID = ccBalanceAndLimit.ID
	creditCard.Balance = ccBalanceAndLimit.Balance
	creditCard.Limit = ccBalanceAndLimit.Limit

	transaction := u.newTransaction(transactionDTO, ccBalanceAndLimit)
	transaction.ProcessAndValidate(creditCard)

	err = u.TransactionRepository.SaveTransaction(*transaction, *creditCard)
	if err != nil {
		return domain.Transaction{}, err
	}

	transactionDTO.ID = transaction.ID
	transactionDTO.CreatedAt = transaction.CreatedAt
	transactionJson, err := json.Marshal(transactionDTO)
	if err != nil {
		return domain.Transaction{}, err
	}

	err = u.KafkaProducer.Publish(string(transactionJson), os.Getenv("KafkaTransactionsTopic"))
	if err != nil {
		return domain.Transaction{}, err
	}

	return *transaction, nil
}

func (u UseCaseTransaction) hydrateCreditCard(transactionDTO dto.Transaction) *domain.CreditCard {
	creditCard := domain.NewCreditCard()
	creditCard.Name = transactionDTO.Name
	creditCard.Number = transactionDTO.Number
	creditCard.ExpirationMonth = transactionDTO.ExpirationMonth
	creditCard.ExpirationYear = transactionDTO.ExpirationYear
	creditCard.CVV = transactionDTO.CVV
	return creditCard
}

func (u UseCaseTransaction) newTransaction(transactionDTO dto.Transaction, creditCard domain.CreditCard) *domain.Transaction {
	transaction := domain.NewTransaction()
	transaction.CreditCardID = creditCard.ID
	transaction.Amount = transactionDTO.Amount
	transaction.Store = transactionDTO.Store
	transaction.Description = transactionDTO.Description
	transaction.CreatedAt = time.Now()
	return transaction
}
