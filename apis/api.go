package api

import (
	"rest-api-demo/types"
)

func NewAPIServer(address string) *types.APIServer{
	return &types.APIServer{
		Address: address,
	}
}