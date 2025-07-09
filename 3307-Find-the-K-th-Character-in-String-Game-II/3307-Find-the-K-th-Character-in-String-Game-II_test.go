package leetcode

import "testing"

func Test_kthCharacter(t *testing.T) {
	type args struct {
		k          int64
		operations []int
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		{
			name: "all type 0, k=5",
			args: args{k: 5, operations: []int{0, 0, 0}},
			want: 'a',
		},
		{
			name: "mixed 0 and 1, k=10",
			args: args{k: 10, operations: []int{0, 1, 0, 1}},
			want: 'b',
		},
		// {
		// 	name: "single op 0, k=2",
		// 	args: args{k: 2, operations: []int{0}},
		// 	want: 'a',
		// },
		// {
		// 	name: "single op 1, k=2",
		// 	args: args{k: 2, operations: []int{1}},
		// 	want: 'b',
		// },
		// {
		// 	name: "type 1 after 0, k=3",
		// 	args: args{k: 3, operations: []int{0, 1}},
		// 	want: 'b',
		// },
		// {
		// 	name: "type 1 after 0, k=1",
		// 	args: args{k: 1, operations: []int{0, 1}},
		// 	want: 'a',
		// },
		// {
		// 	name: "wrap z to a, k=2",
		// 	args: args{
		// 		k: 2,
		// 		operations: []int{
		// 			1,
		// 			1,
		// 			1,
		// 			1,
		// 			1,
		// 			1,
		// 			1,
		// 			1,
		// 			1,
		// 			1,
		// 			1,
		// 			1,
		// 			1,
		// 			1,
		// 			1,
		// 			1,
		// 			1,
		// 			1,
		// 			1,
		// 			1,
		// 			1,
		// 			1,
		// 			1,
		// 			1,
		// 			1,
		// 			1,
		// 		},
		// 	},
		// 	want: 'a',
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kthCharacter(tt.args.k, tt.args.operations); got != tt.want {
				t.Errorf("kthCharacter() = %v, want %v", got, tt.want)
			}
		})
	}
}
