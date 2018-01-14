package routers

import (
	//"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	//"house-payment/common"
	"house-payment/controllers"
	"github.com/codegangsta/negroni"
	"house-payment/common"
)

// Set Protected Routes for cityesconfigures routes for protected city entry
func SetProtectedCityRoutes(router *mux.Router) *mux.Router {
	cityRouter := mux.NewRouter()
	cityRouter.HandleFunc("/cities", controllers.CreateCity).Methods("POST")
	cityRouter.HandleFunc("/cities/{id}", controllers.DeleteCity).Methods("DELETE")
	cityRouter.HandleFunc("/cities", controllers.GetCities).Methods("GET")
	router.PathPrefix("/cities").Handler(negroni.New(
		negroni.HandlerFunc(common.Authorize),
		negroni.Wrap(cityRouter),
	))
	return router
}
