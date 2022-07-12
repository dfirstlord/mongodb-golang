package main

import (
	"context"
	"fmt"
	"log"
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

	//read
	//select * from student
	cursor, err := coll.Find(ctx, bson.D{})
	if err != nil {
		log.Println(err.Error())
	}
	var students []bson.D
	err = cursor.All(ctx, &students)
	if err != nil {
		log.Println(err.Error())
	}
	for _, student := range students {
		fmt.Println("Select *", student)
	}

	//projection
	// opts := options.Find().SetProjection(bson.D{{Key: "age", Value: 1}, {Key: "gender", Value: 1}, {Key: "_id", Value: 1}})
	// cursor, err := coll.Find(ctx, bson.D{}, opts)
	// if err != nil {
	// 	log.Println(err.Error())
	// }
	// var results []bson.D
	// if err = cursor.All(ctx, &results); err != nil {
	// 	log.Println(err.Error())
	// }
	// for _, result := range results {
	// 	fmt.Println(result)
	// }

	//Logical
	filterGenderAndAge := bson.D{
		{
			Key: "$and", Value: bson.A{
				bson.D{
					{"gender", "F"},
					{Key: "age", Value: bson.D{{Key: "$gt", Value: 15}}},
				},
			},
		},
	}

	projection := bson.D{
		{Key: "_id", Value: 1},
		{Key: "gender", Value: 1},
		{Key: "age", Value: 1},
	}

	// studentx, err := coll.Find(ctx, filterGenderAndAge, options.Find().SetProjection(projection))
	// if err != nil {
	// 	log.Println(err)
	// }
	// err = studentx.All(ctx, &students)
	// if err != nil {
	// 	log.Println(err.Error())
	// }
	//mapping result query ke struct
	filterGenderAndAgeResult := make([]*Student, 0)
	cursor, err = coll.Find(ctx, filterGenderAndAge, options.Find().SetProjection(projection))
	if err != nil {
		log.Println(err)
	}
	for cursor.Next(ctx) {
		var student Student
		err := cursor.Decode(&student)
		if err != nil {
			log.Println(err.Error())
		}
		filterGenderAndAgeResult = append(filterGenderAndAgeResult, &student)
	}
	for _, student := range filterGenderAndAgeResult {
		fmt.Println("Filter by gender and age", student)
	}

	//aggregation
	colll := connect.Database("enigma").Collection("products")
	count, err := colll.CountDocuments(ctx, bson.D{{"category", "food"}})
	if err != nil {
		log.Println(err.Error())
	}
	fmt.Println("Product total in food category", count)

	//match
	matchStage := bson.D{
		{
			Key: "$match", Value: bson.D{
				{Key: "category", Value: "food"},
			},
		},
	}

	groupStage := bson.D{
		{
			Key: "$group", Value: bson.D{
				{Key: "_id", Value: "$category"},
				{Key: "Total", Value: bson.D{{"$sum", 1}}},
			},
		},
	}

	cursor, err = colll.Aggregate(ctx, mongo.Pipeline{matchStage, groupStage})
	if err != nil {
		log.Println(err.Error())
	}
	var productCount []bson.M
	err = cursor.All(ctx, &productCount)
	if err != nil {
		log.Println(err.Error())
	}
	for _, product := range productCount {
		fmt.Printf("Group[%v], Total [%v]\n", product["_id"], product["Total"])
	}
}

func parseTime(date string) time.Time {
	layoutFormat := "2006-01-02 15:04:05"
	parse, _ := time.Parse(layoutFormat, date)
	return parse
}

/*
Buat koneksi ke mongodb
*/
