package leetcode

func countPoints(rings string) int {
	hmap := make(map[int][]int)

	for i := 0; i < len(rings); i = i + 2 {
		color := rings[i]
		position := int(rings[i+1] - '0')

		if _, ok := hmap[position]; !ok {
			hmap[position] = make([]int, 3)
		}

		if color == 'R' {
			hmap[position][0]++
		} else if color == 'G' {
			hmap[position][1]++
		} else if color == 'B' {
			hmap[position][2]++
		}
	}

	ans := 0
	for _, val := range hmap {
		if val[0] > 0 && val[1] > 0 && val[2] > 0 {
			ans++
		}
	}

	return ans
}
