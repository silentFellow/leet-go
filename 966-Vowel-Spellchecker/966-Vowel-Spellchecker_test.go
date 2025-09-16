package leetcode

import (
	"reflect"
	"testing"
)

func Test_spellchecker(t *testing.T) {
	type args struct {
		wordlist []string
		queries  []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "example1",
			args: args{
				wordlist: []string{"KiTe", "kite", "hare", "Hare"},
				queries: []string{
					"kite",
					"Kite",
					"KiTe",
					"Hare",
					"HARE",
					"Hear",
					"hear",
					"keti",
					"keet",
					"keto",
				},
			},
			want: []string{"kite", "KiTe", "KiTe", "Hare", "hare", "", "", "KiTe", "", "KiTe"},
		},
		{
			name: "example2",
			args: args{
				wordlist: []string{"yellow"},
				queries:  []string{"YellOw"},
			},
			want: []string{"yellow"},
		},
		{
			name: "no match",
			args: args{
				wordlist: []string{"apple"},
				queries:  []string{"banana"},
			},
			want: []string{""},
		},
		{
			name: "case insensitive",
			args: args{
				wordlist: []string{"Apple"},
				queries:  []string{"apple", "APPLE", "aPpLe"},
			},
			want: []string{"Apple", "Apple", "Apple"},
		},
		{
			name: "vowel error",
			args: args{
				wordlist: []string{"KiTe"},
				queries:  []string{"kete", "kuto", "kita"},
			},
			want: []string{"KiTe", "KiTe", "KiTe"},
		},
		{
			name: "multiple vowels",
			args: args{
				wordlist: []string{"poe", "paa"},
				queries:  []string{"pue"},
			},
			want: []string{"poe"},
		},
		{
			name: "first match precedence",
			args: args{
				wordlist: []string{"kite", "KiTe"},
				queries:  []string{"KITE"},
			},
			want: []string{"kite"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := spellchecker(tt.args.wordlist, tt.args.queries); !reflect.DeepEqual(
				got,
				tt.want,
			) {
				t.Errorf("spellchecker() = %v, want %v", got, tt.want)
			}
		})
	}
}
