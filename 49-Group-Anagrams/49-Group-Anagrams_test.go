package leetcode

import (
	"reflect"
	"sort"
	"testing"
)

func Test_groupAnagrams(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "example1",
			args: args{strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"}},
			want: [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}},
		},
		{
			name: "example2",
			args: args{strs: []string{""}},
			want: [][]string{{""}},
		},
		{
			name: "example3",
			args: args{strs: []string{"a"}},
			want: [][]string{{"a"}},
		},
		{
			name: "example4",
			args: args{strs: []string{"abc", "bca", "cab", "xyz", "zyx"}},
			want: [][]string{{"abc", "bca", "cab"}, {"xyz", "zyx"}},
		},
		{
			name: "example5",
			args: args{strs: []string{"", "b", "bb", "bbb"}},
			want: [][]string{{""}, {"b"}, {"bb"}, {"bbb"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := groupAnagrams(tt.args.strs)
			sortGroups(got)
			sortGroups(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("groupAnagrams() = %v, want %v", got, tt.want)
			}
		})
	}
}

func sortGroups(groups [][]string) {
	for _, group := range groups {
		sort.Strings(group)
	}
	sort.Slice(groups, func(i, j int) bool {
		return groups[i][0] < groups[j][0]
	})
}
