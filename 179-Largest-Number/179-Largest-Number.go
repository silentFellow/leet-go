package leetcode

import (
	"container/heap"
	"strconv"
	"strings"
)

type sortNums []string

// sort interface
func (s sortNums) Len() int {
	return len(s)
}

func (s sortNums) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortNums) Less(i, j int) bool {
	return isLarger(s[i], s[j])
}

// heap interface
func (s *sortNums) Push(v any) {
	*s = append(*s, v.(string))
}

func (s *sortNums) Pop() any {
	old := *s
	n := len(old)
	x := old[n-1]
	*s = old[:n-1]
	return x
}

func largestNumber(nums []int) string {
	sh := &sortNums{}

	heap.Init(sh)
	for _, num := range nums {
		str := strconv.Itoa(num)
		heap.Push(sh, str)
	}

	var ans strings.Builder
	for sh.Len() > 0 {
		top := heap.Pop(sh).(string)

		if top == "0" && ans.String() == "0" {
			break
		}

		ans.WriteString(top)
	}

	return ans.String()
}

func isLarger(str1, str2 string) bool {
	if str1 == str2 {
		return false
	}

	length := len(str1)

	if len(str2) < length {
		length = len(str2)
	}

	i := 0

	for ; i < length; i++ {
		if str1[i] > str2[i] {
			return true
		} else if str1[i] < str2[i] {
			return false
		}
	}

	if i == len(str1) {
		return isLarger(str1, str2[i:])
	}

	return isLarger(str1[i:], str2)
}
