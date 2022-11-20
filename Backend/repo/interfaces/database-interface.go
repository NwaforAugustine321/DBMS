package interfaces

import "gorm.io/gorm"



type DatabaseInterface interface{
	InitDB(dsn string)error
	DBInstance()*gorm.DB
	Migrate()error
	Get(table string,columns []string, where string, model interface{}, conditions ...interface{})interface{}
}