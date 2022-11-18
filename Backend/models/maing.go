package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Users struct {
	UserID           uuid.UUID      `json:"userID" gorm:"size:255"`
	UserName         string         `json:"userName"`
	UserEmail        string         `json:"userEmail"`
	UserPhone        string         `json:"userPhone"`
	Profile          string         `json:"profile"`
	UserInvoiceModel []UserInvoices `gorm:"foreignKey:UserID"`
	gorm.Model
}

type UserInvoices struct {
	UserID   uuid.UUID `json:"userID" gorm:"size:255"`
	ItemName string    `json:"Name"`
	UserName string    `json:"userName"`
	Profile  string    `json:"profile"`
	Rate     int64     `json:"Rate"`
	Hour     int64     `json:"hour"`
	gorm.Model
}

