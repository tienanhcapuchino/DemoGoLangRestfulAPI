package stores

import (
	"rest-api-demo/types"
)

type UserStore interface {
	GetAllUsers() (*types.UserModel, error)
	CreateUser(user types.UserModel) error
}