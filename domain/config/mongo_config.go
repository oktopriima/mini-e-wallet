/*
 * Name : Okto Prima Jaya
 * Github : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 18/11/2019, 20:50
 * Copyright (c) 2019
 */

package config

import (
	"log"

	"gopkg.in/mgo.v2"
)

func NewMongoDBConfig(cfg Config) (error, *mgo.Database) {
	address := cfg.GetString(`mongodb.address`)
	database := cfg.GetString(`mongodb.database`)
	log.Println("mongo config", address, database)
	session, err := mgo.Dial(address)
	session.SetMode(mgo.Monotonic, true)
	return err, session.DB(database)
}
