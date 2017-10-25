package data

import (
	"gopkg.in/mgo.v2"
	"house-payment/models"
	_ "gopkg.in/mgo.v2/bson"
)


type CityRepository struct {
	C *mgo.Collection
}

func (r *CityRepository) GetAll() []models.City{
	// Define slice of City structs
	var cities []models.City
	iter := r.C.Find(nil).Iter()
	result := models.City{}
	for iter.Next(&result){
		cities = append(cities, result)
	}
	return cities
}