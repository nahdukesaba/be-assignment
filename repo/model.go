package repo

import "time"

var (
	AccountTypeDebit  = "DEBIT"
	AccountTypeCredit = "CREDIT"

	TransactionTypeWithdrawal = "WITHDRAWAL"
	TransactionTypeTransfer   = "TRANSFER"
	TransactionTypeDeposit    = "DEPOSIT"

	StatusSuccess = "SUCCESS"
	StatusFail    = "FAIL"
)

// User represents a user entity
type User struct {
	UserID   int       `json:"user_id" pg:",pk,notnull"`
	Username string    `json:"username" pg:",unique,notnull"`
	Password string    `json:"password" pg:",notnull"`
	Accounts []Account `json:"accounts" pg:"rel:has-many"`
}

// PaymentAccount represents a user's payment account
type Account struct {
	AccountID    string        `json:"account_id" pg:",pk,notnull"`
	UserID       int           `json:"user_id" pg:",notnull"`
	AccountType  string        `json:"account_type" pg:",notnull"`
	Balance      float64       `json:"balance" pg:",notnull,use_zero"`
	Limit        float64       `json:"limit"`
	UpdatedAt    time.Time     `json:"updated_at" pg:"default:now()"`
	User         *User         `json:"user" pg:"rel:has-one"`
	Transactions []Transaction `json:"transactions" pg:"rel:has-many"`
}

// Transaction represents a payment transaction
type Transaction struct {
	TransactionID int       `json:"transaction_id" pg:",pk,notnull"`
	AccountID     string    `json:"account_id" pg:",notnull"`
	Amount        float64   `json:"amount" pg:",notnull"`
	Type          string    `json:"type" pg:",notnull"`
	RecieverID    string    `json:"reciever_id"`
	Status        string    `json:"status" pg:",notnull"`
	CreatedAt     time.Time `json:"created_at" pg:"default:now()"`
	Account       *Account  `json:"account" pg:"rel:has-one"`
	Reciever      *Account  `json:"reciever" pg:"rel:has-one"`
}
