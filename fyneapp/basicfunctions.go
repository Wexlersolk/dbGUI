package fyneapp

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
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

func executeQuery(db *sql.DB, query string) (string, error) {
	query = strings.TrimRight(query, ";")

	rows, err := db.Query(query)
	if err != nil {
		return "", fmt.Errorf("Error executing query: %v", err)
	}
	defer func() {
		if err := rows.Close(); err != nil {
			fmt.Println("Error closing rows:", err)
		}
	}()

	columns, err := rows.Columns()
	if err != nil {
		return "", fmt.Errorf("Error executing query: %v", err)
	}

	var resultRows []string

	for rows.Next() {
		rowValues := make([]interface{}, len(columns))
		for i := range columns {
			rowValues[i] = new(sql.RawBytes)
		}

		if err := rows.Scan(rowValues...); err != nil {
			return "", fmt.Errorf("Error executing query: %v", err)
		}

		var rowString string
		for i, col := range rowValues {
			rowString += fmt.Sprintf("%s: %s, ", columns[i], string(*col.(*sql.RawBytes)))
		}

		resultRows = append(resultRows, rowString)
	}

	if err := rows.Err(); err != nil {
		return "", fmt.Errorf("Error executing query: %v", err)
	}

	if len(resultRows) == 0 {
		return "Query executed successfully, but no results found.", nil
	}

	result := strings.Join(resultRows, "\n")

	return result, nil
}

func saveCSV(filename string, data string) error {
	folderPath := "results"

	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		err := os.Mkdir(folderPath, 0755)
		if err != nil {
			return err
		}
	}

	filePath := folderPath + string(os.PathSeparator) + filename

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	lines := strings.Split(data, "\n")

	for _, line := range lines {
		record := strings.Split(line, ",")

		if err := writer.Write(record); err != nil {
			return err
		}
	}

	return nil
}

func executeCommandOO(db *sql.DB, command string, book *Book) (string, error) {
	var result string
	switch command {
	case "create":
		result = "cr result"

	case "read":
		result = "read res"

	case "update":
		result = "upd res"

	case "delete":
		result = "del res"

	}
	return result, nil
}
