package mongodb

import (
	"go.mongodb.org/mongo-driver/bson"
	"web-pet-project/internal/dbms/model"
	"web-pet-project/internal/dbms/repository"
)

type issuesRepository struct {
	//
	connectString string
	dbName        string
}

func NewIssuesRepository() repository.IssuesRepository {
	return &issuesRepository{
		//mongodb://localhost:27017/reportapp?authSource=admin
		connectString: "mongodb://admin:secret@localhost:27017/reportapp?authSource=admin", //connectString,
		dbName:        "reportapp",
	}
}

func (repo *issuesRepository) GetAllIssues() ([]model.Issue, error) {
	db, client, context, cancel := SetupMongoDB(repo.connectString, repo.dbName)
	defer CloseConnection(client, context, cancel)

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
