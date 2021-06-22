package accountdto

import "github.com/authorizer-api/src/api/domain/accountmd"

type AccountResponse struct {
	ActiveCard 		bool     `json:"active-card" example:"true"`
	AvailableLimit  float64  `json:"available-limit" example:"100"`
}

func NewAccountResponse(account accountmd.Account) AccountResponse {
	return AccountResponse{
		ActiveCard:     account.ActiveCard,
		AvailableLimit: account.AvailableLimit,
	}
}