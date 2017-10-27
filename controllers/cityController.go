package controllers

import (
	"encoding/json"
	"net/http"

	"house-payment/data"

	"log"
)

// Handler for HTTP Post - "/cities"
// Insert a new city
func CreateCity(w http.ResponseWriter, r *http.Request) {

	/*
	Example data for HTTP Post - "/cities"
	==============================================
	{
		"data":{
		"townName":"Narva",
		"currentPopulation":123456
		}
	}
	*/

	var dataResource CityResource
	// Decode the incoming City json
	dec := json.NewDecoder(r.Body) // reads data from request body to buffer
	err := dec.Decode(&dataResource) // decodes readed data using CityResourse model
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	city := &dataResource.Data
	log.Println("[Received data]:", city.TownName, city.CurrentPopulation)
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("cities")
	repo := &data.CityRepository{col}
	// Insert city document
	repo.Create(city) //Pointer-receiver function in data package
	if j, err := json.Marshal(CityResource{Data: *city}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		w.Header().Set("Content Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(j)
	}
}

// Handler for HTTP Get - "/cities"
// Return all City Documents
func GetCities(w http.ResponseWriter, r *http.Request) {
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("cities")
	repo := &data.CityRepository{C: col}
	cities := repo.GetAll()
	j, err := json.Marshal(CitiesResource{Data: cities})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}
