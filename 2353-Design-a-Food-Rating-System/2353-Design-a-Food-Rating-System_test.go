package leetcode

import "testing"

func TestFoodRatings_System(t *testing.T) {
	foods := []string{"kimchi", "miso", "sushi", "moussaka", "ramen", "bulgogi"}
	cuisines := []string{"korean", "japanese", "japanese", "greek", "japanese", "korean"}
	ratings := []int{9, 12, 8, 15, 14, 7}
	fr := Constructor(foods, cuisines, ratings)

	tests := []struct {
		op     string
		arg1   string
		arg2   int
		expect string
	}{
		{"highestRated", "korean", 0, "kimchi"},
		{"highestRated", "japanese", 0, "ramen"},
		{"changeRating", "sushi", 16, ""},
		{"highestRated", "japanese", 0, "sushi"},
		{"changeRating", "ramen", 16, ""},
		{"highestRated", "japanese", 0, "ramen"},
	}

	for _, tt := range tests {
		switch tt.op {
		case "highestRated":
			got := fr.HighestRated(tt.arg1)
			if got != tt.expect {
				t.Errorf("HighestRated(%q) = %q, want %q", tt.arg1, got, tt.expect)
			}
		case "changeRating":
			fr.ChangeRating(tt.arg1, tt.arg2)
		}
	}
}
