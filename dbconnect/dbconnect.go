package dbconnect

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	DBUsername = "root"
	DBPassword = "MyNewPass"
	DBHost     = "127.0.0.1"
	DBPort     = "3306"
	DBName     = "Lab1"
)

var db *sql.DB

func InitDB() error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", DBUsername, DBPassword, DBHost, DBPort, DBName)
	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	return nil
}

func ConnectDB() (*sql.DB, error) {
	if db == nil {
		if err := InitDB(); err != nil {
			return nil, err
		}
	}
	return db, nil
}
