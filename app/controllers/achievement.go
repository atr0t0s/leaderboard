package controllers

import (
	//"fmt"
	"github.com/revel/revel"
	"labix.org/v2/mgo/bson"
	"leaderboard/app/models"
)

func (c App) GetAch(name string) *models.Ach {

	// connect to DB server(s)
	d, s := db(achcol)

	// Query
	results := models.Ach{}
	query := bson.M{"achname": name}
	err := d.Find(query).One(&results)

	if err != nil {
		panic(err)
	}
	if len(results.AchName) == 0 {
		return nil
	}

	s.Close()

	return &results

}

func (c App) Achieve(achName string, complete bool) revel.Result {

	if c.Session["user"] == "" || c.Session["role"] != playerrole {
		return c.RenderJson("User is not logged in, or user is not a player")
	} else {
		username := c.Session["user"]
		ach := c.GetAch(achName)
		user := c.GetUser(username)
		// connect to DB server
		d, s := db(userachcol)

		// Query
		var doc models.UserAch
		results := models.UserAch{}
		query := bson.M{"achievement": ach.AchName, "user": user.Username}
		err := d.Find(query).One(&results)

		if err != nil {
			panic(err)
		} else {
			if len(results.AchName) == 0 {
				//do DB operations
				doc = models.UserAch{Id: bson.NewObjectId(), AchName: achName, Complete: complete, Username: user.Username}
				err = d.Insert(doc)
				if err != nil {
					panic(err)
				} else {
					s.Close()
					return c.RenderJson(doc)
				}
			}
		}

		return c.RenderJson(results)
	}

}

func (c App) GetUserAchieves(username string) revel.Result {

	// connect to DB server(s)
	d, s := db(userachcol)

	var user string

	if len(username) == 0 {
		user = c.Session["user"]
	} else {
		user = username
	}

	// Query
	var results []models.UserAch
	query := bson.M{"user": user}
	err := d.Find(query).Sort("-timestamp").All(&results)

	if err != nil {
		panic(err)
	}
	if len(results) == 0 {
		return nil
	}

	s.Close()

	return c.RenderJson(results)

}
