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

func (db *DB) Get(table string,columns []string, where string, model interface{},conditions ...interface{}) interface{}{
	
	return databaseInstance.Get(table,columns,where,model,conditions)
	
  
}