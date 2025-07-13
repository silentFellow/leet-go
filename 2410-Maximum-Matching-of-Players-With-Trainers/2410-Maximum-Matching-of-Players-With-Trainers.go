package leetcode

import "slices"

func matchPlayersAndTrainers(players []int, trainers []int) int {
	slices.Sort(players)
	slices.Sort(trainers)

	ans := 0
	i, j := 0, 0
	for i < len(players) && j < len(trainers) {
		if players[i] <= trainers[j] {
			ans++
			i++
		}

		j++
	}

	return ans
}
