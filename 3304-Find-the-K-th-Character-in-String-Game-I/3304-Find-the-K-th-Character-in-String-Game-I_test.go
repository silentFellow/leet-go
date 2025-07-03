package leetcode

import "testing"

func Test_kthCharacter(t *testing.T) {
	type args struct {
		k int
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		{"k=1", args{1}, 'a'},
		{"k=2", args{2}, 'b'},
		{"k=3", args{3}, 'b'},
		{"k=4", args{4}, 'c'},
		{"k=5", args{5}, 'b'},
		{"k=6", args{6}, 'c'},
		{"k=7", args{7}, 'c'},
		{"k=8", args{8}, 'd'},
		{"k=9", args{9}, 'b'},
		{"k=10", args{10}, 'c'},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kthCharacter(tt.args.k); got != tt.want {
				t.Errorf("kthCharacter() = %v, want %v", got, tt.want)
			}
		})
	}
}
