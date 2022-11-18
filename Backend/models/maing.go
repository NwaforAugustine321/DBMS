package models

import "gorm.io/gorm"

type UserModel struct {
	gorm.Model
	UserName         string `json:"userName"`
	UserEmail        string `json:"userEmail"`
	UserPhone        string `json:"userPhone"`
	Profile          string `json:"profile"`
	UserInvoiceModel []UserInvoiceModel
}

type UserInvoiceModel struct {
	gorm.Model
	Number string
	UserID uint
}
