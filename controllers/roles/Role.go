package roles

import (
	"gobookclub/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateRole(c *gin.Context) {
	var input models.RoleInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create book
	role := models.Role{Name: input.Name}
	models.DB.Create(&role)

	c.JSON(http.StatusOK, gin.H{"data": role})
}
func UpdateRole(c *gin.Context) {
	var role models.Role
	if err := models.DB.Where("id = ?", c.Param("id")).First(&role).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input models.RoleInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&role).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": role})
}
func DeleteRole(c *gin.Context) {
	// Get model if exist
	var role models.Role
	if err := models.DB.Where("id = ?", c.Param("id")).First(&role).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&role)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
