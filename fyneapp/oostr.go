package fyneapp

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type TableData struct {
	BookLibraryCode   string
	Title             string
	YearOfPublication string
	NumberOfPages     string
	Price             string
	GenreID           string
	AuthorID          string
	PublisherID       string
}

var data TableData

type IBookCatalogDAO interface {
	AddBook(book TableData) error
	FindByTitle(title string) ([]TableData, error)
	FindByAuthor(authorID int) ([]TableData, error)
	FindByLibraryCode(libraryCode string) (TableData, error)
	UpdateBook(book TableData) error
	DeleteBook(book TableData) error
}

type ooDB struct {
	DB *sql.DB
}

func NewBookCatalogDAOMySQL(db *sql.DB) *ooDB {
	return &ooDB{DB: db}
}

func (db *ooDB) AddBook(book TableData) error {
	// Check if a book with the same library code already exists
	existingBook, err := db.findIfExists(book.BookLibraryCode)
	if err != nil {
		return err
	}

	if existingBook.BookLibraryCode == "" {
		// No existing book found, insert a new record
		stmt, err := db.DB.Prepare("INSERT INTO bookcatalog (book_library_code, title, year_of_publication, number_of_pages, price, genre_id, author_id, publisher_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?)")
		if err != nil {
			return err
		}
		defer stmt.Close()

		_, err = stmt.Exec(book.BookLibraryCode, book.Title, book.YearOfPublication, book.NumberOfPages, book.Price, book.GenreID, book.AuthorID, book.PublisherID)
		if err != nil {
			return err
		}
	} else {
		// Existing book found, update the record
		stmt, err := db.DB.Prepare("UPDATE bookcatalog SET title=?, year_of_publication=?, number_of_pages=?, price=?, genre_id=?, author_id=?, publisher_id=? WHERE book_library_code=?")
		if err != nil {
			return err
		}
		defer stmt.Close()

		_, err = stmt.Exec(book.Title, book.YearOfPublication, book.NumberOfPages, book.Price, book.GenreID, book.AuthorID, book.PublisherID, book.BookLibraryCode)
		if err != nil {
			return err
		}
	}

	return nil
}

func (dao *ooDB) FindByLibraryCode(libraryCode string) {

	row := dao.DB.QueryRow("SELECT book_library_code, title, year_of_publication, number_of_pages, price, genre_id, author_id, publisher_id FROM bookcatalog WHERE book_library_code = ?", libraryCode)

	err := row.Scan(&data.BookLibraryCode, &data.Title, &data.YearOfPublication, &data.NumberOfPages, &data.Price, &data.GenreID, &data.AuthorID, &data.PublisherID)
	if err != nil {
		fmt.Printf("error scanning")
	}

}

func (dao *ooDB) FindByTitle(title string) ([]TableData, error) {
	// Implement the logic to find books by title in the database
	return nil, nil
}

func (dao *ooDB) FindByAuthor(authorID int) ([]TableData, error) {
	// Implement the logic to find books by author in the database
	return nil, nil
}

func (dao *ooDB) UpdateBook(book TableData) error {
	// Implement the logic to update a book in the database
	return nil
}

func (dao *ooDB) DeleteBook(book TableData) error {
	// Implement the logic to delete a book from the database
	return nil
}

func (dao *ooDB) findIfExists(libraryCode string) (TableData, error) {
	var book TableData

	// Query the database to find a book by library code
	row := dao.DB.QueryRow("SELECT book_library_code, title, year_of_publication, number_of_pages, price, genre_id, author_id, publisher_id FROM bookcatalog WHERE book_library_code=?", libraryCode)

	// Scan the row into the book struct
	err := row.Scan(&book.BookLibraryCode, &book.Title, &book.YearOfPublication, &book.NumberOfPages, &book.Price, &book.GenreID, &book.AuthorID, &book.PublisherID)
	if err != nil {
		// Handle the case when the book is not found
		if err == sql.ErrNoRows {
			return TableData{}, nil
		}
		return TableData{}, err
	}

	return book, nil
}
