package leetcode

import "testing"

func Test_repeatLimitedString(t *testing.T) {
	type args struct {
		s           string
		repeatLimit int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "example1",
			args: args{s: "cczazcc", repeatLimit: 3},
			want: "zzcccac",
		},
		{
			name: "example2",
			args: args{s: "aababab", repeatLimit: 2},
			want: "bbabaa",
		},
		{
			name: "example3",
			args: args{s: "xyutfpopdynbadwtvmxiemmusevduloxwvpkjioizvanetecnuqbqqdtrwrkgt", repeatLimit: 1},
			want: "zyxyxwxwvwvuvuvututstrtrtqpqpqponononmlmkmkjigifiededededcbaba",
		},
		{
			name: "single character",
			args: args{s: "aaaaa", repeatLimit: 1},
			want: "a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := repeatLimitedString(tt.args.s, tt.args.repeatLimit); got != tt.want {
				t.Errorf("repeatLimitedString() = %v, want %v", got, tt.want)
			}
		})
	}
}
