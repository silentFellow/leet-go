package leetcode

import "testing"

func Test_smallestNumber(t *testing.T) {
	type args struct {
		pattern string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test 1",
			args: args{pattern: "IIIDIDDD"},
			want: "123549876",
		},
		{
			name: "Test 2",
			args: args{pattern: "DDD"},
			want: "4321",
		},
		{
			name: "Test 3",
			args: args{pattern: "I"},
			want: "12",
		},
		{
			name: "Test 4",
			args: args{pattern: "D"},
			want: "21",
		},
		{
			name: "Test 5",
			args: args{pattern: "IDID"},
			want: "13254",
		},
		{
			name: "Test 6",
			args: args{pattern: "DDI"},
			want: "3214",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := smallestNumber(tt.args.pattern); got != tt.want {
				t.Errorf("smallestNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
