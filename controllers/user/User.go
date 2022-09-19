package user

import (
	"gobookclub/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var input models.UserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(input.Roles) == 0 {
		input.Roles = make([]models.Role, 0)
		input.Roles = append(input.Roles, models.Role{Name: "USER"})
	}
	// Create user
	user := models.User{FirstName: input.FirstName, LastName: input.LastName, Username: input.Username, Email: input.Email, Password: input.Password, Roles: input.Roles}
	models.DB.Create(&user)

	c.JSON(http.StatusOK, gin.H{"data": user})

}
func UpdateUser(c *gin.Context) {
	var user models.User
	if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input models.UserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&user).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": user})
}
func DeleteUser(c *gin.Context) {
	// Get model if exist
	var user models.User
	if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&user)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
func AddRole(c *gin.Context) {
	var user models.User
	if err := models.DB.Where("id=?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "This user does not exist!"})
	}

	var role models.RoleInput
	user.Roles = append(user.Roles, models.Role{Name: role.Name})
	models.DB.Model(&user).Updates(user)

}
func AddWish(c *gin.Context) {
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
	user.Wishlist = append(user.Wishlist, book)
	models.DB.Model(&user).Updates(user)

	c.JSON(http.StatusOK, gin.H{"data": book})
}
