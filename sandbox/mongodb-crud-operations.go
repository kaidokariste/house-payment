package main

import(
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"fmt"
	"house-payment/common"
	"time"
)

type (
	City struct {
		Id                bson.ObjectId `bson:"_id,omitempty" json:"id"`
		TownName          string        `bson:"townName" json:"townName"`
		CurrentPopulation int           `bson:"currentPopulation" json:"currentPopulation"`
	}
)

func main()  {

	//Load in the configuration data
	common.StartUp()

	//TODO If interestes, then investigate why DialWithInfo brokes the authorisation
	//session, err = mgo.DialWithInfo(&mgo.DialInfo{
	//	Addrs: []string{AppConfig.MongoDBHost},
	//	Username: AppConfig.DBUser,
	//	Password: AppConfig.DBPwd,
	//	Timeout: 60 * time.Second,
	//})

	// Session trough URI
	session, err := mgo.Dial(common.AppConfig.MongoDBFullURI)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	log.Println("[session object]:", session)

	session.SetMode(mgo.Monotonic, true)
	//get collection
	c := session.DB("house-payment").C("cities")
	log.Println("[get collection]:", c)

	// Find all entries
	iter := c.Find(nil).Iter()
	log.Println("[iter object]:", iter)
	result := City{}
	for iter.Next(&result) {
		fmt.Printf("City: %s, Population: %s\n", result.TownName, result.CurrentPopulation)
	}
	if err = iter.Close(); err != nil {
		log.Fatal(err)
	}

}