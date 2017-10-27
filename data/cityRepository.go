package data

import (
	"gopkg.in/mgo.v2"
	"house-payment/models"

	"log"
	"gopkg.in/mgo.v2/bson"
)

var err error

type CityRepository struct {
	C *mgo.Collection
}

// Create new city entri to collection
func (r *CityRepository) Create(city *models.City) error {
	obj_id := bson.NewObjectId()
	city.Id = obj_id
	err := r.C.Insert(&city)
	return err
}

// Function to get all cities from collection
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