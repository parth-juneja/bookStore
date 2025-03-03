package controller

import (
	"bookStore/data/request"
	"bookStore/data/response"
	"bookStore/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BookController struct {
	bookService service.BooksService
}

func NewBookController(service service.BooksService) *BookController {
	return &BookController{bookService: service}
}

func (controller *BookController) Create(ctx *gin.Context) {
	createBookRequest := request.CreateBooksRequest{}
	err := ctx.ShouldBindJSON(&createBookRequest)
	if err != nil {
		panic(err)
	}

	controller.bookService.Create(createBookRequest)

	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *BookController) Update(ctx *gin.Context) {
	updateBookRequest := request.UpdateBooksRequest{}
	err := ctx.ShouldBindJSON(&updateBookRequest)
	if err != nil {
		panic(err)
	}

	tagId := ctx.Param("tagId")
	id, err := strconv.Atoi(tagId)
	if err != nil {
		panic(err)
	}

	updateBookRequest.BookId = id

	controller.bookService.Update(updateBookRequest)

	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *BookController) Delete(ctx *gin.Context) {
	bookId := ctx.Param("BookId")
	id, err := strconv.Atoi(bookId)
	if err != nil {
		panic(err)
	}
	controller.bookService.Delete(id)

	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}

	ctx.JSON(http.StatusOK, webResponse)

}

func (controller *BookController) FindById(ctx *gin.Context) {
	bookId := ctx.Param("bookId")
	id, err := strconv.Atoi(bookId)
	if err != nil {
		panic(err)
	}

	bookResponse := controller.bookService.FindById(id)

	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   bookResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *BookController) FindAll(ctx *gin.Context) {
	bookResponse := controller.bookService.FindAll()

	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   bookResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}
