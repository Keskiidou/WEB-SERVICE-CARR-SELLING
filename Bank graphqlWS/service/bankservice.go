package service

import (
	"Bank_graphqlWS/client"
	"Bank_graphqlWS/graph/model"
	"Bank_graphqlWS/repository"
	"fmt"
	"time"
)

type BankServices struct {
	Bankrepo *repository.BankRepository
}

func (s *BankServices) NewBankAcc(ID string, accountNumber string, balance float64, userid int32) (*model.Bank, error) {
	bank := model.Bank{
		ID:            ID,
		AccountNumber: accountNumber,
		Balance:       balance,
		UserID:        userid,
	}

	err := s.Bankrepo.AddBank(&bank)
	if err != nil {
		return nil, err
	}

	return &bank, nil
}
func (s *BankServices) GetBankAcc(id string) (*model.Bank, error) {
	bank, err := s.Bankrepo.GetBankByID(id)
	if err != nil {
		return nil, err
	}
	return bank, nil
}
func (s *BankServices) CanBuyCar(userId int32, carID int32) error {
	balance, err := s.Bankrepo.GetBalanceByUSERID(userId)
	if err != nil {
		return err
	}

	carPrice, err := client.FetchCarPrice(carID)
	if err != nil {
		return err
	}

	if balance >= carPrice {
		newBalance := balance - carPrice
		err := s.Bankrepo.UpdateBankBalance(userId, newBalance)
		if err != nil {
			return fmt.Errorf("failed to update balance")
		}
		cardetails, err := client.FetchCarmodel(carID)
		if err != nil {
			return fmt.Errorf("failed to fetch car details: %w", err)
		}
		purchase := &model.Purchase{
			UserID:       userId,
			CarID:        carID,
			CarDetails:   cardetails,
			Price:        int32(carPrice),
			PurchaseTime: time.Now(),
		}

		err = s.Bankrepo.AddPurchaseRecord(purchase)
		if err != nil {
			return fmt.Errorf("failed to add purchase record")
		}
		return nil
	}
	return fmt.Errorf("insufficient balance")
}
