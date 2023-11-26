package fyneapp

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	"fyne.io/fyne/v2/widget"
	"github.com/Wexler763/dbGUI/oo"
)

func atoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func atof(s string) float64 {
	f, _ := strconv.ParseFloat(s, 64)
	return f
}

func handleRunQuery(db *sql.DB, query string, isEditingMode bool, inputEntry *widget.Entry) {
	if isEditingMode {
		handleEditingMode(db, inputEntry)
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

func handleEditingMode(db *sql.DB, inputEntry *widget.Entry) {
	lines := strings.Split(inputEntry.Text, "\n")

	instructionLine := lines[0]
	prefix := "Enter book info and then write a command(create, read, update, delete):"

	if strings.HasPrefix(instructionLine, prefix) {
		action := strings.TrimSpace(strings.TrimPrefix(instructionLine, prefix))
		values := parseBookValues(lines)

		book := oo.Book{
			BookLibraryCode:   values["BookLibraryCode"],
			Title:             values["Title"],
			YearOfPublication: atoi(values["YearOfPublication"]),
			NumberOfPages:     atoi(values["NumberOfPages"]),
			Price:             atof(values["Price"]),
			GenreID:           atoi(values["GenreID"]),
			AuthorID:          atoi(values["AuthorID"]),
			PublisherID:       atoi(values["PublisherID"]),
		}

		result, err := executeCommandOO(db, action, book)
		if err != nil {
			resultText := fmt.Sprintf(err.Error())
			inputEntry.SetText(resultText)
			return
		}
		resultText := fmt.Sprintf(result)
		inputEntry.SetText(resultText)
	}
}

func parseBookValues(lines []string) map[string]string {
	values := make(map[string]string)

	for i := 1; i < len(lines)-1; i++ {
		line := strings.TrimSpace(lines[i])
		if line != "" {
			parts := strings.Split(line, "\t")
			if len(parts) == 2 {
				name := strings.TrimSpace(parts[0])
				value := strings.TrimSpace(parts[1])
				values[name] = value
			}
		}
	}

	return values
}
