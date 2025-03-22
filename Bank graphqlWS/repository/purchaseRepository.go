package repository

import (
	"Bank_graphqlWS/graph/model"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type PurchaseRepository struct {
	db *gorm.DB
}

func NewPurchaseRepository(db *gorm.DB) *PurchaseRepository {
	return &PurchaseRepository{db: db}
}

func (r *PurchaseRepository) GetUserIDAndCarIDByPurchaseID(purchaseID int32) (int32, int32, error) {
	var purchase model.Purchase

	result := r.db.Where("id = ?", purchaseID).First(&purchase)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return 0, 0, fmt.Errorf("purchase with ID %d not found", purchaseID)
		}
		return 0, 0, result.Error
	}

	return purchase.UserID, purchase.CarID, nil
}
