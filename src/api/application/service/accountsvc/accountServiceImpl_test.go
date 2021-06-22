package accountsvc

import (
	"github.com/authorizer-api/src/api/application/presentation/accountdto"
	"github.com/authorizer-api/src/api/domain/accountmd"
	"testing"
)

const (
	errorTest = "Valor esperado %v, Valor encontrado %v."
)

func TestCreateAccount(t *testing.T) {
	service := newServiceImpl()
	accountRequest := accountdto.AccountRequest{
		ActiveCard:     true,
		AvailableLimit: 1000,
	}
	account := accountmd.Account{
		ActiveCard:     true,
		AvailableLimit: 1000,
	}

	accountResult := service.CreateAccount(accountRequest)
	if account.ActiveCard != accountResult.ActiveCard {
		t.Errorf(errorTest, account.ActiveCard, accountResult.ActiveCard)
	}

	if account.AvailableLimit != accountResult.AvailableLimit {
		t.Errorf(errorTest, account.AvailableLimit, accountResult.AvailableLimit)
	}
}
