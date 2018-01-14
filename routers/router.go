package routers

import (
	"github.com/gorilla/mux"
)


func InitRoutes() *mux.Router {

	router := mux.NewRouter().StrictSlash(false)
	// Routes for welcome router
	router = SetWelcomeRoutes(router)
	// Routes for the User entity
	router = SetUserRoutes(router)
	// Routes for protected city resources
	router = SetProtectedCityRoutes(router)
	return router
}
