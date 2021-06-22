package transactionsvc

import (
	"github.com/authorizer-api/src/api/application/presentation/accountdto"
	"github.com/authorizer-api/src/api/application/presentation/transactiondto"
	"github.com/authorizer-api/src/api/application/service/accountsvc"
	"github.com/authorizer-api/src/api/domain/transactionmd"
)

var (
	service Service
)

type Service interface {
	CreateTransaction(request transactiondto.TransactionRequest) transactionmd.Transaction
	AuthorizationTransaction(transaction transactionmd.Transaction) (accountdto.AccountResponse, []string)
}

func init() {
	service = newServiceImpl(accountsvc.Inject())
}

func Inject() Service {
	return service
}
