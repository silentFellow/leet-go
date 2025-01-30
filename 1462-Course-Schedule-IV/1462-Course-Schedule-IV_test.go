package leetcode

import (
	"reflect"
	"testing"
)

func Test_checkIfPrerequisite(t *testing.T) {
	type args struct {
		numCourses    int
		prerequisites [][]int
		queries       [][]int
	}
	tests := []struct {
		name string
		args args
		want []bool
	}{
		{
			name: "example1",
			args: args{
				numCourses:    2,
				prerequisites: [][]int{{1, 0}},
				queries:       [][]int{{0, 1}, {1, 0}},
			},
			want: []bool{false, true},
		},
		{
			name: "example2",
			args: args{
				numCourses:    2,
				prerequisites: [][]int{},
				queries:       [][]int{{1, 0}, {0, 1}},
			},
			want: []bool{false, false},
		},
		{
			name: "example3",
			args: args{
				numCourses:    3,
				prerequisites: [][]int{{1, 2}, {1, 0}, {2, 0}},
				queries:       [][]int{{1, 0}, {1, 2}},
			},
			want: []bool{true, true},
		},
		{
			name: "no_prerequisites",
			args: args{
				numCourses:    4,
				prerequisites: [][]int{},
				queries:       [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 0}},
			},
			want: []bool{false, false, false, false},
		},
		{
			name: "all_prerequisites",
			args: args{
				numCourses:    4,
				prerequisites: [][]int{{0, 1}, {1, 2}, {2, 3}},
				queries:       [][]int{{0, 3}, {1, 3}, {0, 2}, {2, 1}},
			},
			want: []bool{true, true, true, false},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkIfPrerequisite(tt.args.numCourses, tt.args.prerequisites, tt.args.queries); !reflect.DeepEqual(
				got,
				tt.want,
			) {
				t.Errorf("checkIfPrerequisite() = %v, want %v", got, tt.want)
			}
		})
	}
}
