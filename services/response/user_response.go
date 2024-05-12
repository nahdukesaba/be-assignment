package response

import "github.com/nahdukesaba/be-assignment/repo"

type UserLoginUser struct {
	Token    string `json:"token"`
	Username string `json:"username"`
}

type AccountsResponse struct {
	AccountID   string  `json:"account_id"`
	UserID      int     `json:"user_id"`
	AccountType string  `json:"account_type"`
	Balance     float64 `json:"balance"`
	Limit       float64 `json:"limit,omitempty"`
}

func NewAccountsResponse(data []repo.Account) []AccountsResponse {
	res := []AccountsResponse{}
	for _, acc := range data {
		res = append(res, AccountsResponse{
			AccountID:   acc.AccountID,
			UserID:      acc.UserID,
			AccountType: acc.AccountType,
			Balance:     acc.Balance,
			Limit:       acc.Limit,
		})
	}
	return res
}
