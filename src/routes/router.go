package routes

import (
	"go-gin-gonic/src/controllers/book"

	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {

	main := router.Group("api/v1")
	{
		books := main.Group("books")
		{
			books.GET("/:id", book.ShowBook)
			books.GET("/", book.ShowBooks)
			books.POST("/", book.CreateBook)
			books.PUT("/", book.UpdateBook)
			books.DELETE("/:id", book.DeleteBook)
		}
	}

	return router
}
