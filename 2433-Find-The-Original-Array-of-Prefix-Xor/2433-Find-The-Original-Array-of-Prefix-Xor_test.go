package leetcode

import (
	"reflect"
	"testing"
)

func Test_findArray(t *testing.T) {
	type args struct {
		pref []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Test case 1",
			args: args{pref: []int{5, 2, 0, 3, 1}},
			want: []int{5, 7, 2, 3, 2},
		},
		{
			name: "Test case 2",
			args: args{pref: []int{13}},
			want: []int{13},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findArray(tt.args.pref); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
