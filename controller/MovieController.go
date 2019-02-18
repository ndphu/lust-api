package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	"github.com/ndphu/drive-manager-api/dao"
	"github.com/ndphu/lust-api/entity"
	"io/ioutil"
	"net/http"
	"os"
)

func MovieController(r *gin.RouterGroup) {
	r.GET("/:id", func(c *gin.Context) {
		mid := c.Param("id")
		m := entity.Movie{}
		err := dao.Collection("movie").FindId(bson.ObjectIdHex(mid)).One(&m)
		if err != nil {
			ServerError("movie not found", err, c)
		}

		baseUrl := os.Getenv("STORAGE_BASE_URL")
		serviceToken := os.Getenv("STORAGE_SERVICE_TOKEN")

		reqUrl := baseUrl +
			"/driveAccount/" + m.DriveId.Hex() +
			"/file/" + m.FileId + "/download"
		req, err := http.NewRequest("GET", reqUrl, nil)
		if err != nil {
			ServerError("fail to get playing link", err, c)
			return
		}
		req.Header.Set("Authorization", "Bearer " + serviceToken)
		client := &http.Client{}
		resp, err := client.Do(req)

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
