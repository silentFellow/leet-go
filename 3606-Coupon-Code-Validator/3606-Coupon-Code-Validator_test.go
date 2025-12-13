package leetcode

import (
	"reflect"
	"testing"
)

func Test_validateCoupons(t *testing.T) {
	type args struct {
		code         []string
		businessLine []string
		isActive     []bool
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Example 1",
			args: args{
				code:         []string{"SAVE20", "", "PHARMA5", "SAVE@20"},
				businessLine: []string{"restaurant", "grocery", "pharmacy", "restaurant"},
				isActive:     []bool{true, true, true, true},
			},
			want: []string{"PHARMA5", "SAVE20"},
		},
		{
			name: "Example 2",
			args: args{
				code:         []string{"GROCERY15", "ELECTRONICS_50", "DISCOUNT10"},
				businessLine: []string{"grocery", "electronics", "invalid"},
				isActive:     []bool{false, true, true},
			},
			want: []string{"ELECTRONICS_50"},
		},
		{
			name: "All invalid codes",
			args: args{
				code:         []string{"", "BAD@CODE", "123*"},
				businessLine: []string{"electronics", "grocery", "pharmacy"},
				isActive:     []bool{true, true, true},
			},
			want: []string{},
		},
		{
			name: "Inactive coupons",
			args: args{
				code:         []string{"CODE1", "CODE2"},
				businessLine: []string{"electronics", "grocery"},
				isActive:     []bool{false, false},
			},
			want: []string{},
		},
		{
			name: "Multiple valid, check sorting",
			args: args{
				code:         []string{"B", "A", "C", "D"},
				businessLine: []string{"pharmacy", "electronics", "pharmacy", "grocery"},
				isActive:     []bool{true, true, true, true},
			},
			want: []string{"A", "D", "B", "C"},
		},
		{
			name: "Business line not allowed",
			args: args{
				code:         []string{"CODE1", "CODE2"},
				businessLine: []string{"toys", "electronics"},
				isActive:     []bool{true, true},
			},
			want: []string{"CODE2"},
		},
		{
			name: "Lexicographical order within business line",
			args: args{
				code:         []string{"Z", "A", "M"},
				businessLine: []string{"restaurant", "restaurant", "restaurant"},
				isActive:     []bool{true, true, true},
			},
			want: []string{"A", "M", "Z"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validateCoupons(tt.args.code, tt.args.businessLine, tt.args.isActive); !reflect.DeepEqual(
				got,
				tt.want,
			) {
				t.Errorf("validateCoupons() = %v, want %v", got, tt.want)
			}
		})
	}
}
