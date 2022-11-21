package database

import (
	"dbms/errorsHandlers"
	"dbms/models"
	"dbms/repo/interfaces"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type database struct {
	DB *gorm.DB
}

func NewDatabase() interfaces.DatabaseInterface {
	return &database{}
}

func (db *database) InitDB(dsn string) error {
	DBInstance, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return &errorsHandlers.DataBaseError{
			Message: "Database connection error",
			Err:     err.Error(),
			Status:  500,
		}
	}

	db.DB = DBInstance

	return nil

}

func (db *database) DBInstance() *gorm.DB {
	return db.DB
}

func (db *database) Migrate()error {
	err := db.DB.AutoMigrate(
		&models.Users{},
		&models.UserInvoices{},
		
	)

	if err != nil {
		return &errorsHandlers.DataBaseError{
			Message: "Database migration failed",
			Err:     err.Error(),
			Status:  500,
		}
	}

	return nil
}

