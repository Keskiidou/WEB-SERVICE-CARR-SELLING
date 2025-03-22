package service

import (
	"Bank_graphqlWS/graph/model"
	"Bank_graphqlWS/insurance"
	"Bank_graphqlWS/repository"
	"fmt"
	"time"
)

type InsuranceService struct {
	insuranceRepo *repository.InsuranceRepository
	purchaseRepo  *repository.PurchaseRepository
}

func NewInsuranceService(insuranceRepo *repository.InsuranceRepository, purchaseRepo *repository.PurchaseRepository) *InsuranceService {
	return &InsuranceService{
		insuranceRepo: insuranceRepo,
		purchaseRepo:  purchaseRepo,
	}
}

// CreateInsurance creates a new insurance record
func (s *InsuranceService) CreateInsurance(purchaseID int32, id string) (*model.Insurance, error) {
	userID, carID, err := s.purchaseRepo.GetUserIDAndCarIDByPurchaseID(purchaseID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch purchase details: %v", err)
	}

	amount, companyName, err := insurance.InsuranceDetails(purchaseID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch insurance details: %v", err)
	}

	// Step 3: Create a new Insurance record
	insurance := &model.Insurance{
		ID:             id,
		PurchaseID:     purchaseID,
		UserID:         userID,
		CarID:          carID,
		DateOfContract: time.Now().Format("2006-01-02"),
		Deadline:       time.Now().AddDate(1, 0, 0).Format("2006-01-02"),
		Amount:         amount,
		CompanyName:    companyName,
	}

	if err := s.insuranceRepo.AddInsurance(insurance); err != nil {
		return nil, fmt.Errorf("failed to save insurance: %v", err)
	}

	return insurance, nil
}
