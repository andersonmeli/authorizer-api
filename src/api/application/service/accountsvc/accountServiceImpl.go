package accountsvc

import (
	"github.com/authorizer-api/src/api/application/presentation/accountdto"
	"github.com/authorizer-api/src/api/domain/accountmd"
)

var(
	account accountmd.Account
	accounts []accountmd.Account
)
type serviceImpl struct {
}

func newServiceImpl() serviceImpl {
	return serviceImpl{}
}

func (service serviceImpl) CreateAccount(accountRequest accountdto.AccountRequest) accountmd.Account {
	accountInitialize := accountRequest.ToModel()
	accounts = append(accounts, accountInitialize)
	return accountInitialize
}

func (service serviceImpl) GetAccounts() []accountmd.Account {
	return accounts
}
