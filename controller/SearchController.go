package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	"github.com/ndphu/drive-manager-api/dao"
	"github.com/ndphu/lust-api/entity"
)

func SearchController(r *gin.RouterGroup) error {
	r.GET("quickSearch", func(c *gin.Context) {
		movies := make([]entity.Movie, 0)
		query := c.Query("query")
		dao.Collection("movie").
			Find(bson.M{
				"title": bson.RegEx{Pattern: query, Options: "i"},
			}).
			Select(bson.M{
				"_id":      1,
				"title":    1,
				"category": 1,
			}).
			Limit(20).
			All(&movies)
		c.JSON(200, gin.H{"movies": movies})
	})
	return nil
}
