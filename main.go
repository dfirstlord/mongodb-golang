package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const mongodbUri = "mongodb://127.0.0.1:27017"

func main() {
	credentials := options.Credential{
		AuthMechanism: "SCRAM-SHA-256",
		Username: "ianlord",
		Password: "password",
	}

	clientOptions := options.Client()
	clientOptions.ApplyURI(mongodbUri).SetAuth(credentials)
	ctx := context.Background()

	connect, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		panic(err)
	}else{
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

	
}

/*
Buat koneksi ke mongodb
*/
