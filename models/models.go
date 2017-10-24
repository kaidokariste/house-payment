package models

import (
	"gopkg.in/mgo.v2/bson"
)

type (
	City struct {
		Id                bson.ObjectId `bson:"_id,omitempty" json:"id"`
		TownName          string        `json:"townName"`
		CurrentPopulation string        `json:"currentPopulation"`
	}
)
