package services

import (
	"AUTH/enums"
	"AUTH/types"
	"errors"
)

func LoginUser(req *types.LoginUser) (types.LoginResponse, error) {

	if req.Email == "D" {
		return types.LoginResponse{}, errors.New(string(enums.Invalid_Credentials))
	}
	return types.LoginResponse{
		Access:  "",
		Refresh: "",
	}, nil
}
