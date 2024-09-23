package postgres

import (
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"reflect"
	"testing"
	"web-pet-project/internal/dbms/model"
)

func TestIssuesRepository_GetAllIssues(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error occurred while creating mock: %s", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"issue_id", "issue_key", "issue_type", "summary"}).
		AddRow("id-1", "key-1", "1", "summary-1")

	mock.ExpectPing()
	mock.ExpectQuery("SELECT issue_id, issue_key,issue_type,summary FROM issues").
		WillReturnRows(rows)

	repo := &IssuesRepository{
		db: db,
	}
	got, err := repo.GetAllIssues()
	if err != nil {
		t.Errorf("GetAllIssues() error = %v", err)
		return
	}

	want := []model.Issue{
		{"id-1", "key-1", 1, "summary-1"},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("GetAllIssues() got = %v, want %v", got, want)
	}

	fmt.Printf("got: %v", got)
}

//func TestIssuesRepository_GetAllIssues(t *testing.T) {
//	type fields struct {
//		psqlInfo string
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		want    []model.Issue
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			repo := &IssuesRepository{
//				psqlInfo: tt.fields.psqlInfo,
//			}
//			got, err := repo.GetAllIssues()
//			if (err != nil) != tt.wantErr {
//				t.Errorf("GetAllIssues() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("GetAllIssues() got = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

//func TestNewIssueRepository(t *testing.T) {
//	type args struct {
//		uri string
//	}
//	tests := []struct {
//		name string
//		args args
//		want repository.IssuesRepository
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := NewIssueRepository(tt.args.uri); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("NewIssueRepository() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
