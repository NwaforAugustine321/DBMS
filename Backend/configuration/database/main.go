package database

import "dbms/repo/interfaces"


type database struct{}

func NewDatabase()interfaces.DatabaseInterface{
	return &database{}
}