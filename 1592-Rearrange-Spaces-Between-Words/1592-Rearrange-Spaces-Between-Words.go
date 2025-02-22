package leetcode

import "strings"

func reorderSpaces(text string) string {
	var ans strings.Builder

	space := strings.Count(text, " ")

	words := strings.Fields(text)

	if len(words) == 1 {
		ans.WriteString(words[0])
		ans.WriteString(getSpaceString(space))

		return ans.String()
	}

	var between int = space / (len(words) - 1)
	extra := space % (len(words) - 1)

	s := getSpaceString(between)
	ans.WriteString(strings.Join(words, s))
	ans.WriteString(getSpaceString(extra))

	return ans.String()
}

func getSpaceString(count int) string {
	var spaceStr strings.Builder
	for range count {
		spaceStr.WriteString(" ")
	}

	return spaceStr.String()
}
