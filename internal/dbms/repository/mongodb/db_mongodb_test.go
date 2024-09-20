package mongodb

import (
	"reflect"
	"testing"
	"web-pet-project/internal/dbms/model"
)

//func TestNewIssuesRepository(t *testing.T) {
//	type args struct {
//		connectString string
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
//			if got := NewIssuesRepository(tt.args.connectString); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("NewIssuesRepository() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

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
