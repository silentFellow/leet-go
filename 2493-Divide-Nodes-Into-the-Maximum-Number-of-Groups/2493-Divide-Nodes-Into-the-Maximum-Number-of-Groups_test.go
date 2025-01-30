package leetcode

import "testing"

func Test_magnificentSets(t *testing.T) {
	type args struct {
		n     int
		edges [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test case 1",
			args: args{
				n: 6,
				edges: [][]int{
					{1, 2},
					{1, 4},
					{1, 5},
					{2, 6},
					{2, 3},
					{4, 6},
				},
			},
			want: 4,
		},
		{
			name: "test case 2",
			args: args{
				n: 3,
				edges: [][]int{
					{1, 2},
					{2, 3},
					{3, 1},
				},
			},
			want: -1,
		},
		{
			name: "test case 3",
			args: args{
				n: 5,
				edges: [][]int{
					{1, 2},
					{2, 3},
					{3, 4},
					{4, 5},
				},
			},
			want: 5,
		},
		{
			name: "test case 4",
			args: args{
				n: 4,
				edges: [][]int{
					{1, 2},
					{2, 3},
					{3, 4},
				},
			},
			want: 4,
		},
		{
			name: "test case 5",
			args: args{
				n: 1,
				edges: [][]int{},
			},
			want: 1,
		},
		{
			name: "test case 6",
			args: args{
				n: 6,
				edges: [][]int{
					{1, 2},
					{2, 3},
					{4, 5},
					{5, 6},
				},
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := magnificentSets(tt.args.n, tt.args.edges); got != tt.want {
				t.Errorf("magnificentSets() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bfsCheckAndFindHeight(t *testing.T) {
	type args struct {
		start int
		graph [][]int
	}
	tests := []struct {
		name            string
		args            args
		wantComponentID int
		wantMaxHeight   int
		wantOk          bool
	}{
		{
			name: "test case 3",
			args: args{
				start: 1,
				graph: [][]int{
					{1},
					{0, 2},
					{1, 3},
					{2},
				},
			},
			wantComponentID: 0,
			wantMaxHeight:   3,
			wantOk:          true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotComponentID, gotMaxHeight, gotOk := bfsCheckAndFindHeight(tt.args.start, tt.args.graph)
			if gotComponentID != tt.wantComponentID {
				t.Errorf("bfsCheckAndFindHeight() gotComponentID = %v, want %v", gotComponentID, tt.wantComponentID)
			}
			if gotMaxHeight != tt.wantMaxHeight {
				t.Errorf("bfsCheckAndFindHeight() gotMaxHeight = %v, want %v", gotMaxHeight, tt.wantMaxHeight)
			}
			if gotOk != tt.wantOk {
				t.Errorf("bfsCheckAndFindHeight() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func Test_absDiff(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test case 1",
			args: args{
				a: 5,
				b: 3,
			},
			want: 2,
		},
		{
			name: "test case 2",
			args: args{
				a: 3,
				b: 5,
			},
			want: 2,
		},
		{
			name: "test case 3",
			args: args{
				a: 7,
				b: 7,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := absDiff(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("absDiff() = %v, want %v", got, tt.want)
			}
		})
	}
}
