package fyneapp

import (
	"database/sql"
	"fmt"
	"log"

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

func Create() {
	myApp := app.New()

	myWindow := myApp.NewWindow("TheHornedDB")

	nameLabel := widget.NewLabel("TheHornedDB")
	nameLabel.Alignment = fyne.TextAlignLeading

	inputEntry := widget.NewEntry()
	inputEntry.SetPlaceHolder("Enter your query here")
	inputEntry.Resize(fyne.NewSize(800, 500))

	runQueryBtn := widget.NewButton("▶️ Run Query", func() {
		// Implement the logic to run the query
		// You can use inputEntry.Text to get the user's input
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

		// Create a dialog to display the tables
		message := fmt.Sprintf("Tables: %v", tables)
		tablesDialog := dialog.NewInformation("Database Tables", message, myWindow)
		tablesDialog.SetDismissText("OK")
		tablesDialog.Show()
	})

	myWindow.SetFullScreen(true)

	nameLabel.Move(fyne.NewPos(10, 10))
	runQueryBtn.Move(fyne.NewPos(50, 600))
	inputEntry.Move(fyne.NewPos(10, 70))
	showTablesBtn.Move(fyne.NewPos(1000, 100))

	runQueryBtn.Resize(fyne.NewSize(100, 50))
	showTablesBtn.Resize(fyne.NewSize(100, 50))

	myWindow.SetContent(
		container.NewWithoutLayout(
			nameLabel,
			inputEntry,
			runQueryBtn,
			showTablesBtn,
		),
	)

	// Close the App when Escape key is pressed
	myWindow.Canvas().SetOnTypedKey(func(keyEvent *fyne.KeyEvent) {
		if keyEvent.Name == fyne.KeyEscape {
			myApp.Quit()
		}
	})

	// Show window and run app
	myWindow.ShowAndRun()
}
