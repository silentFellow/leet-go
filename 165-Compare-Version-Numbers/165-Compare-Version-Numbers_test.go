package leetcode

import "testing"

func Test_compareVersion(t *testing.T) {
	type args struct {
		version1 string
		version2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"equal simple", args{"1.0", "1.0.0"}, 0},
		{"less simple", args{"1.2", "1.10"}, -1},
		{"equal leading zeros", args{"1.01", "1.001"}, 0},
		{"greater", args{"2.0", "1.9"}, 1},
		{"less with trailing zeros", args{"1.0.0", "1.0.1"}, -1},
		{"greater with trailing zeros", args{"1.0.1", "1.0.0"}, 1},
		{"equal single", args{"0", "0"}, 0},
		{"greater single", args{"6", "003479002"}, -1},
		{"less long", args{"306.9.0.0", "3.0691430.7.7.263.005.2.1.6.606.702.7.076519166.507"}, 1},
		{
			"equal long trailing zeros",
			args{
				"30307.038",
				"30307.038.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0",
			},
			0,
		},
		{"less", args{"1", "1.1.0.1.0"}, -1},
		{"greater", args{"101.1.100", "100.1.101"}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compareVersion(tt.args.version1, tt.args.version2); got != tt.want {
				t.Errorf("compareVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}
