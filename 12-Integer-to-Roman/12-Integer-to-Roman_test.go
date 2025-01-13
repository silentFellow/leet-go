package leetcode

import "testing"

func Test_intToRoman(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test 1",
			args: args{num: 3749},
			want: "MMMDCCXLIX",
		},
		{
			name: "Test 2",
			args: args{num: 58},
			want: "LVIII",
		},
		{
			name: "Test 3",
			args: args{num: 1994},
			want: "MCMXCIV",
		},
		{
			name: "Test 4",
			args: args{num: 1},
			want: "I",
		},
		{
			name: "Test 5",
			args: args{num: 3999},
			want: "MMMCMXCIX",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intToRoman(tt.args.num); got != tt.want {
				t.Errorf("intToRoman() = %v, want %v", got, tt.want)
			}
		})
	}
}
