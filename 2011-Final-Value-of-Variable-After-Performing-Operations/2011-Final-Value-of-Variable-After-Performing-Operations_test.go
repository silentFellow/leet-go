package leetcode

import "testing"

func Test_finalValueAfterOperations(t *testing.T) {
	type args struct {
		operations []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{[]string{"--X", "X++", "X++"}}, 1},
		{"Example 2", args{[]string{"++X", "++X", "X++"}}, 3},
		{"Example 3", args{[]string{"X++", "++X", "--X", "X--"}}, 0},
		{"All increments", args{[]string{"++X", "X++", "++X", "X++"}}, 4},
		{"All decrements", args{[]string{"--X", "X--", "--X", "X--"}}, -4},
		{"Mixed", args{[]string{"++X", "X--", "X++", "--X"}}, 0},
		{"Single increment", args{[]string{"++X"}}, 1},
		{"Single decrement", args{[]string{"X--"}}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := finalValueAfterOperations(tt.args.operations); got != tt.want {
				t.Errorf("finalValueAfterOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}
