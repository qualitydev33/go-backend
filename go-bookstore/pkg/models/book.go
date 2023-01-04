package models

import (
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm: ""json: "name"`
	Author      string `json: "author"`
	Publication string `json: "publication"`
}

func (b *Book) createBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func getBooks() []Book {
	var books []Book
	db.Find(&books)
	return books
}

func getBookById(Id int64) (*Book, *gorm.DB) {
	var book Book
	db := db.Where("ID=?", Id).Find(&book)
	return &book, db
}

func delBook(Id int64) Book {
	var book Book
	db.Where("ID=?", Id).Delete(book)
	return book
}

func init() {
	config.connect()
	db = config.getDB()
	db.AutoMigrate(&Book{})
}
