package data

import (
	"gopkg.in/mgo.v2"
	"house-payment/models"
	 _ "gopkg.in/mgo.v2/bson"
	"log"

)


type CityRepository struct {
	C *mgo.Collection
}

func (r *CityRepository) GetAll() []models.City{
	// Define slice of City structs
	var cities []models.City
	iter := r.C.Find(nil).Iter()
	log.Println("GetAll/ iter reult ",iter)
	result := models.City{}
	for iter.Next(&result){
		log.Println("GetAll/ iteration ",result)
		cities = append(cities, result)
	}
	return cities
}