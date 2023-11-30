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

	//  the table
	tablePlus := widget.NewTable(
		func() (int, int) { return 8, 2 }, // 8 rows, 2 columns
		func() fyne.CanvasObject { return widget.NewEntry() },
		func(i widget.TableCellID, obj fyne.CanvasObject) {
			// Populate the table with data
			switch i.Row {
			case 0:
				switch i.Col {
				case 0:
					obj.(*widget.Entry).SetText("Title")
				case 1:
					entry := widget.NewEntry()
					entry.SetText(data.Title)
					entry.OnChanged = func(text string) {
						data.Title = text
					}
					obj = entry
				}
			case 1:
				switch i.Col {
				case 0:
					obj.(*widget.Entry).SetText("Year of Publication")
				case 1:
					obj = widget.NewEntry()
					obj.(*widget.Entry).SetText(data.YearOfPublication)
					obj.(*widget.Entry).OnChanged = func(text string) {
						data.YearOfPublication = text
					}
				}
			case 2:
				switch i.Col {
				case 0:
					obj.(*widget.Entry).SetText("Number of Pages")
				case 1:
					obj.(*widget.Entry).SetText(data.NumberOfPages)
				}
			case 3:
				switch i.Col {
				case 0:
					obj.(*widget.Entry).SetText("Price")
				case 1:
					obj.(*widget.Entry).SetText(data.Price)
				}
			case 4:
				switch i.Col {
				case 0:
					obj.(*widget.Entry).SetText("Genre ID")
				case 1:
					obj.(*widget.Entry).SetText(data.GenreID)
				}
			case 5:
				switch i.Col {
				case 0:
					obj.(*widget.Entry).SetText("Author ID")
				case 1:
					obj.(*widget.Entry).SetText(data.AuthorID)
				}
			case 6:
				switch i.Col {
				case 0:
					obj.(*widget.Entry).SetText("Publisher ID")
				case 1:
					obj.(*widget.Entry).SetText(data.PublisherID)
				}
			case 7:
				switch i.Col {
				case 0:
					obj.(*widget.Entry).SetText("Book Library Code")
				case 1:
					obj.(*widget.Entry).SetText(data.BookLibraryCode)
				}
			default:
				obj.(*widget.Entry).SetText(fmt.Sprintf("%d %d", i.Col, i.Row))
			}
		},
	)
	tablePlus.SetColumnWidth(0, 150)
	tablePlus.SetColumnWidth(1, 150)

	tablePlus.Hide()
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
			tablePlus.Show()
		} else {
			tablePlus.Hide()
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
	tablePlus.Move(fyne.NewPos(30, 100))

	switchBtn.Resize(fyne.NewSize(130, 60))
	inputEntry.Resize(fyne.NewSize(800, 500))
	runQueryBtn.Resize(fyne.NewSize(130, 60))
	showTablesBtn.Resize(fyne.NewSize(130, 60))
	exitBtn.Resize(fyne.NewSize(130, 60))
	saveCSVBtn.Resize(fyne.NewSize(130, 60))
	tablePlus.Resize(fyne.NewSize(320, 320))

	myWindow.SetContent(
		container.NewWithoutLayout(
			nameLabel,
			inputEntry,
			runQueryBtn,
			showTablesBtn,
			saveCSVBtn,
			exitBtn,
			switchBtn,
			tablePlus,
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
