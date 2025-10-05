package middleware

import (
	generatedModels "learning-go-microservice/controllers/target/models"
	model "learning-go-microservice/models/src"
)

// Convert your model to generated User
func ToGeneratedUserArr(input []model.User) []*generatedModels.User {
	output := make([]*generatedModels.User, len(input))
	for i, u := range input {
		output[i] = ToGeneratedUser(u)
	}
	return output
}

func ToGeneratedUser(user model.User) *generatedModels.User {
	return &generatedModels.User{
		ID:       user.ID,
		UserName: user.UserName,
		Email:    user.Email,
	}
}

// Convert generated User to your model
func FromGeneratedUser(u *generatedModels.User) model.User {
	return model.User{
		ID:       u.ID,
		UserName: u.UserName,
		Email:    u.Email,
	}
}
