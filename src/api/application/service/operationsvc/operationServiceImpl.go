package operationsvc

import (
	"encoding/json"
	"fmt"
	"github.com/authorizer-api/src/api/application/presentation/accountdto"
	"github.com/authorizer-api/src/api/application/service/accountsvc"
	"github.com/authorizer-api/src/api/application/service/transactionsvc"
	"log"
)

const (
	accountAlreadyInitialized = "account-already-initialized"
	highFrequencySmallInterval = "high-frequency-small-interval"
	insufficientLimit = "insufficient-limit"
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
				accountRequest := accountdto.AccountRequest{
					ActiveCard:     operations["account"]["active-card"].(bool),
					AvailableLimit: operations["account"]["available-limit"].(float64),
				}

				account := service.accountService.CreateAccount(accountRequest)

				var response accountdto.AccountResponse
				if len(service.accountService.GetAccounts()) > 1 {
					response = accountdto.NewAccountResponse(service.accountService.GetAccounts()[0])
					response.Violations = append(response.Violations, accountAlreadyInitialized)
				}else{
					response = accountdto.NewAccountResponse(account)
				}

				buf, err := json.Marshal(response)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Printf("%s\n", buf)
			}else if operationType == "transaction" {

			}
		}
	}
}

func newServiceImpl(accountService accountsvc.Service, transactionService transactionsvc.Service) serviceImpl {
	return serviceImpl{
		accountService:            accountService,
		transactionService:        transactionService,
	}
}