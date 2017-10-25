package routers

import (
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
)


func index(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Welcome!")
}

// Set Welcome router for routing testing
func SetWelcomeRoutes(router *mux.Router) *mux.Router {
	welcomeRouter := mux.NewRouter()
	welcomeRouter.HandleFunc("/welcome", index).Methods("GET")
	return router
}

