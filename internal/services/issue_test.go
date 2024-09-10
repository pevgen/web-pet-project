package services

import (
	"reflect"
	"testing"
)

func TestGetIssueList(t *testing.T) {
	tests := []struct {
		name string
		want []Issue
	}{
		{name: "Check content",
			want: []Issue{
				{"1", "k1", "type 1", "Рус"},
				{"2", "k2", "type 2", "Рус 2"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetIssueList()
			if len(got) != 2 {
				t.Errorf("Length of issues is different for GetIssueList() = %v, want %v", len(got), 2)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetIssueList() = %v, want %v", got, tt.want)
			}
		})
	}
}
