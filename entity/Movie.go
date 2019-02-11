package entity

import "github.com/globalsign/mgo/bson"

type Movie struct {
	Id        bson.ObjectId `json:"_id" bson:"_id"`
	Category  string        `json:"category,omitempty" bson:"category"`
	Title     string        `json:"title"`
	SourceUrl string        `json:"sourceUrl,omitempty" bson:"sourceUrl"`
	Size      int64         `json:"size" bson:"size"`
	DriveId   bson.ObjectId `json:"driveId,omitempty" bson:"driveId"`
	FileId    string        `json:"fileId,omitempty" bson:"fileId"`
}
