package db

import (
	"os"

	"github.com/api-assignment/pkg/utils/logger"
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
	log := logger.InitializeAppLogger()
	err := sdb.db.AutoMigrate(models...)
	if err != nil {
		log.Panicf("Failed to migrate the database schema: %v", err)
	}
	return err
}

func InitializeSqliteDB() *sqliteDB {
	log := logger.InitializeAppLogger()
	if _, err := os.Stat("users.db"); err == nil {
		log.Info("Database already exists. Skipping initialization.")
		dbConn, err := gorm.Open(sqlite.Open("users.db"), &gorm.Config{})
		if err != nil {
			log.Fatalf("Failed to connect to the database: %v", err)
			return nil
		}
		return &sqliteDB{
			db: dbConn,
		}
	} else if os.IsNotExist(err) {
		log.Info("Database does not exist. Initializing database.")
		dbConn, err := gorm.Open(sqlite.Open("users.db"), &gorm.Config{})
		if err != nil {
			log.Fatalf("Failed to connect to the database: %v", err)
			return nil
		}
		return &sqliteDB{
			db: dbConn,
		}
	} else {
		log.Fatalf("Error checking the database file: %v", err)
		return nil
	}
}
