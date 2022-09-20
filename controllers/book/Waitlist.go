package book

import (
	"gobookclub/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Waitlist struct {
	BookId uint
	UserId uint
}

func AddWaiter(c *gin.Context) {
	var input models.ListInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := models.DB.Where("id = ?", input.UserId).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found!"})
		return
	}

	var book models.Book
	if err := models.DB.Where("id = ?", input.BookId).First(&book).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found!"})
		return
	}

	if !book.Borrowed {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Book available!"})
		return
	}
	book.Waitlist = append(book.Waitlist, user)
	models.DB.Model(&book).Updates(book)
	c.JSON(http.StatusOK, gin.H{"data": book})
}

func DeleteWaiter(c *gin.Context) {
	var input models.ListInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var result []Waitlist
	if err := models.DB.Raw("SELECT * FROM waitlist where user_id=? and book_id=?", input.UserId, input.BookId).Scan(&result).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User or book not on the waitlist!"})
		return
	}

	models.DB.Exec("DELETE FROM waitlist WHERE user_id=? and book_id=?", input.UserId, input.BookId)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
