package book

import (
	"go-gin-gonic/src/database"
	"go-gin-gonic/src/models"

	"github.com/gin-gonic/gin"
)

func ShowBooks(c *gin.Context) {
	db := database.GetDatabase()

	var books []models.Book

	err := db.Find(&books).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot find books: " + err.Error(),
		})

		return
	}

	c.JSON(200, books)

}
