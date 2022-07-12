package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Student struct {
	Id       primitive.ObjectID `bson:"_id"`
	Name     string             `bson:"fullName"`
	Age      int                `bson:"age"`
	Gender   string             `bson:"gender"`
	JoinDate primitive.DateTime `bson:"joinDate"`
	Senior   bool               `bson:"senior"`
}
