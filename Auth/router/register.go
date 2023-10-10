package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"E-COM/helper"
)

// RegisterRoutes add all routing list here automatically get main router
func RegisterRoutes(router *chi.Mux) {
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		helper.SuccessResponse(w, "alive ok")
	})
	//Add All route

}
