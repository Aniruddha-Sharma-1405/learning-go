package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	//go:generate swagger generate spec -o ./swagger/userSwagger.yaml --scan-models
	//go:generate swagger generate server -A users -f ./swagger/swagger.yaml -t ./gen

	//Users CRUD
	// router.GET("/users", controllers.GetUsers)
	// router.GET("/users/:id", controllers.GetUser)
	// router.POST("/users", controllers.CreateUser)
	// router.DELETE("/users/:id", controllers.DeleteUser)

	router.Run(":8080")

}
