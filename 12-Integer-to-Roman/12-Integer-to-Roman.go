package leetcode

func intToRoman(num int) string {
	return convert(num, "")
}

func convert(num int, roman string) string {
	if num == 0 {
		return roman
	} else if num-1000 >= 0 {
		return convert(num-1000, roman+"M")
	} else if num-900 >= 0 {
		return convert(num-900, roman+"CM")
	} else if num-500 >= 0 {
		return convert(num-500, roman+"D")
	} else if num-400 >= 0 {
		return convert(num-400, roman+"CD")
	} else if num-100 >= 0 {
		return convert(num-100, roman+"C")
	} else if num-90 >= 0 {
		return convert(num-90, roman+"XC")
	} else if num-50 >= 0 {
		return convert(num-50, roman+"L")
	} else if num-40 >= 0 {
		return convert(num-40, roman+"XL")
	} else if num-10 >= 0 {
		return convert(num-10, roman+"X")
	} else if num-9 >= 0 {
		return convert(num-9, roman+"IX")
	} else if num-5 >= 0 {
		return convert(num-5, roman+"V")
	} else if num-4 >= 0 {
		return convert(num-4, roman+"IV")
	} else if num-1 >= 0 {
		return convert(num-1, roman+"I")
	}

	return roman
}
