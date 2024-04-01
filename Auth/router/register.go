package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"AUTH/controllers"
	"AUTH/helper"
)

// RegisterRoutes add all routing list here automatically get main router
func RegisterRoutes(router *chi.Mux) {
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {

		response := helper.ResponseFormat{
			Message: "alive ok",
		}

		helper.SuccessResponse(w, response)
	})
	//Add All route
	router.Group(func(r chi.Router) {
		r.Post("/auth/login", controllers.LoginController)

	})

	router.HandleFunc("/*", func(w http.ResponseWriter, r *http.Request) {

		response := helper.ResponseFormat{
			Message: "route dose not exist",
			Status:  http.StatusNotFound,
		}

		helper.BadRequest(w, response)
	})

}
