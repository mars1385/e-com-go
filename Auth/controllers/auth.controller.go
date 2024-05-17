package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/mars1385/e-com-go/auth/helper"
	"github.com/mars1385/e-com-go/auth/services"
	"github.com/mars1385/e-com-go/auth/types"
)

func LoginController(w http.ResponseWriter, request *http.Request) {
	response := helper.ResponseFormat{}
	requestBody := types.LoginUser{}
	err := json.NewDecoder(request.Body).Decode(&requestBody)
	if err != nil {
		response.Message = "Invalid input"
		helper.BadRequest(w, response)
		return
	}

	res, err := services.LoginUser(&requestBody)

	if err != nil {
		response.Message = err.Error()
		helper.BadRequest(w, response)
		return
	}

	response.Message = "Request Success"
	response.Data = res
	helper.SuccessResponse(w, response)
}
