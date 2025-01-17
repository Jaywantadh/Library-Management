package models

import (
	config "github.com/jaywantadh/go-bookstore/PKG/Config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct{
	gorm.Model
	Name 		string `gorm:""json:"name"`
	Author 		string `json:"author"`
	Publication string `json:"publication"`
}

func init(){
	config.Connnect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book{
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book{
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(ID int64) (*Book, *gorm.DB){
	var getBook Book
	db := db.Where("ID=?",ID).Find(&getBook)
	return &getBook, db
}

func DeleteBook(ID int64) Book{
	var book Book
	db.Where("ID:?", ID).Delete(book)
	return book
}



