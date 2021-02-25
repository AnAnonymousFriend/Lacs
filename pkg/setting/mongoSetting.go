package setting

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
	"fmt"
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

func (m Mgo)FindOne(key string, value interface{}) interface{}{
	if m.Collection ==nil {
		fmt.Println("MongoDB Collection is nil")
		return nil
	}
	var result interface{}
	filter := bson.D{{"roleName", "admin"}}
	err := m.Collection.FindOne(context.TODO(), filter).Decode(&result)
	if err !=nil {
		fmt.Println(err)
		return nil
	}
	return result
}




