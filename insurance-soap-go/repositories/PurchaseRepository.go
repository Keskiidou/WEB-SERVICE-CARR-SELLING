package repositories

import (
	"gorm.io/gorm"
	"insurance-soap-go/model"
)

type PurchaseRepository struct {
	db *gorm.DB
}

func NewPurchaseRepository(db *gorm.DB) *PurchaseRepository {
	return &PurchaseRepository{db: db}
}

func (r *PurchaseRepository) FindByID(purchaseID int32) (*model.Purchase, error) {
	var purchase model.Purchase
	result := r.db.First(&purchase, purchaseID)
	if result.Error != nil {
		return nil, result.Error
	}

	return &purchase, nil
}
