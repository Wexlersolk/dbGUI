// fyneapp.go
package fyneapp

import (
	"fmt"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func Create() {
	myApp := app.New()

	myWindow := myApp.NewWindow("TheHornedDB")

	text := widget.NewLabel("Hail the horned one ")
	text.Alignment = fyne.TextAlignCenter

	showTablesBtn := widget.NewButtonWithIcon("Show Tables", theme.ViewRefreshIcon(), func() {
		tables, err := getTables()
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
