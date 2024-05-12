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
