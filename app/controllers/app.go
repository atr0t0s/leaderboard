// All usage of mgo below is simply for kickstarting the app, but proper models
// should be implemented in a models folder

package controllers

import (
	"github.com/revel/revel"
	"labix.org/v2/mgo/bson"
)

const (
	uri = "localhost"
	//change this to your mongodb server including auth (i.e. admin:pass@localhost)
	// TODO: change this to retrieve uri from a config file
)

// User Struct
type User struct {
	Id       bson.ObjectId `bson:"_id"`
	Username string        `bson:"username"`
	Email    string        `bson:"email"`
	Password string        `bson:"password"`
}

// App Struct
type App struct {
	*revel.Controller
}

type Stat struct {
	Id         bson.ObjectId `bson:"_id"`
	StatName   string        `bson:"statname"`
	StatMetric string        `bson:"statmetric"`
}

// Goes to web view -> TODO: display API reference
func (c App) Index() revel.Result {

	var greeting string

	if c.Session["user"] != "" {
		greeting = "Welcome " + c.Session["user"]
	} else {
		greeting = "Please login"
	}

	return c.Render(greeting)

}

// Search available API functions and root to their documentation
// TODO: Should be done once an alpha version is ready
func (c App) RefSearch(findFunction string) revel.Result {

	return c.Render(findFunction)
}
