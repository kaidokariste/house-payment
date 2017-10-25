package controllers

import(
	"encoding/json"
	"net/http"

	"house-payment/data"

	"log"
)

// Handler for HTTP Get - "/cities"
// Return all City Documents
func GetCities(w http.ResponseWriter, r *http.Request){
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("cities")
	repo := &data.CityRepository{C:col}
	cities := repo.GetAll()
	log.Println("Cities from GetAll", cities)
	j, err := json.Marshal(CityResource{Data: cities})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)

}