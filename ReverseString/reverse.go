package reverse

/*
Write a function that takes a string as input and returns the string reversed.
函数功能：将输入的字符串进行翻转

Example:
Given s = "hello", return "olleh".
例如：
输入hello，输出olleh
*/

//第一种：牺牲空间复杂度
func reverseString(s string) string {
	strLen := len(s)
	if strLen <= 0 { //空字符串直接返回
		return ""
	}

	//翻转
	retStr := make([]byte, strLen)
	index := strLen - 1
	for _, b := range s {
		retStr[index] = byte(b) //忘记了转换
		index--
	}
	return string(retStr) //转换语法错了，类型（变量）
}

//第二种：指针方式
func reverseStringV2(s string) string {
	strLen := len(s)
	if strLen <= 0 { //空字符串直接返回
		return ""
	}

	//翻转
	bs := []byte(s)
	start := 0
	end := strLen - 1
	for start < end {
		tmp := bs[start]
		bs[start] = bs[end]
		bs[end] = tmp
		start++
		end--
	}
	return string(bs)
}
