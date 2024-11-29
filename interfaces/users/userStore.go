package stores

import (
	"rest-api-demo/types"
)

type UserStore interface {
	GetAllUsers() ([]types.UserModel, error)
}