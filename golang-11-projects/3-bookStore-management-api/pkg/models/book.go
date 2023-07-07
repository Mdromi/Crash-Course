package models

import (
	"fmt"

	"github.com/Mdromi/Crash-Course/golang-11-projects/3-bookStore-management-api/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"column:name" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	result := db.Create(&b)
	if result.Error != nil {
		fmt.Println(result)
	}
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(Id int64) Book {
	// book := &Book{}
	// db.Where("ID=?", Id).Delete(book)
	// return *book

	var book Book
	db.Where("ID=?", Id).Delete(&book)
	return book
}
