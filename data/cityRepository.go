package data

import (
	"gopkg.in/mgo.v2"
	"house-payment/models"

	"log"
)

var err error

type CityRepository struct {
	C *mgo.Collection
}

func (r *CityRepository) GetAll() []models.City{
	// Define slice of City structs
	var cities []models.City
	iter := r.C.Find(nil).Iter()
	result := models.City{}
	for iter.Next(&result){
		//fmt.Printf("City: %s, Population: %s\n", result.TownName, result.CurrentPopulation)
		cities = append(cities, result)
	}
	if err := iter.Close(); err != nil {
		log.Println("[error on iteration]", err)
	}
	return cities
}