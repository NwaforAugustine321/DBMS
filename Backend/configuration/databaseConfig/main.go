package databaseConfig

import (
	"dbms/repo/database"
	"dbms/repo/interfaces"

	"gorm.io/gorm"
)

type DB struct{}

var databaseInstance = database.NewDatabase()

func NewDatabase() interfaces.DatabaseInterface {
	return &DB{}
}

func (db *DB) InitDB(dsn string) error {
	return databaseInstance.InitDB(dsn)
}

func (db *DB) DBInstance() *gorm.DB {
	return databaseInstance.DBInstance()
}

func (db *DB) Migrate() error {
	return databaseInstance.Migrate()
}

