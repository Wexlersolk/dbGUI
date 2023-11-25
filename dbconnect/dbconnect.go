// dbconnector.go
package dbconnect

import (
	"database/sql"
	"fmt"
	"log"

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

func InitDB() (*sql.DB, error) {
	// Modify the DSN as needed
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", DBUsername, DBPassword, DBHost, DBPort, DBName)
	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func ConnectDB() {
	_, err := InitDB()
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
}
