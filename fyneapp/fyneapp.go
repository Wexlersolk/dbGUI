package fyneapp

import (
	"fmt"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/Wexler763/dbGUI/dbconnect"
)

func Create() {

	myApp := app.New()

	myWindow := myApp.NewWindow("TheHornedDB")

	isEditingMode := false

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

		handleRunQuery(db, query, isEditingMode, inputEntry)
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

		submitBtn := widget.NewButton("Save", func() {
			filename := filenameEntry.Text + ".csv"
			content := inputEntry.Text

			err := saveCSV(filename, content)
			if err != nil {
				log.Println("Error saving CSV:", err)
				inputEntry.SetText("Error saving CSV")
				return
			}

			infoDialog := dialog.NewInformation("CSV Saved", fmt.Sprintf("Content has been saved to %s", filename), myWindow)
			infoDialog.Show()
		})

		formContainer := container.NewWithoutLayout(
			widget.NewLabel("Filename:"),
			filenameEntry,
			submitBtn,
		)

		saveDialog := dialog.NewCustom("Save CSV", "Exit", formContainer, myWindow)
		filenameEntry.Move(fyne.NewPos(100, 0))
		submitBtn.Move(fyne.NewPos(240, 85))

		filenameEntry.Resize(fyne.NewSize(230, 40))
		submitBtn.Resize(fyne.NewSize(50, 40))
		saveDialog.Resize(fyne.NewSize(400, 200))

		saveDialog.Show()
	})

	switchBtn := widget.NewButton("Editing Mode", func() {
		isEditingMode = !isEditingMode
		if isEditingMode {
			inputEntry.SetText("Enter book info and then write a command(create, read, update, delete):\n" +
				"BookLibraryCode\t\n" +
				"Title\t\n" +
				"YearOfPublication\t\n" +
				"NumberOfPages\t\n" +
				"Price\t\n" +
				"GenreID\t\n" +
				"AuthorID\t\n" +
				"PublisherIDt\t\n")
		} else {
			inputEntry.SetText("")
		}
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
	switchBtn.Move(fyne.NewPos(600, 600))

	switchBtn.Resize(fyne.NewSize(130, 60))
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
			switchBtn,
		),
	)

	myWindow.Canvas().SetOnTypedKey(func(keyEvent *fyne.KeyEvent) {
		if keyEvent.Name == fyne.KeyEscape {
			myApp.Quit()
		}
	})

	myWindow.SetFullScreen(true)

	myWindow.ShowAndRun()
}
