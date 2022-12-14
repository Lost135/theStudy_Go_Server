package config

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

var DbClient *mongo.Client

func MongoClient() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := recover()
	DbClient, err = mongo.Connect(ctx, options.Client().ApplyURI(Yml.Mongodb.Addr).SetMaxPoolSize(100))
	CatchFatal(err)
	err = DbClient.Ping(ctx, readpref.Primary())
	CatchFatal(err)
	//defer func() {
	//	if err = DbClient.Disconnect(ctx); err != nil {
	//		panic(err)
	//	}
	//}()
}
