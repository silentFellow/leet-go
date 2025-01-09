package leetcode

import "testing"

func Test_prefixCount(t *testing.T) {
	type args struct {
		words []string
		pref  string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test_case_1",
			args: args{
				words: []string{"pay", "attention", "practice", "attend"},
				pref:  "at",
			},
			want: 2,
		},
		{
			name: "test_case_2",
			args: args{
				words: []string{"leetcode", "win", "loops", "success"},
				pref:  "code",
			},
			want: 0,
		},
		{
			name: "test_case_3",
			args: args{
				words: []string{"apple", "app", "application", "apply"},
				pref:  "app",
			},
			want: 4,
		},
		{
			name: "test_case_4",
			args: args{
				words: []string{"banana", "band", "bandana", "ban"},
				pref:  "ban",
			},
			want: 4,
		},
		{
			name: "test_case_5",
			args: args{
				words: []string{"cat", "dog", "cow", "cattle"},
				pref:  "ca",
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := prefixCount(tt.args.words, tt.args.pref); got != tt.want {
				t.Errorf("prefixCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
