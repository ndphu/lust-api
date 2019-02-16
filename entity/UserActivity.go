package entity

import (
	"github.com/globalsign/mgo/bson"
	"time"
)

type UserActivity struct {
	Id bson.ObjectId `json:"id" bson:"_id"`
	UserId string `json:"userId"  bson:"userId"`
	Action string `json:"action" bson:"action"`
	Timestamp time.Time `json:"timestamp" bson:"timestamp"`
	Details interface{} `json:"details" bson:"details"`
}
