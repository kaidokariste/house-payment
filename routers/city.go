package routers

import (
	//"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	//"house-payment/common"
	"house-payment/controllers"
)



/*
// SeteRoutes configures routes for protected city entry
func SetCityRoutes(router *mux.Router) *mux.Router {
	//cityRouter := mux.NewRouter()
	//cityRouter.HandleFunc("/cities", controllers.CreateCity).Methods("POST")
	//cityRouter.HandleFunc("/cities/{id}", controllers.UpdateCity).Methods("PUT")
	//cityRouter.HandleFunc("/cities/{id}", controllers.GetCityByID).Methods("GET")

	//cityRouter.HandleFunc("/cities/tasks/{id}", controllers.GetCitiesByTask).Methods("GET")
	//cityRouter.HandleFunc("/cities/{id}", controllers.DeleteCity).Methods("DELETE")
	//router.PathPrefix("/cities").Handler(negroni.New(
	//	negroni.HandlerFunc(common.Authorize),
	//	negroni.Wrap(cityRouter),
	//))
	//return router
}*/

func SetUnprotectedCityRoutes(router *mux.Router) *mux.Router  {
	router.HandleFunc("/cities", controllers.GetCities).Methods("GET")
	router.HandleFunc("/cities", controllers.CreateCity).Methods("POST")
	return router
}