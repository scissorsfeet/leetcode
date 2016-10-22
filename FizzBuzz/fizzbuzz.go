package FizzBuzz

import "strconv"

/*
Write a program that outputs the string representation of numbers from 1 to n.
输入为从1到n的程序
But for multiples of three it should output “Fizz” instead of the number
3的倍数输出Fizz
and for the multiples of five output “Buzz”.
5的倍数输出Buzz
For numbers which are multiples of both three and five output “FizzBuzz”.
3和5的倍数输出FizzBuzz
 */

func fizzBuzz(n int) []string {
	if n <= 0 {
		return []string{}
	}
	ret := make([]string, 0, n)
	for i:=1;i<=n;i++{
		ret = append(ret, compute(i))
	}
	return ret
}

func compute(i int) (s string){
	s = strconv.Itoa(i)
	if i % 3 == 0 && i % 5 == 0{
		s = "FizzBuzz"
	}else if i % 3 == 0 {
		s = "Fizz"
	}else if i % 5 == 0{
		s = "Buzz"
	}
	return s
}

