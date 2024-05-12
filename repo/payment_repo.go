package repo

import (
	"log"

	"github.com/gin-gonic/gin"
)

type PaymentRepo struct {
	db *DB
}

type PaymentRepository interface {
	WithdrawBalance(ctx *gin.Context, userAccount *Account, trans *Transaction) error
	SendBalance(ctx *gin.Context, userAccount, recieverAccount *Account, trans *Transaction) error
	GetTransactionsByAccountID(ctx *gin.Context, accountID string) ([]Transaction, error)
}

func NewPaymentRepo(db *DB) PaymentRepository {
	return &PaymentRepo{
		db: db,
	}
}

func (pr *PaymentRepo) WithdrawBalance(ctx *gin.Context, userAccount *Account, trans *Transaction) error {
	tx, _ := pr.db.Begin()
	defer tx.Close()

	if _, err := tx.Model(userAccount).WherePK().Update(); err != nil {
		tx.Rollback()
		log.Println("UpdateAccount error query " + err.Error())
		return err
	}

	if _, err := tx.Model(trans).Insert(); err != nil {
		tx.Rollback()
		log.Println("TransactionInsert error query " + err.Error())
		return err
	}

	if err := tx.Commit(); err != nil {
		log.Println("Commit error " + err.Error())
		return err
	}

	return nil
}

func (pr *PaymentRepo) SendBalance(ctx *gin.Context, userAccount, recieverAccount *Account, trans *Transaction) error {
	tx, _ := pr.db.Begin()
	defer tx.Close()

	if _, err := tx.Model(userAccount).WherePK().Update(); err != nil {
		tx.Rollback()
		log.Println("UpdateAccount error query " + err.Error())
		return err
	}

	if _, err := tx.Model(recieverAccount).WherePK().Update(); err != nil {
		tx.Rollback()
		log.Println("UpdateAccount reciever error query " + err.Error())
		return err
	}

	if _, err := tx.Model(trans).Insert(); err != nil {
		tx.Rollback()
		log.Println("TransactionInsert error query " + err.Error())
		return err
	}

	if err := tx.Commit(); err != nil {
		log.Println("Commit error " + err.Error())
		return err
	}

	return nil
}

func (pr *PaymentRepo) GetTransactionsByAccountID(ctx *gin.Context, accountID string) ([]Transaction, error) {
	var transactions []Transaction
	if err := pr.db.Model(&transactions).Where("account_id = ?", accountID).Select(); err != nil {
		log.Println("GetTransactionsByAccountID error query " + err.Error())
		return nil, err
	}

	return transactions, nil
}
