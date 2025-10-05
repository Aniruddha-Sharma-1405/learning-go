package src

import (
	services "learning-go-microservice/services"
)

func CallGetUsers() User {
	return services.GetUsers
}
