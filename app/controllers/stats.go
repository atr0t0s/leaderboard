package controllers

import (
	"github.com/revel/revel"
	"labix.org/v2/mgo/bson"
	"leaderboard/app/models"
)

func (c App) CreateStat(name, metric string) revel.Result {

	if c.Session["user"] == "" || c.Session["role"] != adminrole {
		return c.RenderJson("Sorry, only the cluster admin can perform that action")
	} else {
		// connect to DB server(s)
		d, s := db(statcol)

		var doc models.Stat
		var results []models.Stat
		err := d.Find(bson.M{"statname": name}).Sort("-timestamp").All(&results)

		if err != nil {
			panic(err)
		} else {
			if len(results) == 0 {
				//do DB operations
				doc = models.Stat{Id: bson.NewObjectId(), StatName: name, StatMetric: metric}
				err = d.Insert(doc)
				if err != nil {
					panic(err)
				} else {
					s.Close()
					return c.RenderJson(doc)
				}
			} else {
				s.Close()
				return c.RenderJson("Error: A statistic with the same name already exists in the database.")
			}
		}

	}

}

// TODO: this func and the ones in achievement.go are for defining, retrieving names,
// and storing in the db when a user achieves something based on a stat
// However, we'll need to add a routing that watches over the database
// when a stat changes for a specific user, which will automatically call the
// Achieve func from achievement.go
func (c App) DefineAchievement(name string, statistic string, min float64) revel.Result {
	if c.Session["user"] == "" || c.Session["role"] != adminrole {
		return c.RenderJson("Sorry, only the cluster admin can perform that action")
	} else {
		// connect to DB server(s)
		d, s := db(achcol)

		var doc models.Ach
		var results []models.Ach
		err := d.Find(bson.M{"achname": name}).Sort("-timestamp").All(&results)

		if err != nil {
			panic(err)
		} else {
			if len(results) == 0 {
				//do DB operations
				doc = models.Ach{Id: bson.NewObjectId(), AchName: name, StatName: statistic, MinVal: min}
				err = d.Insert(doc)
				if err != nil {
					panic(err)
				} else {
					s.Close()
					return c.RenderJson(doc)
				}
			} else {
				s.Close()
				return c.RenderJson("Error: A statistic with the same name already exists in the database.")
			}
		}

	}

}

func (c App) LbSingleGame(GameUsersList string) {

}

func (c App) LbGlobal() {

}
