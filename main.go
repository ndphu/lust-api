package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	"github.com/ndphu/drive-manager-api/dao"
	entity2 "github.com/ndphu/drive-manager-api/entity"
	"github.com/ndphu/lust-api/controller"
	"github.com/ndphu/lust-api/entity"
	"log"
	"strings"
)

func main() {
	r := gin.Default()

	c := cors.DefaultConfig()
	c.AllowAllOrigins = true
	c.AllowCredentials = true
	c.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE"}
	c.AllowHeaders = []string{"Origin", "Authorization", "Content-Type", "Content-Length", "X-Requested-With"}

	r.Use(cors.New(c))

	api := r.Group("/api")
	controller.SearchController(api.Group("/search"))
	controller.CategoryController(api.Group("/category"))
	controller.MovieController(api.Group("/movie"))
	controller.HomeController(api.Group("/home"))
	controller.RecommendController(api.Group("/recommend"))
	controller.TrackingController(api.Group("/tracking"))

	//updateMovies()
	//updateMissingMovie()

	fmt.Println("Starting server")
	api.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "OK"})
	})
	r.Run()
}

func updateMissingMovie()  {
	driveFiles := make([]entity2.DriveFile, 0)
	dao.GetSession().DB("drive-manager").C("file_entry").Find(nil).All(&driveFiles)
	for _, df := range driveFiles {
		if df.Name == "Getting started" {
			continue
		}
		movieName := strings.TrimSuffix(df.Name, ".mp4")
		count, _ := dao.GetSession().DB("lust").C("movie").Find(bson.M{"title": movieName}).Count()
		if count == 0 {
			log.Println("create movie for", df.Name)
			movie := entity.Movie{
				Id: bson.NewObjectId(),
				Category: "uncensored",
				Title: movieName,
				Size: df.Size,
				DriveId: df.DriveAccount,
				FileId: df.DriveFileId,
				Tags: []string{"FullHD", "1080p", "HD"},
			}
			err := dao.GetSession().DB("lust").C("movie").Insert(&movie)
			if err != nil {
				log.Println("fail to create movie for file", df.Name)
			} else {
				log.Println("inserted movie", movie.Id.Hex(), movie.Title)
			}
		}
	}
}

func updateMovies() {
	movies := make([]entity.Movie, 0)
	movieCollection := dao.GetSession().DB("lust").C("movie")
	movieCollection.Find(nil).All(&movies)
	for _, m := range movies {
		query := strings.Replace(m.Title, "[", "\\[", -1)
		query = strings.Replace(query, "]", "\\]", -1)
		fe := entity2.DriveFile{}
		if err := dao.GetSession().DB("drive-manager").C("file_entry").
			Find(bson.M{"name": bson.RegEx{Pattern: query, Options: "i"}}).One(&fe); err != nil {
			log.Println("file not found for movie:", m.Title, err.Error())
			continue
		}
		m.Size = fe.Size
		m.FileId = fe.DriveFileId
		m.DriveId = fe.DriveAccount
		movieCollection.UpdateId(m.Id, &m)
	}
}
