package transactiondto

import (
	"github.com/authorizer-api/src/api/domain/transactionmd"
	"time"
)

type TransactionRequest struct {
	Merchant string
	Amount float64
	Time string
}

func (request TransactionRequest) ToModel() transactionmd.Transaction {
	time, _ := time.Parse(time.RFC3339, request.Time)
	return transactionmd.Transaction{
		Merchant: request.Merchant,
		Amount: request.Amount,
		Time: time,
	}
}
