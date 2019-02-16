package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	"github.com/ndphu/drive-manager-api/dao"
	"github.com/ndphu/drive-manager-api/utils"
	"github.com/ndphu/lust-api/entity"
	"time"
)

type WatchHistory struct {
	Id bson.ObjectId `json:"id" bson:"_id"`
	Movie entity.Movie `json:"movie" bson:"movie"`
	Timestamp time.Time `json:"timestamp" bson:"timestamp"`
}


func TrackingController(r *gin.RouterGroup) {
	r.POST("/watch/movie", func(c *gin.Context) {
		movieId := c.Query("movieId")
		userId := c.Query("userId")
		dao.Collection("user_activity").Insert(&entity.UserActivity{
			Id:        bson.NewObjectId(),
			Timestamp: time.Now(),
			UserId:    userId,
			Action: "WATCH",
			Details: bson.M{
				"movieId": bson.ObjectIdHex(movieId),
			},
		})
		c.JSON(200, gin.H{})
	})

	r.POST("/search/quick", func(c *gin.Context) {
		query := c.Query("query")
		userId := c.Query("userId")
		dao.Collection("user_activity").Insert(&entity.UserActivity{
			Id:        bson.NewObjectId(),
			Timestamp: time.Now(),
			UserId:    userId,
			Action: "SEARCH",
			Details: bson.M{
				"quick": true,
				"query": query,
			},
		})
		c.JSON(200, gin.H{})
	})

	r.GET("/watchHistory", func(c *gin.Context) {
		hist := make([]WatchHistory, 0)
		err := dao.Collection("user_activity").Pipe([]bson.M{
			{"$match": bson.M{"action": "WATCH"}},
			{"$sort": bson.M{"timestamp": -1}},
			{"$lookup": bson.M{
				"from": "movie",
				"localField": "details.movieId",
				"foreignField": "_id",
				"as": "movie",
			}},
			{"$unwind": bson.M{"path": "$movie"}},
			{"$project": bson.M{
				"movie.sourceUrl": 0,
				"movie.driveId": 0,
				"movie.fileId": 0,
			},},
			{"$limit": utils.GetIntQuery(c, "limit", 50)},
		}).All(&hist)

		if err != nil {
			ServerError("fail to get watch history", err, c)
			c.Abort()
		}

		c.JSON(200, hist)
	})

	r.GET("/searchHistory", func(c *gin.Context) {
		hist := make([]entity.UserActivity, 0)
		err := dao.Collection("user_activity").Find(bson.M{
			"action":"SEARCH",
		}).All(&hist)

		if err != nil {
			ServerError("fail to get search history", err, c)
			c.Abort()
		}

		c.JSON(200, hist)
	})
}
