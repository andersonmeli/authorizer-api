package controller

import (
	"github.com/mercadolibre/authorizer-api/src/api/application/service/operationsvc"
)

var (
	operationService = operationsvc.Inject()
)

func ProcessOperations(messages []string) {
	operationService.ProcessOperations(messages)
}