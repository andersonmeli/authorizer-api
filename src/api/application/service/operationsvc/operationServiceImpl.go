package operationsvc

import (
	"encoding/json"
	"fmt"
	"github.com/authorizer-api/src/api/application/presentation/accountdto"
	"github.com/authorizer-api/src/api/application/presentation/operationdto"
	"github.com/authorizer-api/src/api/application/presentation/transactiondto"
	"github.com/authorizer-api/src/api/application/service/accountsvc"
	"github.com/authorizer-api/src/api/application/service/transactionsvc"
	"log"
)

const (
	accountAlreadyInitialized = "account-already-initialized"
)

type serviceImpl struct {
	accountService     accountsvc.Service
	transactionService transactionsvc.Service
}

func (service serviceImpl) ProcessOperations(messages []string) {
	for _, message := range messages {
		var operations map[string]map[string]interface{}
		json.Unmarshal([]byte(message), &operations)

		for operationType := range operations {
			if operationType == "account" {
				createAccount(operations, service)
			} else if operationType == "transaction" {
				authorizationTransaction(operations, service)
			}
		}
	}
}

func authorizationTransaction(operations map[string]map[string]interface{}, service serviceImpl) {
	transactionRequest := transactiondto.TransactionRequest{
		Merchant: operations["transaction"]["merchant"].(string),
		Amount:   operations["transaction"]["amount"].(float64),
		Time:     operations["transaction"]["time"].(string),
	}

	transaction := service.transactionService.CreateTransaction(transactionRequest)
	accountResponse := service.transactionService.AuthorizationTransaction(transaction)
	formatOperationResponse(accountResponse)
}

func createAccount(operations map[string]map[string]interface{}, service serviceImpl) {
	var accountResponse accountdto.AccountResponse
	accountRequest := accountdto.AccountRequest{
		ActiveCard:     operations["account"]["active-card"].(bool),
		AvailableLimit: operations["account"]["available-limit"].(float64),
	}

	account := service.accountService.CreateAccount(accountRequest)
	if len(service.accountService.GetAccounts()) > 1 {
		accountResponse = accountdto.NewAccountResponse(service.accountService.GetAccounts()[0])
		accountResponse.Violations = append(accountResponse.Violations, accountAlreadyInitialized)
	} else {
		accountResponse = accountdto.NewAccountResponse(account)
	}
	formatOperationResponse(accountResponse)
}

func formatOperationResponse(accountResponse accountdto.AccountResponse) {
	buf, err := json.Marshal(operationdto.NewOperationResponse(accountResponse))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", buf)
}

func newServiceImpl(accountService accountsvc.Service, transactionService transactionsvc.Service) serviceImpl {
	return serviceImpl{
		accountService:            accountService,
		transactionService:        transactionService,
	}
}