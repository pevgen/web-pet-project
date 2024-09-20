package mongodb

import (
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"web-pet-project/internal/dbms/model"
	"web-pet-project/internal/dbms/repository"
)

type issuesRepository struct {
	//
	connectString string
	dbName        string
}

func NewIssuesRepository(cs, dbn string) repository.IssuesRepository {
	return &issuesRepository{
		//mongodb://localhost:27017/reportapp?authSource=admin
		connectString: cs,  //"mongodb://admin:secret@localhost:27017/reportapp?authSource=admin", //connectString,
		dbName:        dbn, //"reportapp",
	}
}

func (repo *issuesRepository) GetAllIssues() ([]model.Issue, error) {
	db, _, context, _, err := SetupMongoDB(repo.connectString, repo.dbName)
	//defer CloseConnection(client, context, cancel)
	if err != nil {
		log.Printf("Error connection to mongodb: %v\n", err)
		return nil, err
	}

	collection := db.Collection("issues")

	filter := bson.D{}
	cursor, err := collection.Find(context, filter)
	if err != nil {
		panic(err)
	}

	defer cursor.Close(context)

	var issues []model.Issue
	if err = cursor.All(context, &issues); err != nil {
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
