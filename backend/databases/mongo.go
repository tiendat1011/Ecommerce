package databases

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"

	"ecommerce-project/config"
)

var DB *mongo.Database

func InitMongoDB(uri string) {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(opts)
	if err != nil {
		panic(err)
	}
	
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		panic(err)
	}

	DB = client.Database(config.Cfg.DbName)
	fmt.Println("Connected to MongoDB")

}
