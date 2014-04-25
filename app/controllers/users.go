package controllers

import (
	"github.com/revel/revel"
	"labix.org/v2/mgo/bson"
	"leaderboard/app/models"
)

// Create Users via HTTP POST call to /App/CreateUser
// You can manually add/remove fields by changing the params and 'doc' variable
func (c App) CreateUser(name, user, email, pass, role string) revel.Result {

	if c.Session["user"] == "" || c.Session["role"] != adminrole {
		return c.RenderJson("Sorry, only the cluster admin can perform that action")
	} else {
		// connect to DB server
		d, s := db(usercol)

		// Query to see if user already exists in collection
		var doc models.User
		var results []models.User

		err := d.Find(bson.M{"username": user}).Sort("-timestamp").All(&results)

		if err != nil {
			panic(err)
		} else {
			if len(results) == 0 {
				//do DB operations
				doc = models.User{Id: bson.NewObjectId(), Name: name, Username: user, Email: email, Password: pass}
				err = d.Insert(doc)
				if err != nil {
					panic(err)
				}
			} else {
				return c.RenderJson("Error: User already exists")
			}
		}

		s.Close()

		return c.RenderJson(doc)
	}
}

func (c App) Auth(user, pass string) revel.Result {

	// connect to DB server(s)
	d, s := db(usercol)

	// TODO: Use encryption through crypto package to hash passwords
	// Query to authenticate

	results := models.User{}
	query := bson.M{"username": user, "password": pass}
	err := d.Find(query).One(&results)

	if err != nil {
		panic(err)
	} else {
		//fmt.Println(err)
		c.Session["user"] = results.Username
		c.Session["role"] = results.Role
		c.Flash.Success("Welcome, " + results.Username)
	}

	s.Close()

	return c.RenderJson(results)
}

func (c App) Logout() revel.Result {

	c.Session["user"] = ""

	session := c.Session["user"]

	return c.RenderJson(session)
}

func (c App) SaveUserStat(statName, username string) {

}

func (c App) GetUserStats(username string) {

}
