package operationsvc

import (
	"github.com/mercadolibre/authorizer-api/src/api/application/service/accountsvc"
	"github.com/mercadolibre/authorizer-api/src/api/application/service/transactionsvc"
)

var (
	service Service
)

type Service interface {
	ProcessOperations(messages []string)
}

func init() {
	service = newServiceImpl(accountsvc.Inject(), transactionsvc.Inject())
}

func Inject() Service {
	return service
}