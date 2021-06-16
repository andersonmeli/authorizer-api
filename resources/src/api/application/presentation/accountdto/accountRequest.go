package accountdto

import (
	"github.com/mercadolibre/authorizer-api/resources/src/api/domain/accountmd"
)

type AccountRequest struct {
	ActiveCard 		bool
	AvailableLimit  float64
}

func (request AccountRequest) ToModel() (accountmd.Account) {
	return accountmd.Account{
		ActiveCard: request.ActiveCard,
		AvailableLimit: request.AvailableLimit,
	}
}
