package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type sqliteDB struct {
	db *gorm.DB
}

func (sdb *sqliteDB) GetDB() *gorm.DB {
	return sdb.db
}
func (sdb *sqliteDB) AutoMigrate(models ...interface{}) error {
	err := sdb.db.AutoMigrate(models...)
	if err != nil {
		log.Fatalf("Failed to migrate the database schema: %v", err)
	}
	return err
}

func InitializeSqliteDB() (*sqliteDB, error) {
	if _, err := os.Stat("users.db"); err == nil {
		fmt.Println("Database already exists. Skipping initialization.")
		dbConn, err := gorm.Open(sqlite.Open("users.db"), &gorm.Config{})
		if err != nil {
			log.Fatalf("Failed to connect to the database: %v", err)
			return nil, err
		}
		return &sqliteDB{
			db: dbConn,
		}, nil
	} else if os.IsNotExist(err) {
		fmt.Println("Database does not exist. Initializing database.")
		dbConn, err := gorm.Open(sqlite.Open("users.db"), &gorm.Config{})
		if err != nil {
			log.Fatalf("Failed to connect to the database: %v", err)
			return nil, err
		}
		return &sqliteDB{
			db: dbConn,
		}, nil
	} else {
		log.Fatalf("Error checking the database file: %v", err)
		return nil, err
	}
}
