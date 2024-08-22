package models

import (
	"golang.org/x/crypto/bcrypt"

	"github.com/jinzhu/gorm"
	"github.com/sid/go-bookstore/pkg/config"
)

var db gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

type User struct {
	gorm.Model
	Name     string `gorm:"" json:"name"`
	Age      string `json:"age"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func init() {
	config.Connect()     // establish connection
	db = *config.GetDB() // initializing the db to a variable

	db.AutoMigrate(&Book{})
	db.AutoMigrate(&User{})
}

func (u *User) CreateUser() *User {
	hashPass, _ := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)

	u.Password = string(hashPass)
	db.NewRecord(u)
	db.Create(&u)
	return u
}

func GetUserByUsername(username string) (*User, *gorm.DB) {
	var user User
	db := db.Where("Username=?", username).Find(&user)
	return &user, db
}
func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBook() []Book {
	var books []Book
	db.Find(&books)
	return books
}

func GetBookById(id int64) (*Book, *gorm.DB) {
	var book Book
	db := db.Where("ID=?", id).Find(&book)
	return &book, db
}

func DeleteBook(id int64) Book {
	var book Book
	db.Where("ID=?", id).Delete(book)
	return book
}
