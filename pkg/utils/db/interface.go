package db

import (
	"fmt"

	"gorm.io/gorm"
)

type DB interface {
	GetDB() *gorm.DB
	AutoMigrate(...interface{}) error
}

var dBConn DB

func GetDBConn() (DB, error) {
	if dBConn != nil {
		return dBConn, nil
	}

	dbType := "sqlite"
	var err error
	switch dbType {
	case "sqlite":
		dBConn, err = InitializeSqliteDB()
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("unsupported Database Type")
	}

	return dBConn, nil
}
