package transactiondto

import (
	"github.com/mercadolibre/authorizer-api/src/api/domain/transactionmd"
)

type TransactionRequest struct {
	Merchant string
	Amount int
	Time string
}

func (request TransactionRequest) ToModel() transactionmd.Transaction {
	return transactionmd.Transaction{
		Merchant: request.Merchant,
		Amount: request.Amount,
		Time: request.Time,
	}
}
