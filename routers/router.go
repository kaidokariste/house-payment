package routers

import (
	"github.com/gorilla/mux"
	"log"
)


func InitRoutes() *mux.Router {

	router := mux.NewRouter().StrictSlash(false)
	// Routes for welcome router
	router = SetWelcomeRoutes(router)
	// Routes for the User entity
	//router = SetUserRoutes(router)
	// Routes for the Task entity
	//router = SetCityRoutes(router)
	log.Print("SetCityRoutes are:", router)
	// Routes for the TaskNote entity
	//router = SetNoteRoutes(router)
	return router
}
