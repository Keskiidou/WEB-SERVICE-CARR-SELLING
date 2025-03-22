package model

import "time"

type Purchase struct {
	ID           int32     `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID       int32     `gorm:"not null" json:"user_id"`
	CarID        int32     `gorm:"not null" json:"car_id"`
	CarDetails   string    `gorm:"not null" json:"car_details"`
	Price        int32     `gorm:"not null" json:"price"`
	PurchaseTime time.Time `gorm:"not null" json:"purchase_time"`
}
