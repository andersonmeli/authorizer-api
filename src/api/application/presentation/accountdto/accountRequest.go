package accountdto

import (
	"github.com/authorizer-api/src/api/domain/accountmd"
)

type AccountRequest struct {
	ActiveCard 		bool    `json:"active-card" example:"true"`
	AvailableLimit  float64 `json:"available-limit" example:"100"`
}

func (request AccountRequest) ToModel() accountmd.Account {
	return accountmd.Account{
		ActiveCard: request.ActiveCard,
		AvailableLimit: request.AvailableLimit,
	}
}
