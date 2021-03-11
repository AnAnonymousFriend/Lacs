package setting

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var database  *mongo.Database

type Mgo struct {
	*mongo.Collection
}

func MongoDBSetup() *mongo.Database  {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(MongoDBSetting.Host).SetMaxPoolSize(MongoDBSetting.MaxConn))
	if err !=nil {
		panic(err)
	}
	database =  client.Database(MongoDBSetting.DbName)
	return  database
}

func NewMongoClient(tableName string) *mongo.Collection {
	con := database.Collection(tableName)
	return con
}


func DisConn(client *mongo.Client)  {
	client.Disconnect(context.Background())
}





