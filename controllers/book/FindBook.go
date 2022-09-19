package book

import (
	"gobookclub/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type result struct {
	Id         uint
	Title      string
	Author     string
	OwnerId    uint
	Borrowed   bool
	StartDate  time.Time
	EndDate    time.Time
	BorrowerId uint
	FirstName  string
	LastName   string
	Username   string
	Email      string
}

func FindBook(c *gin.Context) { // Get model if exist
	var book models.Book

	if err := models.DB.Preload("Waitlist").Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
}
func FindBooks(c *gin.Context) {
	var books []models.Book
	models.DB.Preload("Waitlist").Find(&books)

	c.JSON(http.StatusOK, gin.H{"data": books})
}

func OwnedBooks(c *gin.Context) {

	var result []result
	err := models.DB.Model(&models.Book{}).Select("books.id,books.owner_id,books.borrower_id,books.title,books.author,books.borrowed,books.start_date,books.end_date,users.first_name,users.last_name,users.username,users.email").Joins("left join users on books.borrower_id=users.id").Where("books.owner_id=?", c.Param("ownerId")).Scan(&result).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"owned books": result})

}
func AvailableBooks(c *gin.Context) {
	var availableBooks []models.Book
	models.DB.Where("borrowed=?", false).Find(&availableBooks)

	c.JSON(http.StatusOK, gin.H{"data": availableBooks})
}
