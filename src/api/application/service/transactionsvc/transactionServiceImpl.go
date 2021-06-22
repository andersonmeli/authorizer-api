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
	highFrequencySmallInterval = "high-frequency-small-interval"
	insufficientLimit = "insufficient-limit"
	cardNotActive = "card-not-active"
	doubleTransaction = "double-transaction"
)

var(
	transactions []transactionmd.Transaction
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

	//Nenhuma transação deve ser aceita sem que a conta tenha sido inicializada
	if len(service.accountService.GetAccounts()) == 0 {
		accountResponse = accountdto.AccountResponse{
			Violations:     violations,
		}
	}else{
		accountResponse = service.validateTransaction(transaction, &service.accountService.GetAccounts()[0])
	}

	return accountResponse
}

func (service serviceImpl) CreateTransaction(request transactiondto.TransactionRequest) transactionmd.Transaction {
	return request.ToModel()
}

func (service serviceImpl) validateTransaction(transaction transactionmd.Transaction, account *accountmd.Account) accountdto.AccountResponse {
	var violations []string
	accountResponse := accountdto.AccountResponse{
		ActiveCard:     account.ActiveCard,
		AvailableLimit: account.AvailableLimit,
		Violations:     []string{},
	}

	//Nenhuma transação deve ser aceita quando o cartão não estiver ativo
	if account.ActiveCard == false {
		violations = append(violations, cardNotActive)
	}

	//O valor da transação não deve exceder o limite disponível
	if transaction.Amount > account.AvailableLimit {
		violations = append(violations, insufficientLimit)
	}

	if len(violations) == 0 {
		newAmount := account.AvailableLimit - transaction.Amount
		accountResponse.AvailableLimit = newAmount
		account.AvailableLimit = newAmount
		transactions = append(transactions, transaction)
	}else{
		accountResponse.Violations = violations
	}

	return accountResponse
}