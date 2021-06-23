package transactionsvc

import (
	"github.com/authorizer-api/src/api/application/presentation/accountdto"
	"github.com/authorizer-api/src/api/application/presentation/transactiondto"
	"github.com/authorizer-api/src/api/application/service/accountsvc"
	"github.com/authorizer-api/src/api/domain/accountmd"
	"github.com/authorizer-api/src/api/domain/transactionmd"
	"time"
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

func (service serviceImpl) AuthorizationTransaction(transaction transactionmd.Transaction) (accountdto.AccountResponse, []string) {
	var accountResponse accountdto.AccountResponse
	violations := make([]string, 0)
	//Nenhuma transação deve ser aceita sem que a conta tenha sido inicializada
	if len(service.accountService.GetAccounts()) == 0 {
		violations = append(violations, accountNotInitialized)
	}else{
		accountResponse, violations = service.validateTransaction(transaction, &service.accountService.GetAccounts()[0])
	}

	return accountResponse, violations
}

func (service serviceImpl) CreateTransaction(request transactiondto.TransactionRequest) transactionmd.Transaction {
	return request.ToModel()
}

func (service serviceImpl) validateTransaction(transaction transactionmd.Transaction, account *accountmd.Account) (accountdto.AccountResponse, []string) {
	violations := make([]string, 0)
	accountResponse := accountdto.AccountResponse{
		ActiveCard:     account.ActiveCard,
		AvailableLimit: account.AvailableLimit,
	}

	//Nenhuma transação deve ser aceita quando o cartão não estiver ativo
	if account.ActiveCard == false {
		violations = append(violations, cardNotActive)
	}

	//O valor da transação não deve exceder o limite disponível
	if transaction.Amount > account.AvailableLimit {
		violations = append(violations, insufficientLimit)
	}

	//Não deve haver mais que 3 transações de qualquer comerciante em um intervalo de 2 minutos
	if len(transactions) >= 3 {
		firstTransaction := transactions[0]
		firstTransactionTreeMinutes :=  firstTransaction.Time.Add(3 * time.Minute)
		if firstTransactionTreeMinutes.After(transaction.Time) {
			violations = append(violations, highFrequencySmallInterval)
		}
	}

	//Não deve haver mais que 1 transação similar (mesmo valor e comerciante) no intervalo de 2 minutos
	if len(transactions) > 0 {
		for _, transactionFinalized := range transactions {
			transactionTwoMinutes :=  transactionFinalized.Time.Add(2 * time.Minute)
			if transactionFinalized.Merchant == transaction.Merchant && transactionTwoMinutes.After(transaction.Time) {
				violations = append(violations, doubleTransaction)
			}
		}
	}

	if len(violations) == 0 {
		newAmount := account.AvailableLimit - transaction.Amount
		accountResponse.AvailableLimit = newAmount
		account.AvailableLimit = newAmount
		//Guardando apenas transações que não tiveram violações
		transactions = append(transactions, transaction)
	}

	return accountResponse, violations
}

func (service serviceImpl) CleanTransactions() {
	transactions = make([]transactionmd.Transaction, 0)
}