package controllers

import (
	models "learning-go-microservice/models/src"
	"net/http"

	"github.com/gin-gonic/gin"
)

var users []models.User

func GetUsers() []models.User {
	return users
}

func GetUser(id int64) models.User {
	for _, u := range users {
		if u.ID == id {
			return u
		}
	}
	return models.User{}
}

func CreateUser(ginContext *gin.Context, newUser models.User) {
	if err := ginContext.ShouldBindBodyWithJSON(&newUser); err != nil {
		ginContext.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newUser.ID = int64(len(users) + 1)
	users = append(users, newUser)
	ginContext.JSON(http.StatusCreated, newUser)
}

func DeleteUser(id int64) models.User {
	for i, u := range users {
		if u.ID == id {
			users = append(users[:i], users[i+1:]...)
			return u
		}
	}
	return models.User{}
}
