package leetcode

func maxChunksToSorted(arr []int) int {
	chunks := 0

  // calculate sum till the point
	prevsum := 0
	for i, val := range arr {
		target := (i * (i + 1)) / 2 // target sum if arr is sorted
		prevsum += val

		if target == prevsum { // if matches then sorted chunk
			chunks++
		}
	}

	return chunks
}
