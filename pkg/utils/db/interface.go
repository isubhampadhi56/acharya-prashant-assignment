package db

import (
	"log"

	"gorm.io/gorm"
)

type DB interface {
	GetDB() *gorm.DB
	AutoMigrate(...interface{}) error
}

var dBConn DB

func GetDBConn() DB {
	if dBConn != nil {
		return dBConn
	}

	dbType := "sqlite"
	switch dbType {
	case "sqlite":
		dBConn = InitializeSqliteDB()
	default:
		log.Fatalf("invalid database type %s", dbType)
	}

	return dBConn
}
