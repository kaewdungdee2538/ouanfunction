package pg_db

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// var db *gorm.DB

// func GetDB() *gorm.DB {
// 	return db
// }

func SetupDB(dbHost string, dbUserName string, dbPassword string, dbName string, dbPort string) (*gorm.DB,error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Bangkok", dbHost, dbUserName, dbPassword, dbName, dbPort)
	// fmt.Println(dsn)
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil,errors.New("connect to database failed")
	}
	sqlDB, err := database.DB()
	if err != nil {
		return nil,errors.New("get generic database failed")
	}
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	return database,nil
}

func SaveTransactionDB(db *gorm.DB,query string, value map[string]interface{}) error {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Exec(query, value).Error; err != nil {
		return err
	}

	cmt := tx.Commit().Error
	if cmt != nil {
		return cmt
	} else {
		return nil
	}
}

func GetTransactionWithValueDB(db *gorm.DB,query string, result interface{}, value interface{}) error {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}

	// Raw SQL
	rows, err := tx.Raw(query, value).Rows()

	if err != nil {
		defer rows.Close()
		tx.Rollback()
		return err
	} else {
		defer rows.Close()
		for rows.Next() {
			rows.Scan(result)
			tx.ScanRows(rows, result)
			// do something
		}

		cmt := tx.Commit().Error
		if cmt != nil {
			return cmt
		} else {
			return nil
		}
	}
}


func GetTransactionNotalueDB(db *gorm.DB,query string, result interface{}) error {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}

	// Raw SQL
	rows, err := tx.Raw(query).Rows()

	if err != nil {
		defer rows.Close()
		tx.Rollback()
		return err
	} else {
		defer rows.Close()
		for rows.Next() {
			rows.Scan(result)
			tx.ScanRows(rows, result)
			// do something
		}

		cmt := tx.Commit().Error
		if cmt != nil {
			return cmt
		} else {
			return nil
		}
	}
}
