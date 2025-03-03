package router

import (
	"bookStore/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewRouter(bookController *controller.BookController) *gin.Engine {
	service := gin.Default()

	service.GET("", func(context *gin.Context) {
		context.JSON(http.StatusOK, "welcome home")
	})

	service.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	router := service.Group("/api")
	bookRouter := router.Group("/tag")
	bookRouter.GET("", bookController.FindAll)
	bookRouter.GET("/:tagId", bookController.FindById)
	bookRouter.POST("", bookController.Create)
	bookRouter.PATCH("/:tagId", bookController.Update)
	bookRouter.DELETE("/:tagId", bookController.Delete)

	return service
}
