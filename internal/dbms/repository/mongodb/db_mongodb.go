package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"web-pet-project/internal/dbms/model"
	"web-pet-project/internal/dbms/repository"
)

type IssuesRepository struct {
	//
	//connectString string
	dbName string
	client *mongo.Client
}

func NewIssuesRepository(cs, dbn string) repository.IssuesRepository {
	cl, _, _, err := SetupMongoDB(cs)
	if err != nil {
		log.Printf("Error connection to mongodb: %v\n", err)
		//return nil, err
	}
	return &IssuesRepository{
		//mongodb://localhost:27017/reportapp?authSource=admin
		//connectString: cs,  //"mongodb://admin:secret@localhost:27017/reportapp?authSource=admin", //connectString,
		dbName: dbn, //"reportapp",
		client: cl,
	}
}

func (repo *IssuesRepository) Close() {
	// TODO
	// CloseConnection(client, context, cancel)
}

func (repo *IssuesRepository) GetAllIssues() ([]model.Issue, error) {

	//defer CloseConnection(client, context, cancel)

	db := repo.client.Database(repo.dbName)
	//log.Printf("Mongodb successfully connected to %v", db)
	collection := db.Collection("issues")

	filter := bson.D{}
	ctx := context.Background()

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		panic(err)
	}

	defer cursor.Close(ctx)

	var issues []model.Issue
	if err = cursor.All(ctx, &issues); err != nil {
		panic(err)
	}

	//for cursor.Next(context) {
	//	//var result User
	//	//err := cursor.Decode(&result)
	//	var result bson.M
	//	err := cursor.Decode(&result)
	//	fmt.Printf("result:= %v", result)
	//	issues = append(issues,
	//		model.Issue{
	//			IssueId:   fmt.Sprint(result["issueId"]),
	//			IssueKey:  fmt.Sprint(result["issueKey"]),
	//			IssueType: fmt.Sprint(result["issueType"]),
	//			Summary:   fmt.Sprint(result["summary"]),
	//		})
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//}

	return issues, nil
}
