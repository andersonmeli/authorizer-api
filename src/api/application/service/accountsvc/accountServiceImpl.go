package accountsvc

import (
	"github.com/authorizer-api/src/api/application/presentation/accountdto"
	"github.com/authorizer-api/src/api/domain/accountmd"
)

var(
	account accountmd.Account
)
type serviceImpl struct {
}

func newServiceImpl() serviceImpl {
	return serviceImpl{}
}

func (service serviceImpl) CreateAccount(accountRequest accountdto.AccountRequest) accountmd.Account {
	return accountRequest.ToModel()
}