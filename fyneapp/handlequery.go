package fyneapp

import (
	"database/sql"
	"fmt"
	"reflect"
	"strconv"

	"fyne.io/fyne/v2/widget"
	"github.com/Wexler763/dbGUI/oo"
)

const (
	numRows   = 8
	secondCol = 1
)

func atoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func atof(s string) float64 {
	f, _ := strconv.ParseFloat(s, 64)
	return f
}

func handleEditingMode(db *sql.DB, inputEntry *widget.Entry, tablePlus *widget.Table) {
	fmt.Println("reading data, stfu:")
	readDataFromTable(tablePlus)

	// Display the read data
	fmt.Println("Data from the second column:")

	// Using reflection to iterate over struct fields
	dataValue := reflect.ValueOf(data)
	dataType := dataValue.Type()
	for i := 0; i < dataValue.NumField(); i++ {
		fieldName := dataType.Field(i).Name
		fieldValue := dataValue.Field(i).Interface()
		fmt.Printf("%s: %v\n", fieldName, fieldValue)
	}

	// Assuming the first column contains the action (create, read, update, delete)
	actionColumn := 0
	if dataValue.NumField() > 0 {
		action := dataValue.Field(actionColumn).Interface()

		// Assuming the rest of the columns are book data
		bookValues := make(map[string]string)
		for i := 0; i < dataValue.NumField(); i++ {
			if i != actionColumn {
				fieldName := dataType.Field(i).Name
				fieldValue := dataValue.Field(i).Interface()
				bookValues[fieldName] = fmt.Sprintf("%v", fieldValue)
			}
		}

		// Display book values
		fmt.Println("Book Values:")
		for key, value := range bookValues {
			fmt.Printf("%s: %s\n", key, value)
		}

		// Create a book object
		book := oo.Book{
			// Assign values from the map to the book object
			// Modify this part based on your actual book structure
			BookLibraryCode:   bookValues["BookLibraryCode"],
			Title:             bookValues["Title"],
			YearOfPublication: atoi(bookValues["YearOfPublication"]),
			NumberOfPages:     atoi(bookValues["NumberOfPages"]),
			Price:             atof(bookValues["Price"]),
			GenreID:           atoi(bookValues["GenreID"]),
			AuthorID:          atoi(bookValues["AuthorID"]),
			PublisherID:       atoi(bookValues["PublisherID"]),
		}

		// Execute the command with the book object
		result, err := executeCommandOO(db, fmt.Sprintf("%v", action), book)
		if err != nil {
			resultText := fmt.Sprintf(err.Error())
			inputEntry.SetText(resultText)
			return
		}

		resultText := fmt.Sprintf(result)
		inputEntry.SetText(resultText)
	}

	// ... rest of the code
}
func readDataFromTable(table *widget.Table) {

	for row := 0; row < numRows; row++ {
		cellObject := table.CreateCell()

		if entry, ok := cellObject.(*widget.Entry); ok {
			switch row {
			case 0:
				data.Title = entry.Text
			case 1:
				data.YearOfPublication = entry.Text
				// Update other fields as needed
			}
		}
	}
}

func handleRunQuery(db *sql.DB, query string, isEditingMode bool, inputEntry *widget.Entry, tablePlus *widget.Table) {
	if isEditingMode {
		handleEditingMode(db, inputEntry, tablePlus)
		return
	}

	result, err := executeQuery(db, query)
	if err != nil {
		resultText := fmt.Sprintf(err.Error())
		inputEntry.SetText(resultText)
		return
	}

	resultText := fmt.Sprintf(result)
	inputEntry.SetText(resultText)
}
