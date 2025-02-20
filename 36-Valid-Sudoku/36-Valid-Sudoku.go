package leetcode

func isValidSudoku(board [][]byte) bool {
	n := 9

	rows := make([]map[byte]struct{}, n)
	cols := make([]map[byte]struct{}, n)
	box := make([]map[byte]struct{}, n)

	for i := range n {
		rows[i] = make(map[byte]struct{})
		cols[i] = make(map[byte]struct{})
		box[i] = make(map[byte]struct{})
	}

	for i, row := range board {
		for j, col := range row {
			v := col

			if v == '.' {
				continue
			}

			if _, ok := rows[i][v]; ok {
				return false
			}
			rows[i][v] = struct{}{}

			if _, ok := cols[j][v]; ok {
				return false
			}
			cols[j][v] = struct{}{}

			var idx int = (i / 3) + (j/3)*3
			if _, ok := box[idx][v]; ok {
				return false
			}
			box[idx][v] = struct{}{}
		}
	}

	return true
}
