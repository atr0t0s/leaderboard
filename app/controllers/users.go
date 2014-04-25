package controllers

import (
	"code.google.com/p/go.crypto/bcrypt"
	//"encoding/base64"
	"fmt"
	"github.com/revel/revel"
	"labix.org/v2/mgo/bson"
	"leaderboard/app/models"
)

// Create Users via HTTP POST call to /App/CreateUser
// You can manually add/remove fields by changing the params and 'doc' variable
func (c App) CreateUser(name string, user string, email string, pass []byte, role string) revel.Result {

	/*if c.Session["user"] == "" || c.Session["role"] != adminrole {
		return c.RenderJson("Sorry, only the cluster admin can perform that action")
	} else {*/
	// connect to DB server
	d, s := db(usercol)

	// Query to see if user already exists in collection
	var doc models.User
	var results []models.User
	doc.HashPass, _ = bcrypt.GenerateFromPassword(
		pass, bcrypt.DefaultCost)

	err := d.Find(bson.M{"username": user}).Sort("-timestamp").All(&results)
	fmt.Println(pass)
	if err != nil {
		panic(err)
	} else {
		if len(results) == 0 {
			//do DB operations
			doc = models.User{Id: bson.NewObjectId(), Name: name, Username: user, Email: email, HashPass: doc.HashPass, Role: role}
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
	//}
}

func (c App) Auth(user string, pass []byte) revel.Result {

	// connect to DB server(s)
	d, s := db(usercol)

	// Query to authenticate

	results := models.User{}
	query := bson.M{"username": user}
	err := d.Find(query).One(&results)

	// TODO: fix this, there's something seriously wrong with getting the
	// pass []byte from the argument, it comes out as null, i.e. []
	verify := bcrypt.CompareHashAndPassword(results.HashPass, pass)

	if err != nil {
		panic(err)
	} else {
		if verify != nil {
			panic(verify)
		} else {
			c.Session["user"] = results.Username
			c.Session["role"] = results.Role
			c.Flash.Success("Welcome, " + results.Username)
		}
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
