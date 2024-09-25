package mongodb

import (
	"context"
	"fmt"
	"github.com/testcontainers/testcontainers-go/modules/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"log"
	"reflect"
	"testing"
	"web-pet-project/internal/dbms/model"
)

func Test_issuesRepository_GetAllIssues_tcont(t *testing.T) {
	ctx := context.Background()
	mongodbContainer, err := mongodb.Run(ctx, "mongo:6")
	fmt.Println(mongodbContainer)
	defer func() {
		//if err := testcontainers.TerminateContainer(mongodbContainer); err != nil {
		//	log.Printf("failed to terminate container: %s", err)
		//}
	}()
	if err != nil {
		log.Printf("failed to start container: %s", err)
		return
	}

	endpoint, err := mongodbContainer.ConnectionString(ctx)
	if err != nil {
		log.Printf("failed to get connection string: %s", err)
		return
	}

	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(endpoint))

	if err != nil {
		log.Printf("failed to connect to MongoDB: %s", err)
		return
	}

	dbNameTest := "test"
	coll := mongoClient.Database(dbNameTest).Collection("issues")

	newIssues := []interface{}{
		model.Issue{IssueId: "id-1", IssueType: 1, IssueKey: "k1", Summary: "s1"},
		model.Issue{IssueId: "id-2", IssueType: 2, IssueKey: "k2", Summary: "s2"},
	}
	_, err = coll.InsertMany(context.TODO(), newIssues)

	if err != nil {
		log.Println(err)
	}

	t.Run("mock mongo db", func(t *testing.T) {
		repo := IssuesRepository{
			dbName: dbNameTest,
			client: mongoClient,
		}

		got, err := repo.GetAllIssues()

		if err != nil {
			t.Errorf("GetAllIssues() error = %v", err)
			return
		}

		if len(got) != 2 {
			t.Errorf("Length of issues is different for GetIssueList() = %v, want %v", len(got), 2)
			return
		}

		//if !reflect.DeepEqual(got, tt.want) {
		//	t.Errorf("GetAllIssues() got = %v, want %v", got, tt.want)
		//}
	})

}

func Test_issuesRepository_GetAllIssues(t *testing.T) {
	type fields struct {
		connectString string
	}
	var tests = []struct {
		name    string
		fields  fields
		want    []model.Issue
		wantErr bool
	}{
		{name: "Test 1",
			want:    make([]model.Issue, 2),
			wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := NewIssuesRepository(
				"mongodb://admin:secret@localhost:27017/reportapp?authSource=admin",
				"reportapp")
			got, err := repo.GetAllIssues()

			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllIssues() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) != 2 {
				t.Errorf("Length of issues is different for GetIssueList() = %v, want %v", len(got), 2)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllIssues() got = %v, want %v", got, tt.want)
			}
		})
	}
}
