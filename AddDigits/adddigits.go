package adddigits

import (
	"strconv"
)

/*
Given a non-negative integer num, repeatedly add all its digits until the result has only one digit.
给一个非负整数，重复将其数字相加直到最后成为个位数

For example:
Given num = 38, the process is like: 3 + 8 = 11, 1 + 1 = 2. Since 2 has only one digit, return it.
例如：
对于38：3 + 8 = 11, 1 + 1 = 2，2是个个位数，返回即可。
*/
func addDigits(num int) int {
	if num < 10 { //个位数直接返回
		return num
	}

	numStr := strconv.Itoa(num)
	sum := 11
	for sum >= 10 {
		sum = 0
		for _, b := range numStr {
			tmp, _ := strconv.Atoi(string(b))
			sum += tmp
		}
		numStr = strconv.Itoa(sum)
	}
	return sum
}
