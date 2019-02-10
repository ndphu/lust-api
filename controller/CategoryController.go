package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	"github.com/ndphu/drive-manager-api/dao"
	"github.com/ndphu/drive-manager-api/utils"
	"github.com/ndphu/lust-api/entity"
)

func CategoryController(r *gin.RouterGroup) {
	r.GET("/:id", func(c *gin.Context) {
		movies := make([]entity.Movie, 0)
		page := utils.GetIntQuery(c, "page", 1)
		size := utils.GetIntQuery(c, "size", 10)
		dao.GetSession().DB("lust").C("movie").Find(bson.M{
			"category": c.Param("id"),
		}).Skip((page - 1) * size).Limit(size + 1).All(&movies)
		hasMore := len(movies) > size
		if len(movies) > 0 {
			c.JSON(200, gin.H{
				"movies": movies[:len(movies) - 1],
				"hasMore": hasMore,
				"page": page,
				"size": size,
			})
		} else {
			c.JSON(200, gin.H{
				"movies": movies,
				"hasMore": false,
				"page": page,
				"size": size,
			})
		}
	})
}
