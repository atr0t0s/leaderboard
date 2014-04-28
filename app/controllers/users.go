package controllers

import (
	"code.google.com/p/go.crypto/bcrypt"
	//"encoding/base64"
	//"fmt"
	"github.com/revel/revel"
	"labix.org/v2/mgo/bson"
	"leaderboard/app/models"
)

// Create Users via HTTP POST call to /App/CreateUser
// You can manually add/remove fields by changing the params and 'doc' variable
func (c App) CreateUser(name string, username string, email string, password string, role string) revel.Result {

	if c.Session["user"] == "" || c.Session["role"] != adminrole {
		return c.RenderJson("Sorry, only the cluster admin can perform that action")
	} else {
		// connect to DB server
		d, s := db(usercol)

		// Query to see if user already exists in collection
		var doc models.User
		var results []models.User
		doc.HashPass, _ = bcrypt.GenerateFromPassword(
			[]byte(password), bcrypt.DefaultCost)

		err := d.Find(bson.M{"username": username}).Sort("-timestamp").All(&results)

		if err != nil {
			panic(err)
		} else {
			if len(results) == 0 {
				//do DB operations
				doc = models.User{Id: bson.NewObjectId(), Name: name, Username: username, Email: email, HashPass: doc.HashPass, Role: role}
				err = d.Insert(doc)
				if err != nil {
					panic(err)
				} else {
					s.Close()
					return c.RenderJson(doc)
				}
			} else {
				s.Close()
				return c.RenderJson("Error: User already exists")
			}
		}

	}
}

func (c App) getUser(username string) *models.User {

	// connect to DB server(s)
	d, s := db(usercol)

	// Query
	results := models.User{}
	query := bson.M{"username": username}
	err := d.Find(query).One(&results)

	if err != nil {
		panic(err)
	}
	if len(results.Username) == 0 {
		return nil
	}

	s.Close()

	return &results

}

func (c App) Auth(username string, password string, remember bool) revel.Result {

	user := c.getUser(username)
	if user != nil {
		err := bcrypt.CompareHashAndPassword(user.HashPass, []byte(password))
		if err == nil {
			c.Session["user"] = username
			if remember {
				c.Session.SetDefaultExpiration()
			} else {
				c.Session.SetNoExpiration()
			}
			c.Flash.Success("Welcome, " + username)
			return c.Redirect("/")
		}
	}

	c.Flash.Out["username"] = username
	c.Flash.Error("Login failed")

	return c.Redirect("/")
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
