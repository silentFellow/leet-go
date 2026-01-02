package leetcode

import "testing"

func Test_minDeletionSize(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "already sorted, no deletion needed",
			args: args{[]string{"abc", "bcd", "cde"}},
			want: 0,
		},
		{
			name: "all columns need deletion",
			args: args{[]string{"zyx", "wvu", "tsr"}},
			want: 3,
		},
		{
			name: "single column, already sorted",
			args: args{[]string{"a", "a"}},
			want: 0,
		},
		{
			name: "single column, not sorted",
			args: args{[]string{"b", "a"}},
			want: 1,
		},
		{
			name: "example 1",
			args: args{[]string{"ca", "bb", "ac"}},
			want: 1,
		},
		{
			name: "example 2",
			args: args{[]string{"xc", "yb", "za"}},
			want: 0,
		},
		{
			name: "example 3",
			args: args{[]string{"zyx", "wvu", "tsr"}},
			want: 3,
		},
		{
			name: "partial deletion needed",
			args: args{[]string{"abx", "agz", "bgc", "bfc"}},
			want: 1,
		},
		{
			name: "tricky case",
			args: args{[]string{"xga", "xfb", "yfa"}},
			want: 1,
		},
		{
			name: "random order, multiple deletions",
			args: args{[]string{"vdy", "vei", "zvc", "zld"}},
			want: 2,
		},
		{
			name: "no deletion, all same",
			args: args{[]string{"aaa", "aaa", "aaa"}},
			want: 0,
		},
		{
			name: "minimum input size",
			args: args{[]string{"a"}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDeletionSize(tt.args.strs); got != tt.want {
				t.Errorf("minDeletionSize() = %v, want %v", got, tt.want)
			}
		})
	}
}
