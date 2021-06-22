package accountsvc

import (
	"github.com/authorizer-api/src/api/application/presentation/accountdto"
	"github.com/authorizer-api/src/api/domain/accountmd"
)

var (
	service Service
)

type Service interface {
	CreateAccount(request accountdto.AccountRequest) accountmd.Account
	GetAccounts() []accountmd.Account
	CleanAccounts()
}

func init() {
	service = newServiceImpl()
}

func Inject() Service {
	return service
}
