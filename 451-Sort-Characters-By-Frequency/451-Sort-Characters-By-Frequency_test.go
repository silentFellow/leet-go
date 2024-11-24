package leetcode

import "testing"

func isValidOutput(got, want string) bool {
	if len(got) != len(want) {
		return false
	}
	gotMap := make(map[rune]int)
	wantMap := make(map[rune]int)
	for _, char := range got {
		gotMap[char]++
	}
	for _, char := range want {
		wantMap[char]++
	}
	for key, val := range wantMap {
		if gotMap[key] != val {
			return false
		}
	}
	return true
}

func Test_frequencySort(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "example1",
			args: args{s: "tree"},
			want: "eert",
		},
		{
			name: "example2",
			args: args{s: "cccaaa"},
			want: "aaaccc",
		},
		{
			name: "example3",
			args: args{s: "Aabb"},
			want: "bbAa",
		},
		{
			name: "empty string",
			args: args{s: ""},
			want: "",
		},
		{
			name: "single character",
			args: args{s: "a"},
			want: "a",
		},
		{
			name: "all unique characters",
			args: args{s: "abcd"},
			want: "abcd",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := frequencySort(tt.args.s); !isValidOutput(got, tt.want) {
				t.Errorf("frequencySort() = %v, want %v", got, tt.want)
			}
		})
	}
}
