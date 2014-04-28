package controllers

import (
	"fmt"
	"github.com/revel/revel"
	"labix.org/v2/mgo/bson"
	"leaderboard/app/models"
)

func (c App) GetStat(statName string) *models.Stat {

	// connect to DB server(s)
	d, s := db(statcol)

	// Query
	results := models.Stat{}
	query := bson.M{"statname": statName}
	err := d.Find(query).One(&results)

	if err != nil {
		panic(err)
	}
	if len(results.StatName) == 0 {
		return nil
	}

	s.Close()

	return &results

}

func (c App) SaveUserStat(statName string, statValue float64) revel.Result {

	if c.Session["user"] == "" || c.Session["role"] != playerrole {
		return c.RenderJson("User is not logged in, or user is not a player")
	} else {
		username := c.Session["user"]
		stat := c.GetStat(statName)
		user := c.GetUser(username)
		// connect to DB server
		d, s := db(userstatcol)

		// Query
		var doc models.UserStat
		results := models.UserStat{}
		query := bson.M{"statistic": stat.StatName, "user": user.Username}
		err := d.Find(query).One(&results)

		if len(results.StatName) == 0 {
			//do DB operations
			doc = models.UserStat{Id: bson.NewObjectId(), StatName: statName, Value: statValue, Username: user.Username}
			err = d.Insert(doc)
			if err != nil {
				panic(err)
			} else {
				s.Close()
				return c.RenderJson(doc)
			}
		} else {

			newValue := statValue + results.Value
			fmt.Println(newValue) // debug
			colQuerier := bson.M{"statistic": stat.StatName, "user": user.Username}
			change := bson.M{"$set": bson.M{"value": newValue}}
			err = d.Update(colQuerier, change)
			if err != nil {
				panic(err)
			} else {
				s.Close()
				return c.RenderJson(err)
			}
		}

		if err != nil {
			panic(err)
		}

		return c.RenderJson(results)
	}

}

func (c App) GetUserStats(username string) {

}
