package models

import (
	"gopkg.in/mgo.v2/bson"
)

type (
	User struct {
		Id bson.ObjectId `bson:"_id,omitempty" json:"id"`
		FirstName string `json:"firstname"`
		LastName string `json:"lastname"`
		Email string `json:"email"`
		Password string `json:"password,omitempty"`
		HashPassword []byte `json:"hashpassword,omitempty"`
		}
	City struct {
		Id                bson.ObjectId `bson:"_id,omitempty" json:"id"`
		TownName          string        `bson:"townName" json:"townName"`
		CurrentPopulation int        	`bson:"currentPopulation" json:"currentPopulation"`
	}
)
