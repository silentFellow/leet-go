package leetcode

import "testing"

func Test_restoreString(t *testing.T) {
	type args struct {
		s       string
		indices []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "example1",
			args: args{s: "codeleet", indices: []int{4, 5, 6, 7, 0, 2, 1, 3}},
			want: "leetcode",
		},
		{
			name: "example2",
			args: args{s: "abc", indices: []int{0, 1, 2}},
			want: "abc",
		},
		{
			name: "reverse",
			args: args{s: "abc", indices: []int{2, 1, 0}},
			want: "cba",
		},
		{
			name: "single char",
			args: args{s: "a", indices: []int{0}},
			want: "a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := restoreString(tt.args.s, tt.args.indices); got != tt.want {
				t.Errorf("restoreString() = %v, want %v", got, tt.want)
			}
		})
	}
}
