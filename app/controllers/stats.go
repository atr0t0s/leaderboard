package controllers

import (
	"github.com/revel/revel"
	"labix.org/v2/mgo/bson"
	"leaderboard/app/models"
)

func (c App) CreateStat(statName, statMetric string) revel.Result {

	// connect to DB server(s)
	d, s := db(statcol)

	var doc models.Stat
	var results []models.Stat
	err := d.Find(bson.M{"statname": statName}).Sort("-timestamp").All(&results)

	if err != nil {
		panic(err)
	} else {
		if len(results) == 0 {
			//do DB operations
			doc = models.Stat{Id: bson.NewObjectId(), StatName: statName, StatMetric: statMetric}
			err = d.Insert(doc)
			if err != nil {
				panic(err)
			}
		} else {
			return c.RenderJson("Error: A statistic with the same name already exists in the database.")
		}
	}

	s.Close()

	return c.RenderJson(doc)
}

func (c App) DefineAchievement(achName string, stat models.Stat, minVal int) {

}

func (c App) LbSingleGame(GameUsersList string) {

}

func (c App) LbGlobal() {

}
