package entity

import "github.com/globalsign/mgo/bson"

type Movie struct {
	Id        bson.ObjectId `json:"_id" bson:"_id"`
	Category  string        `json:"category" bson:"category"`
	Title     string        `json:"title"`
	SourceUrl string        `json:"sourceUrl" bson:"sourceUrl"`
	Size      int64         `json:"size" bson:"size"`
	DriveId   bson.ObjectId `json:"driveId" bson:"driveId"`
	FileId    string        `json:"fileId" bson:"fileId"`
}
