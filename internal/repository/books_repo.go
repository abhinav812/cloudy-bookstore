package repository

import (
	"github.com/abhinav812/cloudy-bookstore/internal/model"
	"gorm.io/gorm"
)

//ListBooks - Returns all the books from book table.
func ListBooks(db *gorm.DB) (model.Books, error) {
	books := make([]*model.Book, 0)
	if err := db.Find(&books).Error; err != nil {
		return nil, err
	}

	return books, nil
}

//CreateBook - Create new book based on supplied model.Book instance.
func CreateBook(db *gorm.DB, book *model.Book) (*model.Book, error) {
	if err := db.Create(book).Error; err != nil {
		return nil, err
	}
	return book, nil
}

//ReadBook - Returns model.Book based on supplied book-id
func ReadBook(db *gorm.DB, id uint) (*model.Book, error) {
	book := &model.Book{}
	if err := db.Where("id = ?", id).First(&book).Error; err != nil {
		return nil, err
	}

	return book, nil
}

//DeleteBook - deletes book from books table based on book-id
func DeleteBook(db *gorm.DB, id uint) error {
	book := &model.Book{}
	if err := db.Where("id = ?", id).Delete(&book).Error; err != nil {
		return err
	}

	return nil
}

//UpdateBook - Update existing book based on supplied model.Book instance
func UpdateBook(db *gorm.DB, book *model.Book) error {
	if err := db.First(&model.Book{}, book.ID).Updates(book).Error; err != nil {
		return err
	}

	return nil
}
