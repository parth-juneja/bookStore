package service

import (
	"bookStore/data/request"
	"bookStore/data/response"
)

type BooksService interface {
	Create(books request.CreateBooksRequest)
	Update(books request.UpdateBooksRequest)
	Delete(bookId int)
	FindById(bookId int) response.BooksResponse
	FindAll() []response.BooksResponse
}
