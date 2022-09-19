package book

import (
	"gobookclub/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func BorrowBook(c *gin.Context) {
	var input models.BorrowInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := models.DB.Where("id = ?", input.BorrowerId).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found!"})
		return
	}

	var book models.Book
	if err := models.DB.Where("id = ?", input.BookId).First(&book).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found!"})
		return
	}
	book.Borrowed = true
	book.StartDate = time.Now()
	book.EndDate = time.Now().Add(time.Hour * 24 * time.Duration(input.PeriodDays))
	book.BorrowerId = user.ID
	models.DB.Model(&book).Updates(book)
	c.JSON(http.StatusOK, gin.H{"data": book})
}

func ExtendRentingPeriod(c *gin.Context) {
	var input models.BorrowInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var book models.Book
	if err := models.DB.Where("id=?", input.BookId).First(&book).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found!"})
		return
	}
	book.EndDate = book.EndDate.Add(time.Hour * 24 * time.Duration(input.PeriodDays))
	models.DB.Model(&book).Updates(book)

	c.JSON(http.StatusOK, gin.H{"data": book})
}
