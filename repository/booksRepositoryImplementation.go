package repository

import (
	"bookStore/data/request"
	"bookStore/model"
	"errors"
	"gorm.io/gorm"
)

type BooksRepositoryImpl struct {
	Db *gorm.DB
}

func NewBooksRepositoryImpl(Db *gorm.DB) BooksRepository {
	return &BooksRepositoryImpl{Db: Db}
}

func (b BooksRepositoryImpl) Save(books model.Books) {
	result := b.Db.Create(&books)
	if result.Error != nil {
		panic(result.Error)
	}
}

// Update import request
func (b BooksRepositoryImpl) Update(books model.Books) {
	var updateBook = request.UpdateBooksRequest{
		BookId:   books.BookId,
		BookName: books.BookName,
		Author:   books.Author,
	}
	result := b.Db.Model(&books).Updates(updateBook)
	if result.Error != nil {
		panic(result.Error)
	}
}

func (b BooksRepositoryImpl) Delete(bookId int) {
	var book model.Books
	result := b.Db.Where("id = ?", bookId).Delete(&book)
	if result.Error != nil {
		panic(result.Error)
	}
}

func (b BooksRepositoryImpl) FindById(tagsId int) (model.Books, error) {
	var book model.Books
	result := b.Db.Find(&book, tagsId)
	if result != nil {
		return book, nil
	} else {
		return book, errors.New("tag is not found")
	}
}

func (b BooksRepositoryImpl) FindAll() []model.Books {
	var books []model.Books
	results := b.Db.Find(&books)
	if results.Error != nil {
		panic(results.Error)
	}
	return books
}
