// All usage of mgo below is simply for kickstarting the app, but proper models
// should be implemented in a models folder

package controllers

import (
	"fmt"
	"github.com/revel/revel"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"os"
)

const (
	uri = "localhost" //change this to your mongodb server including auth (i.e. admin:pass@localhost)
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

// Create Users via HTTP POST call to /App/CreateUser
// You can manually add/remove fields by changing the params and 'doc' variable
// -----------
// Parameters:
// dbname -> the mongodb database name, collection -> the mongodb collection
// -----------
func (c App) CreateUser(dbname, collection, user, email, pass string) revel.Result {

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

	// TODO: Use encryption through crypto package to hash passwords

	// Query to see if user already exists in collection
	var doc User
	var results []User
	err = d.Find(bson.M{"username": user}).Sort("-timestamp").All(&results)

	if err != nil {
		panic(err)
	} else {
		if len(results) == 0 {
			//do DB operations
			doc = User{Id: bson.NewObjectId(), Username: user, Email: email, Password: pass}
			err = d.Insert(doc)
			if err != nil {
				panic(err)
			}
		} else {
			return c.RenderJson("Error: User already exists")
		}
	}

	return c.RenderJson(doc)

}

func (c App) Auth(dbname, collection, user, pass string) revel.Result {

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

	// TODO: Use encryption through crypto package to hash passwords

	// Query to authenticate

	var results []User

	err = d.Find(bson.M{"username": user, "password": pass}).Sort("-timestamp").All(&results)

	if err != nil {
		panic(err)
	} else {
		//fmt.Println(err)
		c.Session["user"] = user
		c.Flash.Success("Welcome, " + user)
	}

	return c.RenderJson(results)
}

func (c App) Logout() revel.Result {

	c.Session["user"] = ""

	session := c.Session["user"]

	return c.RenderJson(session)
}

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

func (c App) GetUserStats(username string) {

}

func (c App) DefineAchievement(achName, statName string, minVal int) {

}

func (c App) LbSingleGame(GameUsersList string) {

}

func (c App) LbGlobal() {

}

// Search available API functions and root to their documentation
// TODO: Should be done once an alpha version is ready
func (c App) RefSearch(findFunction string) revel.Result {

	/*
		c.Validation.Required(finfFunction).Message("Your name is required!")
		c.Validation.MinSize(findFunction, 3).Message("Your name is not long enough!")

		if c.Validation.HasErrors() {
			c.Validation.Keep()
			c.FlashParams()
			return c.Redirect(App.Index)
		}
	*/

	return c.Render(findFunction)
}
