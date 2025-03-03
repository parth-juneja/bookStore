package service

import (
	"bookStore/data/request"
	"bookStore/data/response"
	"bookStore/model"
	"bookStore/repository"
	"github.com/go-playground/validator/v10"
)

type BooksServiceImpl struct {
	BookRepository repository.BooksRepository
	Validate       *validator.Validate
}

func NewBookServiceImpl(tagRepository repository.BooksRepository, validate *validator.Validate) BooksService {
	return &BooksServiceImpl{
		BookRepository: tagRepository,
		Validate:       validate,
	}
}

func (b BooksServiceImpl) Create(book request.CreateBooksRequest) {
	err := b.Validate.Struct(book)
	if err != nil {
		panic(err)
	}
	//helper.ErrorPanic(err)
	bookModel := model.Books{
		BookName: book.BookName,
	}
	b.BookRepository.Save(bookModel)
}

func (b BooksServiceImpl) Update(book request.UpdateBooksRequest) {
	bookData, err := b.BookRepository.FindById(book.BookId)
	if err != nil {
		panic(err)
	}
	bookData.BookName = book.BookName
	b.BookRepository.Update(bookData)
}

func (b BooksServiceImpl) Delete(bookId int) {
	b.BookRepository.Delete(bookId)
}

func (b BooksServiceImpl) FindById(bookId int) response.BooksResponse {
	bookData, err := b.BookRepository.FindById(bookId)
	//helper.ErrorPanic(err)
	if err != nil {
		panic(err)
	}
	bookResponse := response.BooksResponse{
		BookId:   bookData.BookId,
		BookName: bookData.BookName,
	}
	return bookResponse
}

func (b BooksServiceImpl) FindAll() []response.BooksResponse {
	result := b.BookRepository.FindAll()

	var books []response.BooksResponse
	for _, value := range result {
		book := response.BooksResponse{
			BookId:   value.BookId,
			BookName: value.BookName,
		}
		books = append(books, book)
	}
	return books
}
