package models

type BookInput struct {
	Title   string `json:"title" binding:"required"`
	Author  string `json:"author" binding:"required"`
	OwnerId uint   `json:"ownerId" binding:"required"`
}

type RoleInput struct {
	Name string `json:"name" binding:"required"`
}

type UserInput struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Username  string `json:"username" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
	Roles     []Role `json:"roles"`
}

type BorrowInput struct {
	BookId     uint `json:"bookId" binding:"required"`
	BorrowerId uint `json:"borrowerId" binding:"required"`
	PeriodDays uint `json:"period" binding:"required"`
}

type ListInput struct {
	BookId uint `json:"bookId" binding:"required"`
	UserId uint `json:"userId" binding:"required"`
}
