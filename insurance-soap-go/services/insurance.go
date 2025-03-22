package services

import (
	"errors"
	"insurance-soap-go/repositories"
)

type InsuranceService struct {
	purchaseRepo *repositories.PurchaseRepository
}

func NewInsuranceService(purchaseRepo *repositories.PurchaseRepository) *InsuranceService {
	return &InsuranceService{purchaseRepo: purchaseRepo}
}

func (s *InsuranceService) CalculateInsurance(purchaseID int32) (float64, string, error) {

	purchase, err := s.purchaseRepo.FindByID(purchaseID)
	if err != nil {
		return 0, "", errors.New("purchase not found")
	}

	insuranceAmount := float64(purchase.Price) * 0.07

	insuranceCompany := " PI Insurance Co."

	return insuranceAmount, insuranceCompany, nil
}
