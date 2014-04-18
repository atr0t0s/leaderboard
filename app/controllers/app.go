package controllers

import (
	"fmt"
	"github.com/revel/revel"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"os"
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

// Goes to web view -> TODO: display API reference
func (c App) Index() revel.Result {

	greeting := "Will use this page to post the API reference"
	return c.Render(greeting)

}

// Create Users via -> /App/CreateUser?dbase=<Database>&collect=<Collection>&user=<Username>&email=<Email>&pass=<Password>
// You can manually add/remove fields by changing the params and 'doc' variable
// -----------
// Parameters:
// uri -> database host with auth info. i.e. for mongodb you could pass "admin:pass@localhost"
// dbname -> the mongodb database name, collection -> the mongodb collection
// -----------
func (c App) CreateUser(uri, dbname, collection, user, email, pass string) revel.Result {

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

// DB params (uri, dbname, collection) explained in func CreateUser
func (c App) Auth(uri, dbname, collection, user, pass string) revel.Result {

	// connect to DB server(s)

	uri := dbdetails[0]
	dbname := dbdetails[1]
	collection := dbdetails[2]

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

	// TODO: Fix the below to search for matches of username+password
	err = d.Find(bson.M{"username": user}).Sort("-timestamp").All(&results)

	if err != nil {
		panic(err)
	}

	return c.Render(doc)
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
