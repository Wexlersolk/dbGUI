// main.go
package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Wexler763/dbGUI/dbconnect"
	"github.com/Wexler763/dbGUI/fyneapp"
)

func getTables(db *sql.DB) ([]string, error) {
	query := "SHOW TABLES"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := rows.Close(); err != nil {
			fmt.Println("Error closing rows:", err)
		}
	}()

	var tables []string
	for rows.Next() {
		var tableName string
		if err := rows.Scan(&tableName); err != nil {
			return nil, err
		}
		tables = append(tables, tableName)
	}

	return tables, nil
}

func main() {
	db, err := dbconnect.ConnectDB()
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	fyneapp.Create()

}
