package repository

import (
	"Bank_graphqlWS/graph/model"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type BankRepository struct {
	DB *gorm.DB
}

func NewBankRepository(db *gorm.DB) *BankRepository {
	return &BankRepository{DB: db}
}

func (repo *BankRepository) AddBank(bank *model.Bank) error {
	return repo.DB.Create(&bank).Error
}
func (repo *BankRepository) GetBankByID(id string) (*model.Bank, error) {
	var bank model.Bank
	err := repo.DB.Where("id = ?", id).First(&bank).Error
	return &bank, err
}
func (repo *BankRepository) GetBalanceByUSERID(UserID int32) (float64, error) {
	var bank model.Bank
	if err := repo.DB.Where("user_id = ?", UserID).First(&bank).Error; err != nil {
		return 0, errors.New("bank account not found")
	}

	return bank.Balance, nil
}
func (repo *BankRepository) UpdateBankBalance(userID int32, newBalance float64) error {
	return repo.DB.Model(&model.Bank{}).
		Where("user_id = ?", userID).
		Update("balance", newBalance).
		Error
}
func (r *BankRepository) AddPurchaseRecord(purchase *model.Purchase) error {
	if err := r.DB.Create(purchase).Error; err != nil {
		return fmt.Errorf("failed to insert purchase record: %v", err)
	}
	return nil
}
