package operationsvc

import (
	"encoding/json"
	"fmt"
	"github.com/mercadolibre/authorizer-api/resources/src/api/application/presentation/accountdto"
	"github.com/mercadolibre/authorizer-api/resources/src/api/application/service/accountsvc"
	"github.com/mercadolibre/authorizer-api/resources/src/api/application/service/transactionsvc"
	"log"
)

type serviceImpl struct {
	accountService            accountsvc.Service
	transactionService        transactionsvc.Service
}

func (service serviceImpl) ProcessOperations(messages []string) {
	for _, message := range messages {
		fmt.Println(message)
		var result map[string]map[string]interface{}
		json.Unmarshal([]byte(message), &result)
		fmt.Println(result)

		for k := range result {
			if k == "account" {
				fmt.Println(k)
				account := accountdto.AccountRequest{
					ActiveCard:     result["account"]["active-card"].(bool),
					AvailableLimit: result["account"]["available-limit"].(float64),
				}
				fmt.Println(account)
				buf, err := json.Marshal(account)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Printf("%s\n", buf)
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