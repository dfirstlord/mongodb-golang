package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const mongodbUri = "mongodb://127.0.0.1:27017"

type Student struct {
	Id       primitive.ObjectID `bson:"_id"`
	Name     string             `bson:"fullName"`
	Age      int                `bson:"age"`
	Gender   string             `bson:"gender"`
	JoinDate primitive.DateTime `bson:"joinDate"`
	Senior   bool               `bson:"senior"`
}

func main() {
	credentials := options.Credential{
		AuthMechanism: "SCRAM-SHA-256",
		Username:      "ianlord",
		Password:      "password",
	}

	clientOptions := options.Client()
	clientOptions.ApplyURI(mongodbUri).SetAuth(credentials)
	ctx := context.Background()

	connect, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("connected...")
	}
	defer func() {
		if err := connect.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	//membuat sebuah db - collection
	db := connect.Database("enigma")
	coll := db.Collection("student")

	//create
	//insert into one
	// newId, err := coll.InsertOne(ctx, bson.D{
	// 	{"name", "Mama"},
	// 	{"age", 21},
	// 	{"gender", "F"},
	// 	{"senior", false},
	// })
	// if err != nil {
	// 	log.Println(err.Error())
	// }
	// fmt.Printf("inserted document with ID %v\n", newId.InsertedID)

	// jd01 := parseTime("2022-07-02 15:04:05")
	// jd02 := parseTime("2022-07-03 15:04:05")

	// //insert into many
	// students := []interface{}{
	// 	bson.D{
	// 		{"name", "jack"},
	// 		{"age", 30},
	// 		{"senior", true},
	// 		{"gender", "M"},
	// 		{"joinDate", primitive.NewDateTimeFromTime(jd01)}},
	// 	bson.D{
	// 		{"name", "dimitri"},
	// 		{"age", 20},
	// 		{"senior", false},
	// 		{"gender", "F"},
	// 		{"joinDate", jd02}},
	// }
	// result, err := coll.InsertMany(ctx, students)
	// if err != nil {
	// 	log.Println(err.Error())
	// }
	// fmt.Printf("inserted document with ID %v\n", result.InsertedIDs)

	// newStudent := Student{
	// 	Id:       primitive.NewObjectID(),
	// 	Name:     "Dino",
	// 	Age:      25,
	// 	Gender:   "M",
	// 	JoinDate: primitive.NewDateTimeFromTime(parseTime("2022-05-22 14:00:45")),
	// 	Senior:   false,
	// }

	// result, err := coll.InsertOne(ctx, newStudent)
	// if err != nil {
	// 	log.Println(err.Error())
	// }
	// fmt.Printf("inserted document with ID %v\n", result.InsertedID)

	//delete
	// filter := bson.D{{"fullName", "Dona"}}
	// result, err := coll.DeleteOne(ctx, filter)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("Number of documents deleted: %d\n", result.DeletedCount)

	// //update
	// filter := bson.D{{Key: "fullName", Value: "Doni"}}
	// update := bson.D{{Key: "$set", Value: bson.D{{Key: "fullName", Value: "Ridwan"}}}}
	// result, err := coll.UpdateOne(ctx, filter, update)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("Documents updated: %v\n", result.ModifiedCount)

	
}

func parseTime(date string) time.Time {
	layoutFormat := "2006-01-02 15:04:05"
	parse, _ := time.Parse(layoutFormat, date)
	return parse
}

/*
Buat koneksi ke mongodb
*/
