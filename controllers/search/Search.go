package search

import (
	"gobookclub/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Search(c *gin.Context) {
	var input models.SearchInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var result []models.Book
	if err := models.DB.Raw("SELECT * FROM books where title LIKE ? or author LIKE ?", "%"+input.Text+"%", "%"+input.Text+"%").Scan(&result).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": result})
}

func Suggestions(c *gin.Context) {
	var input models.SearchInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if len(input.Text) >= 3 {
		var result []models.Book
		if err := models.DB.Raw("SELECT * FROM books where title LIKE ? or author LIKE ?", "%"+input.Text+"%", "%"+input.Text+"%").Scan(&result).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Book not found!"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"result": result})
	}
}
