package main

import (
	"bookStore/config"
	"bookStore/controller"
	"bookStore/model"
	"bookStore/repository"
	"bookStore/router"
	"bookStore/service"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
)

func main() {

	//Database
	db := config.DataBaseConnection()
	validate := validator.New()

	err := db.Table("tags").AutoMigrate(&model.Books{})
	if err != nil {
		return
	}

	//Init Repository
	bookRepository := repository.NewBooksRepositoryImpl(db)

	//Init Service
	bookService := service.NewBookServiceImpl(bookRepository, validate)

	//Init controller
	bookController := controller.NewBookController(bookService)

	//Router
	routes := router.NewRouter(bookController)

	server := &http.Server{
		Addr:           ":8888",
		Handler:        routes,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err2 := server.ListenAndServe()
	if err2 != nil {
		panic(err2)
	}

}
