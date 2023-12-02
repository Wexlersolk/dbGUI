package fyneapp

import (
	"database/sql"
	"fmt"
	"strconv"

	"fyne.io/fyne/v2/widget"
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
	readDataFromTable()
	fmt.Println("Data from the second column:")

	fmt.Println("Title:", data.Title)
	fmt.Println("Year of Publication:", data.YearOfPublication)
	fmt.Println("Number of Pages:", data.NumberOfPages)
	fmt.Println("Price:", data.Price)
	fmt.Println("Genre ID:", data.GenreID)
	fmt.Println("Author ID:", data.AuthorID)
	fmt.Println("Publisher ID:", data.PublisherID)
	fmt.Println("Book Library Code:", data.BookLibraryCode)

}
func readDataFromTable() {

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
