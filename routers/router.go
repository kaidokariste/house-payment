package routers

import (
	"github.com/gorilla/mux"
)


func InitRoutes() *mux.Router {

	router := mux.NewRouter().StrictSlash(false)
	// Routes for welcome router
	router = SetWelcomeRoutes(router)
	// Routes for the User entity
	//router = SetUserRoutes(router)
	// Routes for the Unprotected City routes entity
	router = SetUnprotectedCityRoutes(router)
	// Routes for the TaskNote entity
	//router = SetNoteRoutes(router)
	return router
}
