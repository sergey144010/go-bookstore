package models

import (
	"github.com/jinzhu/gorm"
	"github.com/sergey144010/go-bookstore/pkg/config"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDb()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(id int64) (*Book, *gorm.DB) {
	var Book Book
	db := db.Where("ID=?", id).Find(&Book)
	return &Book, db
}

func DeleteBook(id int64) Book {
	var Book Book
	db.Where("ID=?", id).Delete(&Book)
	return Book
}
