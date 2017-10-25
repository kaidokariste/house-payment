package routers

import (
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
	"log"
)

func index(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Welcome!")
	log.Println("Inside index welcome")
}

// Set Welcome router for routing testing
func SetWelcomeRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/welcome", index).Methods("GET")
	return router
}

