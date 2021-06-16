package controller

import (
	"github.com/mercadolibre/authorizer-api/resources/src/api/application/service/operationsvc"
)

var (
	operationService = operationsvc.Inject()
)

func ProcessOperations(messages []string) {
	operationService.ProcessOperations(messages)
}