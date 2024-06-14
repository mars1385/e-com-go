package services

import (
	"errors"

	"github.com/mars1385/e-com-go/auth/enums"
	"github.com/mars1385/e-com-go/auth/types"
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
