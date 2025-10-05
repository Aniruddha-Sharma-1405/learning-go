package src

import (
	model "learning-go-microservice/controllers/target/models"
	operations "learning-go-microservice/controllers/target/restapi/operations"
	mapper "learning-go-microservice/middleware"
	services "learning-go-microservice/services"
	"log"

	"github.com/go-openapi/runtime/middleware"
)

// Sample in-memory store
var users = map[int64]*model.User{}
var nextID int64 = 1

type UserAPIHandler struct{}

func (h *UserAPIHandler) GetUsers(params operations.GetUsersParams) middleware.Responder {
	log.Print("Received Request to get all users from the system")
	users := services.GetUsers
	responseUsers := mapper.ToGeneratedUserArr(users())
	return operations.NewGetUsersOK().WithPayload(responseUsers)
}
