package models

import (
	"github.com/Adejare77/bookStore/config"
)

type Book config.Book

var db = config.GetDB() // create a custome Error

type CustomError struct {
	Message string
}

// creates a receiver pointer with method Error()
func (err *CustomError) Error() string {
	return err.Message
}

// 'Error() string' satisfies error type interface
func NewCustomError(message string) error {
	return &CustomError{Message: message} // Thus this returns error
}

func (b *Book) CreateBook() error {
	err := db.Create(b).Error

	return err
}

func GetAllBooks() ([]Book, error) {
	var books []Book
	err := db.Find(&books).Error

	return books, err
}

func GetBookById(id uint64) (Book, error) {
	var book Book = Book{}

	// err := db.Where("ID=?", id).Find(&book).Error
	err := db.Find(&book, id).Error

	if book.ID == 0 {
		return book, NewCustomError("Invalid Book ID")
	}

	return book, err
}

func UpdateBookById(id uint64, values Book) error {
	var book Book // Just as a Table

	db.Find(&book, id) // First check if the ID exists

	if book.ID == 0 {
		return NewCustomError("Invalid Book ID")
	}

	err := db.Model(&book).Where("ID=?", id).Updates(values).Error

	return err
}

func DeleteBookById(id uint64) error {
	var book, check Book // The book tells the db which table
	db.Find(&check, id)

	if check.ID == 0 {
		return NewCustomError("Invalid Book ID")
	}

	// db.Delete() uses "soft delete - by populating the deleteAt Column but not permanently delete it"
	// Useful in case there's need for recovery
	// err := db.Delete(&book, id).Error

	// db.Unscoped().Delete() permanently deletes it from the database
	err := db.Unscoped().Delete(&book, id).Error

	return err
}
