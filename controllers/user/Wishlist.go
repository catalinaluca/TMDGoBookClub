package user

import (
	"gobookclub/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

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

type Result struct {
	UserId    uint
	BookId    uint
	FirstName string
	LastName  string
	Username  string
	Title     string
	Author    string
}
type Wishlist struct {
	BookId uint
	UserId uint
}

func FindWishes(c *gin.Context) {
	var result []Result
	err := models.DB.Raw("SELECT wishlist.user_id,wishlist.book_id,users.first_name,users.last_name, users.username,books.title,books.author FROM wishlist left join users on wishlist.user_id=users.id left join books on books.id=wishlist.book_id where user_id=?", c.Param("userId")).Scan(&result).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result:": result})
}

func DeleteWish(c *gin.Context) {
	var input models.ListInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var result []Wishlist
	if err := models.DB.Raw("SELECT * FROM wishlist where user_id=? and book_id=?", input.UserId, input.BookId).Scan(&result).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User or book not on the waitlist!"})
		return
	}

	models.DB.Exec("DELETE FROM wishlist WHERE user_id=? and book_id=?", input.UserId, input.BookId)
	c.JSON(http.StatusOK, gin.H{"data": true})

}
