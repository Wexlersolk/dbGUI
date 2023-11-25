package fyneapp

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"

	"github.com/Wexler763/dbGUI/dbconnect"
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
	file, err := os.Create(filename)
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

func Create() {
	myApp := app.New()

	myWindow := myApp.NewWindow("TheHornedDB")

	nameLabel := widget.NewLabel("TheHornedDB")
	nameLabel.Alignment = fyne.TextAlignLeading

	inputEntry := widget.NewEntry()
	inputEntry.SetPlaceHolder("Enter your query here")
	inputEntry.MultiLine = true

	runQueryBtn := widget.NewButton("▶️ Run Query", func() {
		db, err := dbconnect.ConnectDB()
		if err != nil {
			log.Println("Failed to connect to the database:", err)
			return
		}

		query := inputEntry.Text
		if query == "" {
			return
		}

		result, err := executeQuery(db, query)
		if err != nil {
			log.Println("Error executing query:", err)
			return
		}

		resultText := fmt.Sprintf(result)
		inputEntry.SetText(resultText)
	})

	showTablesBtn := widget.NewButton("Show Tables", func() {
		db, err := dbconnect.ConnectDB()
		if err != nil {
			log.Println("Failed to connect to the database:", err)
			return
		}

		tables, err := getTables(db)
		if err != nil {
			log.Println("Error fetching tables:", err)
			return
		}

		message := fmt.Sprintf("Tables: %v", tables)
		tablesDialog := dialog.NewInformation("Database Tables", message, myWindow)
		tablesDialog.SetDismissText("OK")
		tablesDialog.Show()
	})

	saveCSVBtn := widget.NewButton("Save CSV", func() {
		filenameEntry := widget.NewEntry()
		filenameEntry.SetPlaceHolder("Enter filename (without .csv)")

		form := &widget.Form{
			OnSubmit: func() {
				filename := filenameEntry.Text + ".csv"
				content := inputEntry.Text

				err := saveCSV(filename, content)
				if err != nil {
					log.Println("Error saving CSV:", err)
					inputEntry.SetText("Error saving CSV")
					return
				}

				dialog.NewInformation("CSV Saved", fmt.Sprintf("Content has been saved to %s", filename), myWindow).Show()
			},
			OnCancel: func() {},
			Items: []*widget.FormItem{
				{Text: "Filename", Widget: filenameEntry},
			},
		}

		dialog.ShowCustom("Save CSV", "Save", form, myWindow)
	})

	exitBtn := widget.NewButton("Exit", func() {
		myApp.Quit()
	})

	nameLabel.Move(fyne.NewPos(10, 10))
	runQueryBtn.Move(fyne.NewPos(50, 600))
	inputEntry.Move(fyne.NewPos(10, 70))
	showTablesBtn.Move(fyne.NewPos(1000, 100))
	exitBtn.Move(fyne.NewPos(1000, 500))
	saveCSVBtn.Move(fyne.NewPos(1000, 300))

	inputEntry.Resize(fyne.NewSize(800, 500))
	runQueryBtn.Resize(fyne.NewSize(130, 60))
	showTablesBtn.Resize(fyne.NewSize(130, 60))
	exitBtn.Resize(fyne.NewSize(130, 60))
	saveCSVBtn.Resize(fyne.NewSize(130, 60))

	myWindow.SetContent(
		container.NewWithoutLayout(
			nameLabel,
			inputEntry,
			runQueryBtn,
			showTablesBtn,
			saveCSVBtn,
			exitBtn,
		),
	)
	myWindow.SetFullScreen(true)

	myWindow.Canvas().SetOnTypedKey(func(keyEvent *fyne.KeyEvent) {
		if keyEvent.Name == fyne.KeyEscape {
			myApp.Quit()
		}
	})

	myWindow.ShowAndRun()
}
