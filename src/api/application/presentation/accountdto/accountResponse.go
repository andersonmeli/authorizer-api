package accountdto

import "github.com/authorizer-api/src/api/domain/accountmd"

type AccountResponse struct {
	ActiveCard 		bool     `json:"active-card,omitempty" example:"true"`
	AvailableLimit  float64  `json:"available-limit,omitempty" example:"100"`
}

func NewAccountResponse(account accountmd.Account) AccountResponse {
	return AccountResponse{
		ActiveCard:     account.ActiveCard,
		AvailableLimit: account.AvailableLimit,
	}
}