// main.go
package main

import (
	"fmt"

	"github.com/Wexler763/dbGUI/dbconnect"
)

func getTables() ([]string, error) {
	query := "SHOW TABLES"
	rows, err := dbconnect.DB().Query(query)
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
	dbconnect.ConnectDB()
	tables, err := getTables()
	if err != nil {
		fmt.Println("Error getting tables:", err)
		return
	}

	fmt.Println("Tables:", tables)
}
