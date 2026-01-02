package leetcode

import (
	"slices"
	"strings"
)

func minDeletionSize(strs []string) int {
	ans := 0

	colStrsBuilder := make([]strings.Builder, len(strs[0]))
	for _, str := range strs {
		for j, char := range str {
			colStrsBuilder[j].WriteRune(char)
		}
	}

	colStrs := make([]string, len(colStrsBuilder))
	for i, builder := range colStrsBuilder {
		colStrs[i] = builder.String()
	}

	for _, str := range colStrs {
		runes := []rune(str)
		slices.Sort(runes)
		sortedStr := string(runes)

		if str != sortedStr {
			ans++
		}
	}

	return ans
}

// NOTE: Easy Way
// func minDeletionSize(strs []string) int {
//     ans := 0
//     for col := 0; col < len(strs[0]); col++ {
//         for row := 1; row < len(strs); row++ {
//             if strs[row][col] < strs[row-1][col] {
//                 ans++
//                 break
//             }
//         }
//     }
//     return ans
// }
