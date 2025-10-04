package controllers

import (
	"learning-go-microservice/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var users = []models.User{
	{ID: 1, Name: "Alice", Age: 25},
	{ID: 2, Name: "Bob", Age: 35},
}

func GetUsers(ginContext *gin.Context) {
	ginContext.JSON(http.StatusOK, users)
}

func GetUser(ginContext *gin.Context) {
	id, _ := strconv.Atoi(ginContext.Param("id"))
	for _, u := range users {
		if u.ID == id {
			ginContext.JSON(http.StatusOK, u)
			return
		}
	}
	ginContext.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
}

func CreateUser(ginContext *gin.Context) {
	var newUser models.User
	if err := ginContext.ShouldBindBodyWithJSON(&newUser); err != nil {
		ginContext.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newUser.ID = len(users) + 1
	users = append(users, newUser)
	ginContext.JSON(http.StatusCreated, newUser)
}

func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for i, u := range users {
		if u.ID == id {
			users = append(users[:i], users[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"result": "Deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
}
