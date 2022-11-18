package interfaces

import "gorm.io/gorm"


type DatabaseInterface interface{
	InitDB(dsn string)error
	DBInstance()*gorm.DB
}