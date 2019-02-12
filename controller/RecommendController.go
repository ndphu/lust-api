package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	"github.com/ndphu/drive-manager-api/dao"
	"github.com/ndphu/lust-api/entity"
)

func RecommendController(r *gin.RouterGroup) {
	r.GET("/movie/:id", func(id *gin.Context) {

	})

	r.GET("/category/:cat", func(c *gin.Context) {
		movies := make([]entity.Movie, 0)
		dao.Collection("movie").Pipe([]bson.M{
			{"$match": bson.M{"category": c.Param("cat")}},
			{"$sample": bson.M{"size": 20}},
			{"$project": bson.M{"_id": 1, "title": 1, "size": 1, "category": 1}},
		}).All(&movies)
		c.JSON(200, gin.H{
			"movies": movies,
		})
	})
}
