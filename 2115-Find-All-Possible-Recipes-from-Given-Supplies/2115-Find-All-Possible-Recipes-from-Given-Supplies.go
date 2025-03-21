package leetcode

// TODO: try to learn and use topological sort

func findAllRecipes(recipes []string, ingredients [][]string, supplies []string) []string {
	canCook := make(map[string]bool)
	recipeIdx := make(map[string]int)

	for _, v := range supplies {
		canCook[v] = true
	}

	for i, v := range recipes {
		recipeIdx[v] = i
	}

	var dfs func(recipe string) bool
	dfs = func(recipe string) bool {
		if v, ok := canCook[recipe]; ok {
			return v
		}

		if _, ok := recipeIdx[recipe]; !ok {
			return false
		}

		canCook[recipe] = false
		for _, ing := range ingredients[recipeIdx[recipe]] {
			if !dfs(ing) {
				return false
			}
		}

		canCook[recipe] = true
		return true
	}

	ans := make([]string, 0)
	for _, v := range recipes {
		if dfs(v) {
			ans = append(ans, v)
		}
	}

	return ans
}
