package repo

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"github.com/nahdukesaba/be-assignment/helpers"
)

type UserRepo struct {
	db *DB
}

type UserRepository interface {
	CreateUser(ctx *gin.Context, user *User) error
	GetUserByUsername(ctx *gin.Context, username string) (*User, error)
	GetUserAccountByAccountNumber(ctx *gin.Context, accNumber string) (*Account, error)
	GetAccountsByUserID(ctx *gin.Context, userID int) ([]Account, error)
}

func NewUserRepo(db *DB) UserRepository {
	return &UserRepo{
		db: db,
	}
}

func (ur *UserRepo) CreateUser(ctx *gin.Context, user *User) error {
	tx, _ := ur.db.Begin()
	defer tx.Close()

	_, err := tx.Model(user).Insert()
	if err != nil {
		log.Println("CreateUser error query " + err.Error())
		_ = tx.Rollback()
		return err
	}

	if err := ur.GenerateBonusNewUser(ctx, tx, user); err != nil {
		log.Println("GenerateUserAccount error " + err.Error())
		_ = tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		log.Println("Commit error " + err.Error())
		return err
	}

	return nil
}

func (ur *UserRepo) GetUserByUsername(ctx *gin.Context, username string) (*User, error) {
	user := new(User)
	if err := ur.db.Model(user).Where("username = ?", username).Select(); err != nil {
		log.Println("GetUserByUsername error query " + err.Error())
		return nil, err
	}

	return user, nil
}

func (ur *UserRepo) GetUserAccountByAccountNumber(ctx *gin.Context, accNumber string) (*Account, error) {
	account := new(Account)
	if err := ur.db.Model(account).Relation("User").Where("account_id = ?", accNumber).Select(); err != nil {
		log.Println("GetUserAccountByAccountNumber error query " + err.Error())
		return nil, err
	}

	return account, nil
}

func (ur *UserRepo) GenerateBonusNewUser(ctx *gin.Context, tx *pg.Tx, user *User) error {
	accountDebit := &Account{
		AccountID:   helpers.GenerateAccountNumber(AccountTypeDebit),
		UserID:      user.UserID,
		AccountType: AccountTypeDebit,
		Balance:     1000000,
		User:        user,
	}
	accountCredit := &Account{
		AccountID:   helpers.GenerateAccountNumber(AccountTypeCredit),
		UserID:      user.UserID,
		AccountType: AccountTypeCredit,
		Balance:     0,
		Limit:       10000000,
		User:        user,
	}

	trans := &Transaction{
		Amount:    1000000,
		Type:      TransactionTypeDeposit,
		Status:    StatusSuccess,
		CreatedAt: time.Now(),
	}

	if err := ur.db.RunInTransaction(ctx, func(tx *pg.Tx) error {
		if _, err := tx.Model(accountCredit).Insert(); err != nil {
			log.Println("Insert account error query " + err.Error())
			return err
		}

		if _, err := tx.Model(accountDebit).Insert(); err != nil {
			log.Println("Insert account error query " + err.Error())
			return err
		}

		trans.AccountID = accountDebit.AccountID
		if _, err := tx.Model(trans).Insert(); err != nil {
			log.Println("Insert transaction error query " + err.Error())
			return err
		}
		return nil
	}); err != nil {
		return err
	}

	return nil
}

func (ur *UserRepo) GetAccountsByUserID(ctx *gin.Context, userID int) ([]Account, error) {
	var accounts []Account
	if err := ur.db.Model(&accounts).Where("user_id = ?", userID).Select(); err != nil {
		log.Println("GetUserAccountByAccountNumber error query " + err.Error())
		return nil, err
	}

	return accounts, nil
}
