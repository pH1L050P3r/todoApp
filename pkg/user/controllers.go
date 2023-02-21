package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pH1L050P3r/todoApp/pkg/utils"
)

func GetUserById(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "END POINT IS WORKING"})
}

func CreateUser(c *gin.Context) {
	var userInput CreateUserRequest
	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(userInput.Password) < 8 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "password length must be greater than or equal to 8"})
		return
	}

	if userInput.Password != userInput.ConfirmPassword {
		c.JSON(http.StatusBadRequest, gin.H{"error": "password and confirm_password is not matched"})
		return
	}

	_, err := GetUserByUsername(userInput.Username)
	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user with username already exists"})
		return
	}

	hashPassword, _ := utils.HashPassword(userInput.Password)
	user := User{
		Username:  userInput.Username,
		FirstName: userInput.FirstName,
		LastName:  userInput.LastName,
		Password:  hashPassword,
	}

	err = InsertUser(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"detail": "User created successfully"})
}
