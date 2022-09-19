package user

import (
	"gobookclub/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindUser(c *gin.Context) {
	var user models.User

	if err := models.DB.Preload("OwnedBooks").Preload("BorrowedBooks").Preload("Wishlist").Preload("Roles").Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}
func FindUsers(c *gin.Context) {
	var users []models.User
	models.DB.Preload("OwnedBooks").Preload("Wishlist").Preload("BorrowedBooks").Preload("Roles").Find(&users)

	c.JSON(http.StatusOK, gin.H{"users": users})
}
