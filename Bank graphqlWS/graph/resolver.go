package graph

import (
	"Bank_graphqlWS/graph/model"
	"Bank_graphqlWS/service"
	"context"
	"fmt"
)

type Resolver struct {
	BankService      *service.BankServices
	InsuranceService *service.InsuranceService
}

func (r *Resolver) CreateBankAccount(ctx context.Context, input model.CreateBankAccountInput) (*model.Bank, error) {
	if r.BankService == nil {
		return nil, fmt.Errorf("BankService is not initialized")
	}

	Id := input.ID
	accountNumber := input.AccountNumber
	balance := input.Balance
	userId := input.UserID

	// Logging the input values
	fmt.Printf(" ID : %s, AccountNumber: %s, Balance: %f, UserID: %d\n", Id, accountNumber, balance, userId)

	bank, err := r.BankService.NewBankAcc(Id, accountNumber, balance, userId)
	if err != nil {
		return nil, err
	}

	return &model.Bank{
		ID:            bank.ID,
		AccountNumber: bank.AccountNumber,
		Balance:       bank.Balance,
		UserID:        bank.UserID,
	}, nil
}
func (r *Resolver) GetBankAccountById(ctx context.Context, id string) (*model.Bank, error) {

	bank, err := r.BankService.Bankrepo.GetBankByID(id)
	if err != nil {
		return nil, fmt.Errorf("could not find bank account with id: %s", id)
	}

	return bank, nil
}
func (r *Resolver) CanBuyCar(ctx context.Context, carID int32, userId int32) (string, error) {
	err := r.BankService.CanBuyCar(userId, carID)
	if err != nil {
		return "", err
	}
	return "Purchase successful! Balance updated.", nil
}
func (r *Resolver) CreateInsurance(ctx context.Context, input model.CreateInsuranceInput) (*model.Insurance, error) {
	if r.InsuranceService == nil {
		return nil, fmt.Errorf("InsuranceService is not initialized")
	}
	ID := input.ID
	PurchaseID := input.PurchaseID

	// Logging the input values
	fmt.Printf(" purchaseID %d\n", PurchaseID)

	insurance, err := r.InsuranceService.CreateInsurance(PurchaseID, ID)
	if err != nil {
		return nil, err
	}

	return &model.Insurance{
		ID:             insurance.ID,
		PurchaseID:     insurance.PurchaseID,
		UserID:         insurance.UserID,
		CarID:          insurance.CarID,
		DateOfContract: insurance.DateOfContract,
		Deadline:       insurance.Deadline,
		Amount:         insurance.Amount,
		CompanyName:    insurance.CompanyName,
	}, nil
}
