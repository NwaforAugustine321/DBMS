package databaseConfig

import (
	"dbms/repo/database"
	"dbms/repo/interfaces"

	"gorm.io/gorm"
)

type db struct{}

var databaseInstance = database.NewDatabase()

func NewDatabase() interfaces.DatabaseInterface {
	return &db{}
}

func (db *db) InitDB(dsn string) error {
	return databaseInstance.InitDB(dsn)
}

func (db *db) DBInstance() *gorm.DB {
	return databaseInstance.DBInstance()
}

func (db *db) Migrate() error {
	return databaseInstance.Migrate()
}
