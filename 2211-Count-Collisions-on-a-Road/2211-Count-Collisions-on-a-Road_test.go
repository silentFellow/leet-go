package leetcode

import "testing"

func Test_countCollisions(t *testing.T) {
	type args struct {
		directions string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{"RLRSLL"}, 5},
		{"Example 2", args{"LLRR"}, 0},
		{"All stationary", args{"SSSSS"}, 0},
		{"All right", args{"RRRRR"}, 0},
		{"All left", args{"LLLLL"}, 0},
		{"Mixed no collision", args{"LLSSRR"}, 0},
		{"Mixed collision", args{"RSLLR"}, 3},
		{"Single car", args{"S"}, 0},
		{"Single moving car", args{"R"}, 0},
		{"Edge collision", args{"SRL"}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countCollisions(tt.args.directions); got != tt.want {
				t.Errorf("countCollisions() = %v, want %v", got, tt.want)
			}
		})
	}
}
