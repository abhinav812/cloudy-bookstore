package repository

import (
	"github.com/abhinav812/cloudy-bookstore/internal/model"
	"gorm.io/gorm"
)

//Repo - generic contract for handling CRUD operation on model.Book
type Repo interface {
	ReadBook(db *gorm.DB, id uint) (book *model.Book, err error)

	DeleteBook(db *gorm.DB, id uint) (err error)

	CreateBook(db *gorm.DB, book *model.Book) (createdBook *model.Book, err error)

	UpdateBook(db *gorm.DB, book *model.Book) (err error)

	ListBooks(db *gorm.DB) (books model.Books, err error)
}
