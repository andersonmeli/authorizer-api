package accountsvc

import (
	"github.com/authorizer-api/src/api/application/presentation/accountdto"
	"github.com/authorizer-api/src/api/domain/accountmd"
)

var(
	accounts []accountmd.Account
)

type serviceImpl struct {
}

func newServiceImpl() serviceImpl {
	return serviceImpl{}
}

func (service serviceImpl) CreateAccount(request accountdto.AccountRequest) accountmd.Account {
	accountInitialize := request.ToModel()
	accounts = append(accounts, accountInitialize)
	return accountInitialize
}

func (service serviceImpl) GetAccounts() []accountmd.Account {
	return accounts
}

func (service serviceImpl) CleanAccounts() {
	accounts = make([]accountmd.Account, 0)
}

