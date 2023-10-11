package helper

import (
	"encoding/json"
	"net/http"
)

type ResponseFormat struct {
	Message string
	Data    interface{}
}

func SuccessResponse(w http.ResponseWriter, result ResponseFormat) interface{} {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	return json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"message": &result.Message,
		"result":  &result.Data,
	})
}

func BadRequest(w http.ResponseWriter, result ResponseFormat) interface{} {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	return json.NewEncoder(w).Encode(map[string]interface{}{
		"success": false,
		"message": &result.Message,
		"result":  &result.Data,
	})
}

func InternalServerError(w http.ResponseWriter, result ResponseFormat) interface{} {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	return json.NewEncoder(w).Encode(map[string]interface{}{
		"success": false,
		"message": &result.Message,
		"result":  &result.Data,
	})
}
