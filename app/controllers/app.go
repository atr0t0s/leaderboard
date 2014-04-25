// All usage of mgo below is simply for kickstarting the app, but proper models
// should be implemented in a models folder
// TODO *: Use encryption through crypto package to hash passwords
// <- BRANCH CRYPTO
// START WORK ON ABOVE TODO * IN THIS BRANCH
// <- REMOVE COMMENT BEFORE MERGING
package controllers

import (
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
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
