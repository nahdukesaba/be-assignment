package services

import (
	"errors"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nahdukesaba/be-assignment/repo"
	"github.com/nahdukesaba/be-assignment/services/request"
	"github.com/nahdukesaba/be-assignment/services/response"
)

type PaymentService struct {
	paymentRepo repo.PaymentRepository
	userRepo    repo.UserRepository
}

type Payments interface {
	WithdrawBalance(ctx *gin.Context, form *request.TransactionRequest) (*response.SendWithdrawResponse, error)
	SendBalance(ctx *gin.Context, form *request.TransactionRequest) (*response.SendWithdrawResponse, error)
	GetAllAccountsUser(ctx *gin.Context, username string) ([]response.AccountsResponse, error)
	GetAllTransactionsAccount(ctx *gin.Context, form *request.TransactionsAccountRequest) ([]response.TransactionsResponse, error)
}

func NewPaymentService(db *repo.DB) Payments {
	return &PaymentService{
		paymentRepo: repo.NewPaymentRepo(db),
		userRepo:    repo.NewUserRepo(db),
	}
}

func (ps *PaymentService) WithdrawBalance(ctx *gin.Context, form *request.TransactionRequest) (*response.SendWithdrawResponse, error) {
	if err := form.Validate(false); err != nil {
		return nil, err
	}

	userAccount, err := ps.userRepo.GetUserAccountByAccountNumber(ctx, form.Sender)
	if err != nil {
		log.Printf("Error GetUserAccountByAccountNumber. err: %s\n", err)
		return nil, err
	}

	if userAccount.User.Username != form.Username {
		log.Println("Error wrong account")
		return nil, errors.New("error wrong account")
	}

	if userAccount.AccountType == repo.AccountTypeDebit {
		remaining := userAccount.Balance - form.Amount
		if remaining < 0 {
			log.Println("insufficient balance")
			return nil, errors.New("your balance is insufficient")
		}
		userAccount.Balance = remaining
	} else if userAccount.AccountType == repo.AccountTypeCredit {
		total := userAccount.Balance + form.Amount
		if total > userAccount.Limit {
			log.Println("credit limit exceeded")
			return nil, errors.New("your credit limit exceeded")
		}
		userAccount.Balance = total
	}

	userAccount.UpdatedAt = time.Now()
	trans := &repo.Transaction{
		AccountID: userAccount.AccountID,
		Amount:    form.Amount,
		Status:    repo.StatusSuccess,
		CreatedAt: time.Now(),
		Type:      repo.TransactionTypeWithdrawal,
	}

	if err := ps.paymentRepo.WithdrawBalance(ctx, userAccount, trans); err != nil {
		log.Printf("Error WithdrawBalance. err: %s\n", err)
		return nil, err
	}

	return response.NewSendWithdrawResponse(userAccount, repo.StatusSuccess), nil
}

func (ps *PaymentService) SendBalance(ctx *gin.Context, form *request.TransactionRequest) (*response.SendWithdrawResponse, error) {
	if err := form.Validate(true); err != nil {
		return nil, err
	}

	userAccount, err := ps.userRepo.GetUserAccountByAccountNumber(ctx, form.Sender)
	if err != nil {
		log.Printf("Error GetUserAccountByAccountNumber. err: %s\n", err)
		return nil, err
	}

	recieverAccount, err := ps.userRepo.GetUserAccountByAccountNumber(ctx, form.Reciever)
	if err != nil {
		log.Printf("Error reciever GetUserAccountByAccountNumber. err: %s\n", err)
		return nil, err
	}

	if userAccount.User.Username != form.Username {
		log.Println("Error wrong account")
		return nil, errors.New("error wrong account")
	}

	if userAccount.AccountType == repo.AccountTypeDebit {
		remaining := userAccount.Balance - form.Amount
		if remaining < 0 {
			log.Println("insufficient balance")
			return nil, errors.New("your balance is insufficient")
		}
		userAccount.Balance = remaining
	} else if userAccount.AccountType == repo.AccountTypeCredit {
		total := userAccount.Balance + form.Amount
		if total > userAccount.Limit {
			log.Println("credit limit exceeded")
			return nil, errors.New("your credit limit exceeded")
		}
		userAccount.Balance = total
	}

	if recieverAccount.AccountType == repo.AccountTypeDebit {
		recieverAccount.Balance += form.Amount
	} else if recieverAccount.AccountType == repo.AccountTypeCredit {
		recieverAccount.Balance -= form.Amount
	}

	userAccount.UpdatedAt = time.Now()
	recieverAccount.UpdatedAt = time.Now()
	trans := &repo.Transaction{
		AccountID:  userAccount.AccountID,
		Amount:     form.Amount,
		Status:     repo.StatusSuccess,
		CreatedAt:  time.Now(),
		Type:       repo.TransactionTypeTransfer,
		RecieverID: form.Reciever,
	}

	if err := ps.paymentRepo.SendBalance(ctx, userAccount, recieverAccount, trans); err != nil {
		log.Printf("Error SendBalance. err: %s\n", err)
		return nil, err
	}

	return response.NewSendWithdrawResponse(userAccount, repo.StatusSuccess), nil
}

func (ps *PaymentService) GetAllAccountsUser(ctx *gin.Context, username string) ([]response.AccountsResponse, error) {
	user, err := ps.userRepo.GetUserByUsername(ctx, username)
	if err != nil {
		log.Printf("Error GetUserAccountByAccountNumber. err: %s\n", err)
		return nil, err
	}

	res, err := ps.userRepo.GetAccountsByUserID(ctx, user.UserID)
	if err != nil {
		log.Printf("Error GetAccountsByUserID. err: %s\n", err)
		return nil, err
	}

	return response.NewAccountsResponse(res), nil
}

func (ps *PaymentService) GetAllTransactionsAccount(ctx *gin.Context, form *request.TransactionsAccountRequest) ([]response.TransactionsResponse, error) {
	if err := form.Validate(); err != nil {
		return nil, err
	}

	userAccount, err := ps.userRepo.GetUserAccountByAccountNumber(ctx, form.AccountID)
	if err != nil {
		log.Printf("Error GetUserAccountByAccountNumber. err: %s\n", err)
		return nil, err
	}

	if userAccount.User.Username != form.Username {
		log.Println("Error wrong account")
		return nil, errors.New("error wrong account")
	}

	res, err := ps.paymentRepo.GetTransactionsByAccountID(ctx, form.AccountID)
	if err != nil {
		log.Printf("Error GetTransactionsByAccountID. err: %s\n", err)
		return nil, err
	}

	return response.NewTransactionsResponse(res), nil
}
