package accountsvc

import (
	"github.com/mercadolibre/authorizer-api/src/api/application/presentation/accountdto"
	"github.com/mercadolibre/authorizer-api/src/api/domain/accountmd"
)

var (
	service Service
)

type Service interface {
	CreateAccount(accountRequest accountdto.AccountRequest) accountmd.Account
}

func init() {
	service = newServiceImpl()
}

func Inject() Service {
	return service
}
