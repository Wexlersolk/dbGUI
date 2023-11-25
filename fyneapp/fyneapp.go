// fyneapp.go
package fyneapp

import (
	"database/sql"
	"fmt"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
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

func Create() {
	myApp := app.New()

	myWindow := myApp.NewWindow("TheHornedDB")

	text := widget.NewLabel("Hail the horned one ")
	text.Alignment = fyne.TextAlignCenter

	showTablesBtn := widget.NewButtonWithIcon("Show Tables", theme.ViewRefreshIcon(), func() {
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

		// Create a dialog to display the tables
		tablesDialog := dialog.NewInformation("Database Tables", fmt.Sprintf("Tables: %v", tables), myWindow)
		tablesDialog.SetDismissText("OK")
		tablesDialog.Show()
	})
	showTablesBtn.Importance = widget.HighImportance

	box := container.NewVBox(
		text,
		layout.NewSpacer(),
		showTablesBtn,
	)

	myWindow.SetContent(box)

	// Close the App when Escape key is pressed
	myWindow.Canvas().SetOnTypedKey(func(keyEvent *fyne.KeyEvent) {
		if keyEvent.Name == fyne.KeyEscape {
			myApp.Quit()
		}
	})

	// Show window and run app
	myWindow.ShowAndRun()
}
