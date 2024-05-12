package request

import "errors"

type TransactionRequest struct {
	Sender   string  `json:"sender"`
	Amount   float64 `json:"amount"`
	Reciever string  `json:"reciever"`
	Username string  `json:"-"`
}

func (t *TransactionRequest) Validate(isSend bool) error {
	if t.Sender == "" {
		return errors.New("sender account number cannot be empty")
	}

	if t.Sender != "" && len(t.Sender) != 10 {
		return errors.New("invalid sender account number")
	}

	if t.Amount < 1 {
		return errors.New("amount cannot be empty or negative")
	}

	if isSend {
		if t.Reciever == "" {
			return errors.New("reciever account number cannot be empty")
		}

		if t.Reciever != "" && len(t.Reciever) != 10 {
			return errors.New("invalid reciever account number")
		}

		if t.Reciever == t.Sender {
			return errors.New("cannot send to the same account number")
		}
	}
	return nil
}

type TransactionsAccountRequest struct {
	Username  string `json:"-"`
	AccountID string `json:"-"`
}

func (t *TransactionsAccountRequest) Validate() error {
	if t.Username == "" {
		return errors.New("username cannot be empty")
	}

	if t.AccountID == "" {
		return errors.New("account number cannot be empty")
	}

	return nil
}
