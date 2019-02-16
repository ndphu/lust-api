package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	"github.com/ndphu/drive-manager-api/dao"
	"github.com/ndphu/lust-api/entity"
	"io/ioutil"
	"net/http"
)

func MovieController(r *gin.RouterGroup) {
	r.GET("/:id", func(c *gin.Context) {
		mid := c.Param("id")
		m := entity.Movie{}
		err := dao.Collection("movie").FindId(bson.ObjectIdHex(mid)).One(&m)
		if err != nil {
			ServerError("movie not found", err, c)
		}

		resp, err := http.Get("https://drive-manager-api-villose-bassist.cfapps.io/api/manage" +
			"/driveAccount/" + m.DriveId.Hex() +
			"/file/" + m.FileId + "/download")
		if err != nil {
			ServerError("fail to get playing link", err, c)
			return
		}
		defer resp.Body.Close()

		bytes, _ := ioutil.ReadAll(resp.Body)
		data := make(map[string]string, 0)
		json.Unmarshal(bytes, &data)

		c.JSON(200, gin.H{
			"movie": m,
			"link":  data["link"],
		})
	})
}
