package repository

import (
	"Bank_graphqlWS/graph/model"
	"gorm.io/gorm"
)

type InsuranceRepository struct {
	DB *gorm.DB
}

func NewInsuranceRepository(db *gorm.DB) *InsuranceRepository {
	return &InsuranceRepository{
		DB: db,
	}
}
func (repo *InsuranceRepository) AddInsurance(insurance *model.Insurance) error {
	return repo.DB.Create(&insurance).Error
}
