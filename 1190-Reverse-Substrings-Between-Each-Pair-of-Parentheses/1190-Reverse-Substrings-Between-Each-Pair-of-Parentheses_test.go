package leetcode

import "testing"

func Test_reverseParentheses(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Example 1",
			args: args{s: "(abcd)"},
			want: "dcba",
		},
		{
			name: "Example 2",
			args: args{s: "(u(love)i)"},
			want: "iloveu",
		},
		{
			name: "Example 3",
			args: args{s: "(ed(et(oc))el)"},
			want: "leetcode",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseParentheses(tt.args.s); got != tt.want {
				t.Errorf("reverseParentheses() = %v, want %v", got, tt.want)
			}
		})
	}
}
