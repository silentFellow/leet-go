package leetcode

import "testing"

func Test_reorderSpaces(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "example1",
			args: args{text: "  this   is  a sentence "},
			want: "this   is   a   sentence",
		},
		{
			name: "example2",
			args: args{text: " practice   makes   perfect"},
			want: "practice   makes   perfect ",
		},
		{
			name: "noExtraSpaces",
			args: args{text: "hello world"},
			want: "hello world",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reorderSpaces(tt.args.text); got != tt.want {
				t.Errorf("reorderSpaces() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getSpaceString(t *testing.T) {
	type args struct {
		count int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSpaceString(tt.args.count); got != tt.want {
				t.Errorf("getSpaceString() = %v, want %v", got, tt.want)
			}
		})
	}
}
