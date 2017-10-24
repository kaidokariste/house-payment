package common

import (
	"gopkg.in/mgo.v2"
	"time"
	"log"
)

var session *mgo.Session

// This function is used in CRUD calls and session is obtained trough Copy command
func GetSession() *mgo.Session {
	if session == nil {
		var err error
		session, err = mgo.DialWithInfo(&mgo.DialInfo{
			Addrs:    []string{AppConfig.MongoDBHost}, //Global configuration struct from utils.go
			Username: AppConfig.DBUser,
			Password: AppConfig.DBPwd,
			Timeout:  60 * time.Second,
		})
		if err != nil {
			log.Fatalf("[GetSession]: %s\n", err)
		}

	}
	return session
}

// This will go to booter and is called in main.go before HTTP server start.
func createDBSession() {
	var err error
	session, err = mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    []string{AppConfig.MongoDBHost},
		Username: AppConfig.DBUser,
		Password: AppConfig.DBPwd,
		Timeout:  60 * time.Second,
	})
	if err != nil {
		log.Fatalf("[createDBSession]: %s\n", err)
	}
}
