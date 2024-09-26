package mongodb

import (
	"context"
	"github.com/google/go-cmp/cmp"
	"github.com/testcontainers/testcontainers-go/modules/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"reflect"
	"testing"
	"web-pet-project/internal/dbms/model"
)

func Test_issuesRepository_GetAllIssues_tcont(t *testing.T) {
	ctx := context.Background()
	mongodbContainer, err := mongodb.Run(ctx, "mongo:6")
	t.Log(mongodbContainer)

	t.Cleanup(func() {
		if err := mongodbContainer.Terminate(ctx); err != nil {
			t.Fatalf("failed to terminate mongoContainer: %s", err)
		}
	})

	if err != nil {
		t.Fatalf("failed to start container: %s", err)
	}

	endpoint, err := mongodbContainer.ConnectionString(ctx)
	if err != nil {
		t.Fatalf("failed to get connection string: %s", err)
	}

	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(endpoint))

	if err != nil {
		t.Fatalf("failed to connect to MongoDB: %s", err)
	}

	dbNameTest := "test"
	coll := mongoClient.Database(dbNameTest).Collection("issues")

	issues := []interface{}{
		model.Issue{IssueId: "id-1", IssueType: 1, IssueKey: "k1", Summary: "s1"},
		model.Issue{IssueId: "id-2", IssueType: 2, IssueKey: "k2", Summary: "s2"},
	}
	_, err = coll.InsertMany(context.TODO(), issues)

	// convert []interface{} to []model.Issue for comparing
	want := make([]model.Issue, len(issues))
	for i, issue := range issues {
		want[i] = issue.(model.Issue)
	}

	if err != nil {
		t.Fatal(err)
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

		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("mismatch:\n%s", diff)
		}

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
