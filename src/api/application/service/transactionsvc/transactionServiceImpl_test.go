package transactionsvc

import (
	"github.com/authorizer-api/src/api/application/presentation/accountdto"
	"github.com/authorizer-api/src/api/application/presentation/transactiondto"
	"github.com/authorizer-api/src/api/application/service/accountsvc"
	"github.com/authorizer-api/src/api/domain/transactionmd"
	"testing"
	"time"
)

const (
	errorTest = "Valor esperado %v, Valor encontrado %v."
)

func TestCreateTransaction(t *testing.T) {
	service := newServiceImpl(accountsvc.Inject())
	service.accountService.CleanAccounts()

	transactionRequest := transactiondto.TransactionRequest{
		Merchant: "Samsung",
		Amount:   8000,
		Time:     "2019-02-13T11:00:00.000Z",
	}

	time, _ := time.Parse(time.RFC3339, "2019-02-13T11:00:00.000Z")
	transaction := transactionmd.Transaction{
		Merchant: "Samsung",
		Amount:   8000,
		Time:     time,
	}

	transactionResult := service.CreateTransaction(transactionRequest)
	if transaction.Merchant != transactionResult.Merchant {
		t.Errorf(errorTest, transaction.Merchant, transactionResult.Merchant)
	}

	if transaction.Amount != transactionResult.Amount {
		t.Errorf(errorTest, transaction.Amount, transactionResult.Amount)
	}

	if transaction.Time != transactionResult.Time {
		t.Errorf(errorTest, transaction.Time, transactionResult.Time)
	}
}

func TestAuthorizationTransaction(t *testing.T){
	service := newServiceImpl(accountsvc.Inject())
	service.accountService.CleanAccounts()

	accountRequest := accountdto.AccountRequest{
		ActiveCard:     true,
		AvailableLimit: 10000,
	}

	accountResult := service.accountService.CreateAccount(accountRequest)
	accountRespose := accountdto.NewAccountResponse(accountResult)

	transactionRequest := transactiondto.TransactionRequest{
		Merchant: "Samsung",
		Amount:   8000,
		Time:     "2019-02-13T11:00:00.000Z",
	}

	transactionResult := service.CreateTransaction(transactionRequest)
	accountResponseResult, violations := service.AuthorizationTransaction(transactionResult)
	if accountRespose.ActiveCard != accountResponseResult.ActiveCard {
		t.Errorf(errorTest, accountRespose.ActiveCard, accountResponseResult.ActiveCard)
	}

	if accountResponseResult.AvailableLimit != accountRespose.AvailableLimit - transactionResult.Amount {
		t.Errorf(errorTest, accountRespose.AvailableLimit - transactionResult.Amount, accountResponseResult.AvailableLimit)
	}

	if len(violations) != 0 {
		t.Errorf(errorTest, 0, len(violations))
	}
}

func TestAuthorizationTransactionAccountNotInitialized(t *testing.T){
	service := newServiceImpl(accountsvc.Inject())
	service.accountService.CleanAccounts()

	transactionRequest := transactiondto.TransactionRequest{
		Merchant: "Samsung",
		Amount:   8000,
		Time:     "2019-02-13T11:00:00.000Z",
	}

	transactionResult := service.CreateTransaction(transactionRequest)
	_, violations := service.AuthorizationTransaction(transactionResult)

	if len(violations) == 0 {
		t.Errorf(errorTest, 1, len(violations))
	}

	if violations[0] != accountNotInitialized {
		t.Errorf(errorTest, accountNotInitialized, violations[0])
	}
}

func TestAuthorizationTransactionInsufficientLimit(t *testing.T){
	service := newServiceImpl(accountsvc.Inject())
	service.accountService.CleanAccounts()

	accountRequest := accountdto.AccountRequest{
		ActiveCard:     true,
		AvailableLimit: 10000,
	}

	accountResult := service.accountService.CreateAccount(accountRequest)
	accountRespose := accountdto.NewAccountResponse(accountResult)

	transactionRequest := transactiondto.TransactionRequest{
		Merchant: "Samsung",
		Amount:   18000,
		Time:     "2019-02-13T11:00:00.000Z",
	}

	transactionResult := service.CreateTransaction(transactionRequest)
	accountResponseResult, violations := service.AuthorizationTransaction(transactionResult)
	if accountRespose.ActiveCard != accountResponseResult.ActiveCard {
		t.Errorf(errorTest, accountRespose.ActiveCard, accountResponseResult.ActiveCard)
	}

	if accountResponseResult.AvailableLimit != accountRespose.AvailableLimit {
		t.Errorf(errorTest, accountRespose.AvailableLimit, accountResponseResult.AvailableLimit)
	}

	if len(violations) == 0 {
		t.Errorf(errorTest, 0, len(violations))
	}

	if violations[0] != insufficientLimit {
		t.Errorf(errorTest, insufficientLimit, violations[0])
	}
}

func TestAuthorizationTransactionCardNotActive(t *testing.T){
	service := newServiceImpl(accountsvc.Inject())
	service.accountService.CleanAccounts()

	account := service.accountService.CreateAccount(
		accountdto.AccountRequest{
			ActiveCard:     false,
			AvailableLimit: 10000,
		})
	accountRespose := accountdto.NewAccountResponse(account)

	transaction := service.CreateTransaction(
		transactiondto.TransactionRequest{
			Merchant: "Samsung",
			Amount:   8000,
			Time:     "2019-02-13T11:00:00.000Z",
		})

	accountResponseResult, violations := service.AuthorizationTransaction(transaction)
	if accountRespose.ActiveCard != accountResponseResult.ActiveCard {
		t.Errorf(errorTest, accountRespose.ActiveCard, accountResponseResult.ActiveCard)
	}

	if accountResponseResult.AvailableLimit != accountRespose.AvailableLimit {
		t.Errorf(errorTest, accountRespose.AvailableLimit, accountResponseResult.AvailableLimit)
	}

	if len(violations) == 0 {
		t.Errorf(errorTest, 0, len(violations))
	}

	if violations[0] != cardNotActive {
		t.Errorf(errorTest, cardNotActive, violations[0])
	}
}


func TestAuthorizationTransactionDoubleTransaction(t *testing.T){
	service := newServiceImpl(accountsvc.Inject())
	service.accountService.CleanAccounts()

	account := service.accountService.CreateAccount(
		accountdto.AccountRequest{
		ActiveCard:     true,
		AvailableLimit: 10000,
	})
	accountRespose := accountdto.NewAccountResponse(account)

	transaction1 := service.CreateTransaction(
		transactiondto.TransactionRequest{
		Merchant: "Samsung",
		Amount:   4000,
		Time:     "2019-02-13T11:00:00.000Z",
	})

	var violations []string
	accountResponseResult1, violations := service.AuthorizationTransaction(transaction1)
	if accountRespose.ActiveCard != accountResponseResult1.ActiveCard {
		t.Errorf(errorTest, accountRespose.ActiveCard, accountResponseResult1.ActiveCard)
	}

	if accountResponseResult1.AvailableLimit != accountRespose.AvailableLimit - transaction1.Amount {
		t.Errorf(errorTest, accountRespose.AvailableLimit - transaction1.Amount, accountResponseResult1.AvailableLimit)
	}

	if len(violations) != 0 {
		t.Errorf(errorTest, 0, len(violations))
	}

	transaction2 := service.CreateTransaction(
		transactiondto.TransactionRequest{
			Merchant: "Samsung",
			Amount:   4000,
			Time:     "2019-02-13T11:01:00.000Z",
		})

	accountResponseResult2, violations := service.AuthorizationTransaction(transaction2)
	if accountRespose.ActiveCard != accountResponseResult2.ActiveCard {
		t.Errorf(errorTest, accountRespose.ActiveCard, accountResponseResult2.ActiveCard)
	}

	if accountResponseResult2.AvailableLimit != accountResponseResult1.AvailableLimit {
		t.Errorf(errorTest, accountResponseResult2.AvailableLimit, accountResponseResult1.AvailableLimit)
	}

	if len(violations) == 0 {
		t.Errorf(errorTest, 0, len(violations))
	}

	if violations[0] != doubleTransaction {
		t.Errorf(errorTest, doubleTransaction, violations[0])
	}
}