package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	"github.com/ndphu/drive-manager-api/dao"
	"github.com/ndphu/lust-api/entity"
	"sync"
)

func HomeController(r*gin.RouterGroup)  {
	r.GET("", func(c *gin.Context) {
		wg := sync.WaitGroup{}
		result := make(map[string]interface{})
		for _, cat := range []string{"censored","uncensored", "uniform", "beauty"} {
			wg.Add(1)
			go func(category string) {
				movies := make([]entity.Movie, 0)
				dao.Collection("movie").Pipe([]bson.M{
					{"$match": bson.M{"category": category, "size": bson.M{"$gt": 0},}},
					{"$sample": bson.M{"size": 20}},
					{"$project": bson.M{"_id": 1, "title": 1, "size": 1, "category": 1}},
				}).All(&movies)
				result[category] = gin.H{
					"category": category,
					"movies": movies,
				}
				wg.Done()
			}(cat)
		}
		wg.Wait()

		c.JSON(200, result)

	})
}
