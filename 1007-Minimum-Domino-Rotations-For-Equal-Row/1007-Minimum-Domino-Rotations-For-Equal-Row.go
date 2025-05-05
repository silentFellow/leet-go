package leetcode

func minDominoRotations(tops []int, bottoms []int) int {
	// val => [total, top, bottom]
	hmap := make(map[int][3]int)

	// len(tops) == len(bottoms)
	for i := range len(tops) {
		tv, bv := tops[i], bottoms[i]

		if _, ok := hmap[tv]; !ok {
			hmap[tv] = [3]int{}
		}

		if _, ok := hmap[bv]; !ok {
			hmap[bv] = [3]int{}
		}

		if tv == bv {
			temp := hmap[tv]
			temp[0]++
			temp[1]++
			temp[2]++
			hmap[tv] = temp
		} else {
			t1, t2 := hmap[tv], hmap[bv]
			t1[0]++
			t1[1]++

			t2[0]++
			t2[2]++
			hmap[tv], hmap[bv] = t1, t2
		}
	}

	// check not possible case
	flag := false
	for _, v := range hmap {
		if v[0] == len(tops) {
			flag = true
			break
		}
	}
	if !flag {
		return -1
	}

	// find answer
	maxv := 0
	for _, v := range hmap {
		if v[0] == len(tops) {
			maxv = max(maxv, max(v[1], v[2]))
		}
	}

	return len(tops) - maxv
}
