package transactiondto

import (
	"github.com/authorizer-api/src/api/domain/transactionmd"
)

type TransactionRequest struct {
	Merchant string
	Amount float64
	Time string
}

func (request TransactionRequest) ToModel() transactionmd.Transaction {
	return transactionmd.Transaction{
		Merchant: request.Merchant,
		Amount: request.Amount,
		Time: request.Time,
	}
}
