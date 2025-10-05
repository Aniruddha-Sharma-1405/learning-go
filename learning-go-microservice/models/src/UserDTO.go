package models

import (
	"github.com/go-openapi/strfmt"
)

type User struct {
	ID       int64        `json:"id"`
	UserName string       `json:"userName"`
	Email    strfmt.Email `json:"email"`
}
