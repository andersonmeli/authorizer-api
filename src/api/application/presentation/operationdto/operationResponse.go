package operationdto

import (
	"github.com/authorizer-api/src/api/application/presentation/accountdto"
)

type OperationResponse struct {
	Account accountdto.AccountResponse `json:"account,omitempty"`
	Violations []string 			   `json:"violations"`
}

func NewOperationResponse(accountResponse accountdto.AccountResponse, violations []string) OperationResponse {
	return OperationResponse{
		Account: accountResponse,
		Violations: violations,
	}
}