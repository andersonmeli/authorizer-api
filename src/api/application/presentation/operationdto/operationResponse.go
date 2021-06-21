package operationdto

import (
	"github.com/authorizer-api/src/api/application/presentation/accountdto"
)

type OperationResponse struct {
	Account accountdto.AccountResponse `json:"account"`
}

func NewOperationResponse(accountResponse accountdto.AccountResponse) OperationResponse {
	return OperationResponse{
		Account: accountResponse,
	}
}