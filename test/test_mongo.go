package main

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"context"
)

func main()  {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		fmt.Println(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Connected to MongoDB!")
	obj_id, err := primitive.ObjectIDFromHex("603878bcbea3e34b66f5a474")
	collection := client.Database("Lacs").Collection("role")

	filter := bson.D{{"_id", obj_id}}
	result,err := collection.DeleteOne(context.Background(),filter)
	fmt.Println(result)

	fmt.Println("%T",obj_id)

}