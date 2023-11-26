package oo

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type Book struct {
	BookLibraryCode   string
	Title             string
	YearOfPublication int
	NumberOfPages     int
	Price             float64
	GenreID           int
	AuthorID          int
	PublisherID       int
}

type IBookCatalogDAO interface {
	AddBook(book Book) error
	FindByTitle(title string) ([]Book, error)
	FindByAuthor(authorID int) ([]Book, error)
	FindByLibraryCode(libraryCode string) (Book, error)
	UpdateBook(book Book) error
	DeleteBook(book Book) error
}

type ooDB struct {
	DB *sql.DB
}

func NewBookCatalogDAOMySQL(db *sql.DB) *ooDB {
	return &ooDB{DB: db}
}

func (db *ooDB) AddBook(book Book) error {
	// Prepare the SQL statement
	stmt, err := db.DB.Prepare("INSERT INTO bookcatalog (book_library_code, title, year_of_publication, number_of_pages, price, genre_id, author_id, publisher_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Execute the SQL statement with the provided book values
	_, err = stmt.Exec(book.BookLibraryCode, book.Title, book.YearOfPublication, book.NumberOfPages, book.Price, book.GenreID, book.AuthorID, book.PublisherID)
	if err != nil {
		return err
	}

	return nil
}

func (dao *ooDB) FindByTitle(title string) ([]Book, error) {
	// Implement the logic to find books by title in the database
	return nil, nil
}

func (dao *ooDB) FindByAuthor(authorID int) ([]Book, error) {
	// Implement the logic to find books by author in the database
	return nil, nil
}

func (dao *ooDB) FindByLibraryCode(libraryCode string) (Book, error) {
	// Implement the logic to find a book by library code in the database
	return Book{}, nil
}

func (dao *ooDB) UpdateBook(book Book) error {
	// Implement the logic to update a book in the database
	return nil
}

func (dao *ooDB) DeleteBook(book Book) error {
	// Implement the logic to delete a book from the database
	return nil
}
