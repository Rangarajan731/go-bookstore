package models

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Rangarajan731/go-bookstore/pkg/app"
	"github.com/Rangarajan731/go-bookstore/pkg/utils"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Book struct {
	Id     int    `json:"id" gorm:"primaryKey"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Desc   string `json:"desc"`
}

var db *gorm.DB

func init() {
	db = app.CreateConnection("postgres://postgres:rang97@localhost:5432/postgres")
	db.AutoMigrate(&Book{})
}

func GetBook(r *http.Request) *Book {
	var book Book
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	db.First(&book, id)
	return &book
}

func GetBooks(r *http.Request) *[]Book {
	var books []Book
	db.Find(&books)
	return &books
}

func AddBook(r *http.Request) *Book {
	var book Book
	utils.ParseBody(r, &book)
	fmt.Println(book)
	db.Create(&book)
	return &book
}

func DeleteBook(r *http.Request) *Book {
	var book Book
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	db.First(&book, id)
	db.Delete(&book)
	return &book
}

func UpdateBook(r *http.Request) *Book {
	var book Book
	utils.ParseBody(r, &book)
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	book.Id = id
	db.Save(&book)
	return &book
}
