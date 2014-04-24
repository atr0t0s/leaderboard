package controllers

import (
	"fmt"
	"labix.org/v2/mgo"
	"os"
)

func db(collection string) *mgo.Collection {
	if uri == "" {
		fmt.Println("no connection string provided")
		os.Exit(1)
	}
	session, err := mgo.Dial(uri)
	if err != nil {
		panic(err)
	}

	defer session.Close()

	// select DB and Collection
	d := session.DB(dbname).C(collection)

	return d
}
