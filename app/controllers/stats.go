package controllers

import (
	"fmt"
	"github.com/revel/revel"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"os"
)

func (c App) CreateStat(dbname, collection, statName, statMetric string) revel.Result {

	// connect to DB server(s)

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

	var doc Stat
	var results []Stat
	err = d.Find(bson.M{"statname": statName}).Sort("-timestamp").All(&results)

	if err != nil {
		panic(err)
	} else {
		if len(results) == 0 {
			//do DB operations
			doc = Stat{Id: bson.NewObjectId(), StatName: statName, StatMetric: statMetric}
			err = d.Insert(doc)
			if err != nil {
				panic(err)
			}
		} else {
			return c.RenderJson("Error: A statistic with the same name already exists in the database.")
		}
	}

	return c.RenderJson(doc)
}

func (c App) DefineAchievement(achName, statName string, minVal int) {

}

func (c App) LbSingleGame(GameUsersList string) {

}

func (c App) LbGlobal() {

}
