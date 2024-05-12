package response

import "github.com/nahdukesaba/be-assignment/repo"

type TransactionsResponse struct {
	TransactionID int     `json:"transaction_id"`
	AccountID     string  `json:"account_id"`
	Amount        float64 `json:"amount"`
	Type          string  `json:"type"`
	Status        string  `json:"status"`
	RecieverID    string  `json:"reciever_id,omitempty"`
}

func NewTransactionsResponse(data []repo.Transaction) []TransactionsResponse {
	res := []TransactionsResponse{}
	for _, trans := range data {
		res = append(res, TransactionsResponse{
			TransactionID: trans.TransactionID,
			AccountID:     trans.AccountID,
			Amount:        trans.Amount,
			Type:          trans.Type,
			Status:        trans.Status,
			RecieverID:    trans.RecieverID,
		})
	}

	return res
}

type SendWithdrawResponse struct {
	AccountID string  `json:"account_id"`
	Balance   float64 `json:"balance"`
	Limit     float64 `json:"limit_credit,omitempty"`
	Status    string  `json:"status_transaction"`
}

func NewSendWithdrawResponse(acc *repo.Account, status string) *SendWithdrawResponse {
	res := new(SendWithdrawResponse)
	if acc != nil {
		res.AccountID = acc.AccountID
		res.Balance = acc.Balance
		res.Limit = acc.Limit
		res.Status = status
	}
	return res
}
