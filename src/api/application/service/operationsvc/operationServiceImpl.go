package operationsvc

import (
	"encoding/json"
	"fmt"
	"github.com/mercadolibre/authorizer-api/src/api/application/presentation/accountdto"
	"github.com/mercadolibre/authorizer-api/src/api/application/service/accountsvc"
	"github.com/mercadolibre/authorizer-api/src/api/application/service/transactionsvc"
	"log"
)

type serviceImpl struct {
	accountService     accountsvc.Service
	transactionService transactionsvc.Service
}

func (service serviceImpl) ProcessOperations(messages []string) {
	for _, message := range messages {
		var result map[string]map[string]interface{}
		json.Unmarshal([]byte(message), &result)

		for k := range result {
			if k == "account" {
				accountRequest := accountdto.AccountRequest{
					ActiveCard:     result["account"]["active-card"].(bool),
					AvailableLimit: result["account"]["available-limit"].(float64),
				}

				account := service.accountService.CreateAccount(accountRequest)
				response := accountdto.NewAccountResponse(account)

				fmt.Println(response)
				buf, err := json.Marshal(response)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Printf("%s\n", buf)
			}else if k == "transaction" {

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