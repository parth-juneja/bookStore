package repository

import model "bookStore/model"

type BooksRepository interface {
	Save(books model.Books)
	Update(books model.Books)
	Delete(bookId int)
	FindById(bookId int) (books model.Books, err error)
	FindAll() []model.Books
}
