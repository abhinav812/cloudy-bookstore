package model

import (
	"gorm.io/gorm"
	"time"
)

//Books - array of Book
type Books []*Book

// Book - represents books table
type Book struct {
	gorm.Model
	Title         string
	Author        string
	PublishedDate time.Time
	ImageURL      string `gorm:"column:image_url"`
	Description   string
	CreatedAt     time.Time
}

//BookDtos - array of BookDto
type BookDtos []*BookDto

//BookDto - represents Book http response json
type BookDto struct {
	ID            uint   `json:"id"`
	Title         string `json:"title"`
	Author        string `json:"author"`
	PublishedDate string `json:"published_date"`
	ImageURL      string `json:"image_url"`
	Description   string `json:"description"`
}

//ToDto - Converts Book database model to BooksDto http json response object
func (b Book) ToDto() *BookDto {
	return &BookDto{
		ID:            b.ID,
		Title:         b.Title,
		Author:        b.Author,
		PublishedDate: b.PublishedDate.Format("2006-01-02"),
		ImageURL:      b.ImageURL,
		Description:   b.Description,
	}
}

//ToDto - Converts Books objects to BookDtos http json response object
func (bs Books) ToDto() BookDtos {
	dtos := make([]*BookDto, len(bs))
	for i, b := range bs {
		dtos[i] = b.ToDto()
	}
	return dtos
}
