package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

func SetupMongoDB(uri string, dbName string) (*mongo.Database, *mongo.Client, context.Context, context.CancelFunc, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		fmt.Printf("Could not connect to mongodb service: %v\n\n", err)
		return nil, client, ctx, cancel, err
		//panic(err)
	}

	log.Printf("Mongodb ping")
	ctx, _ = context.WithTimeout(context.Background(), 2*time.Second)
	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		fmt.Printf("Could not ping to mongo db service: %v\n\n", err)
		return nil, client, ctx, cancel, err
	}

	db := client.Database(dbName)
	log.Printf("Mongodb successfully connected to %v", db)
	//return db, client, ctx, nil, nil
	return db, client, ctx, cancel, nil
}

// Close the connection
func CloseConnection(client *mongo.Client, context context.Context, cancel context.CancelFunc) {
	defer func() {
		cancel()
		if err := client.Disconnect(context); err != nil {
			//panic(err)
			log.Fatalf("Mongodb can't close connection %v\n", err)
		}
		log.Fatalln("Close connection is called")
	}()
}
