package controllers

import (
	"AUTH/helper"
	"AUTH/services"
	"AUTH/types"
	"encoding/json"
	"net/http"
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
