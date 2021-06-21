package transactionsvc

import (
	"github.com/authorizer-api/src/api/application/presentation/accountdto"
	"github.com/authorizer-api/src/api/application/presentation/transactiondto"
	"github.com/authorizer-api/src/api/application/service/accountsvc"
	"github.com/authorizer-api/src/api/domain/accountmd"
	"github.com/authorizer-api/src/api/domain/transactionmd"
)

const (
	accountNotInitialized = "account-not-initialized"
)
type serviceImpl struct {
	accountService     accountsvc.Service
}

func newServiceImpl(accountService accountsvc.Service) serviceImpl {
	return serviceImpl{
		accountService: accountService,
	}
}

func (service serviceImpl) AuthorizationTransaction(transaction transactionmd.Transaction) accountdto.AccountResponse {

	var accountResponse accountdto.AccountResponse
	var violations []string
	violations = append(violations, accountNotInitialized)

	if len(service.accountService.GetAccounts()) == 0 {
		accountResponse = accountdto.AccountResponse{
			Violations:     violations,
		}
	}else{
		account := service.accountService.GetAccounts()[0]
		service.validateTransaction(transaction, account)
	}

	return accountResponse
}

func (service serviceImpl) CreateTransaction(request transactiondto.TransactionRequest) transactionmd.Transaction {
	return request.ToModel()
}

func (service serviceImpl) validateTransaction(transaction transactionmd.Transaction, account accountmd.Account) string {

	return ""
}