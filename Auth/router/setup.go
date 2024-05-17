package router

import (
	middlewares "github.com/mars1385/e-com-go/auth/middleware"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func SetupRoute() *chi.Mux {

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Use(middlewares.Cors())
	RegisterRoutes(router) //routes register

	return router
}
