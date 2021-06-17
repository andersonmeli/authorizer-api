package operationdto

import (
	"github.com/mercadolibre/authorizer-api/src/api/domain/operationmd"
)

type OperationRequest struct {
	Type string
	Msg  map[string]interface{}
}

func (request OperationRequest) ToModel() operationmd.Operation {
	return operationmd.Operation{
		Type: request.Type,
		Msg: request.Msg,
	}
}