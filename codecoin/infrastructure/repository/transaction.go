package repository

import (
	"database/sql"
	"errors"

	"github.com/filipe-freitas/codecoin/domain"
)

type TransactionRepositoryDB struct {
	DB *sql.DB
}

func NewTransactionRepositoryDB(DB *sql.DB) *TransactionRepositoryDB {
	return &TransactionRepositoryDB{DB: DB}
}

func (t *TransactionRepositoryDB) GetCreditCard(creditCard domain.CreditCard) (domain.CreditCard, error) {
	var c domain.CreditCard
	stmt, err := t.DB.Prepare(`
		SELECT ID,
		      BALANCE,
					BALANCE_LIMIT
		 FROM CREDIT_CARDS
		WHERE NUMBER = $1`)
	if err != nil {
		return c, err
	}
	if err = stmt.QueryRow(creditCard.Number).Scan(&c.ID, &c.Balance, &c.Limit); err != nil {
		return c, errors.New("Credit card does not exist")
	}
	return c, nil
}

func (t *TransactionRepositoryDB) SaveTransaction(transaction domain.Transaction, creditCard domain.CreditCard) error {
	stmt, err := t.DB.Prepare(`
	INSERT
	  INTO TRANSACTIONS (ID,
											 CREDIT_CARD_ID,
											 AMOUNT,
											 STATUS,
											 DESCRIPTION,
											 STORE,
											 CREATED_AT)
		           VALUES ($1,
								       $2,
											 $3,
											 $4,
											 $5,
											 $6,
											 $7)`)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(
		transaction.ID,
		transaction.CreditCardID,
		transaction.Amount,
		transaction.Status,
		transaction.Description,
		transaction.Store,
		transaction.CreatedAt)
	if err != nil {
		return err
	}

	if transaction.Status == "approved" {
		err = t.updateBalance(creditCard)
		if err != nil {
			return err
		}
	}

	err = stmt.Close()
	if err != nil {
		return err
	}

	return nil
}

func (t *TransactionRepositoryDB) CreateCreditCard(creditCard domain.CreditCard) error {
	stmt, err := t.DB.Prepare(`
	  INSERT
		  INTO CREDIT_CARDS (ID,
				                 NAME,
												 NUMBER,
												 EXPIRATION_MONTH,
												 EXPIRATION_YEAR,
												 CVV,
												 BALANCE,
												 BALANCE_LIMIT)
								 VALUES ($1,
									       $2,
												 $3,
												 $4,
												 $5,
												 $6,
												 $7,
												 $8)`)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(
		creditCard.ID,
		creditCard.Name,
		creditCard.Number,
		creditCard.ExpirationMonth,
		creditCard.ExpirationYear,
		creditCard.CVV,
		creditCard.Balance,
		creditCard.Limit)
	if err != nil {
		return err
	}

	err = stmt.Close()
	if err != nil {
		return err
	}

	return nil
}

func (t *TransactionRepositoryDB) updateBalance(creditCard domain.CreditCard) error {
	_, err := t.DB.Exec(`
		UPDATE CREDIT_CARDS
		   SET BALANCE = $1
	   WHERE ID      = $2`,
		creditCard.Balance,
		creditCard.ID)
	if err != nil {
		return err
	}
	return nil
}
