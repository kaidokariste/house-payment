package models

import (
	"gopkg.in/mgo.v2/bson"
)

type (
	City struct {
		Id                bson.ObjectId `bson:"_id,omitempty" json:"id"`
		TownName          string        `bson:"townName" json:"townName"`
		CurrentPopulation string        `bson:"currentPopulation" json:"currentPopulation"`
	}
)
