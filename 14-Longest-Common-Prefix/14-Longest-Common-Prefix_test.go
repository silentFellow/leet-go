package leetcode

import "testing"

func Test_longestCommonPrefix(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "common prefix 'fl'",
			args: args{strs: []string{"flower", "flow", "flight"}},
			want: "fl",
		},
		{
			name: "no common prefix",
			args: args{strs: []string{"dog", "racecar", "car"}},
			want: "",
		},
		{
			name: "common prefix 'a'",
			args: args{strs: []string{"apple", "ape", "april"}},
			want: "ap",
		},
		{
			name: "single string",
			args: args{strs: []string{"single"}},
			want: "single",
		},
		{
			name: "array with empty string",
			args: args{strs: []string{"", "b", "c"}},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonPrefix(tt.args.strs); got != tt.want {
				t.Errorf("longestCommonPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
