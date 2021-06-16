package transactiondto

import (
	"github.com/mercadolibre/authorizer-api/resources/src/api/domain/transactionmd"
	"github.com/mercadolibre/authorizer-api/resources/src/api/infrastructure/excp"
)

type TransactionRequest struct {
	Merchant string
	Amount int
	Time string
}

func (request TransactionRequest) ToModel() (transactionmd.Transaction, excp.Exception) {
	return transactionmd.Transaction{
		Merchant: request.Merchant,
		Amount: request.Amount,
		Time: request.Time,
	}, nil
}
