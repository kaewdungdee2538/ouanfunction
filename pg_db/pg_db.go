package pg_db

import (
	"errors"
	"fmt"
	"time"

	ouan_time "github.com/kaewdungdee2538/ouanfunction/v2/time"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// var db *gorm.DB

// func GetDB() *gorm.DB {
// 	return db
// }

/*
args[0] is conMaxIdle default equal 0 is pooling all time
args[1] is conMaxOpenConns default equal 100
*/
func SetupDB(dbHost string, dbUserName string, dbPassword string, dbName string, dbPort string, args ...int) (*gorm.DB, error) {

	conMaxIdle := ouan_time.ConvertIntToDuration(0)
	if len(args) > 0 {
		conMaxIdle = ouan_time.ConvertIntToDuration(args[0])
	}

	conMaxOpenConns := 100
	if len(args) > 1 {
		conMaxOpenConnsInput := args[1]
		if conMaxOpenConnsInput > 0 {
			conMaxOpenConns = conMaxOpenConnsInput
		}
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Bangkok", dbHost, dbUserName, dbPassword, dbName, dbPort)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, errors.New("connect to database failed")
	}
	sqlDB, err := database.DB()
	if err != nil {
		return nil, errors.New("get generic database failed")
	}
	// SetConnMaxIdleTime sets the maximum amount of time a connection may be idle.
	sqlDB.SetConnMaxIdleTime(conMaxIdle)
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(conMaxOpenConns)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	return database, nil
}

func SaveTransactionDB(db *gorm.DB, query string, value map[string]interface{}) error {
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

func GetTransactionWithValueDB(db *gorm.DB, query string, result interface{}, value interface{}) error {
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

func GetTransactionNoneValueDB(db *gorm.DB, query string, result interface{}) error {
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
