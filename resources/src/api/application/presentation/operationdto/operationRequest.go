package operationdto

import (
	"github.com/mercadolibre/authorizer-api/resources/src/api/domain/operationmd"
	"github.com/mercadolibre/authorizer-api/resources/src/api/infrastructure/excp"
)

type OperationRequest struct {
	Type string
	Msg  map[string]interface{}
}

func (request OperationRequest) ToModel() (operationmd.Operation, excp.Exception) {
	return operationmd.Operation{
		Type: request.Type,
		Msg: request.Msg,
	}, nil
}