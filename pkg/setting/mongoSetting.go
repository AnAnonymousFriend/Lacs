package setting

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
	"fmt"
)

func MongoDBSetup() *mongo.Database  {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	fmt.Printf("%s\n", MongoDBSetting.Host)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(MongoDBSetting.Host))
	if err !=nil {
		panic(err)
	}
	database :=  client.Database(MongoDBSetting.DbName)
	return  database
}

func DisConn(client *mongo.Client)  {
	client.Disconnect(context.Background())
}


func NewDataTableCollent(dataTable string) *mongo.Collection  {
	return MongoDataBase.Collection(dataTable)
}




