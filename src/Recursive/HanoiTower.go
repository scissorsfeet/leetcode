package Recursive

import "fmt"

/*
汉诺塔问题：
有a、b、c三个柱子，a柱子上从小到大码放n个盘子
要将盘子从a移动到c，依然从小到大码放
移动期间小盘子不能放到大盘子下面

1.问题拆分子问题
如果要将n个盘子从from柱借助assist柱移动到to柱，需要分三步
第一步：将第n个盘子上面的n-1个盘子（即n上面的盘子）从from柱子借助to柱移动到assist柱
第二步：将第n个盘子从from柱移动到to柱
第三步：将之前在assist柱子上面的n-1盘在从assist柱借助from柱移动到to柱

2.公式
move(n, from, to, assist)=
move(n-1, from, assist, to)
move n 从 from 到 to
move(n-1, assist, to, from)

3.终止条件
当只剩下一个盘子时，只需要将这个盘子从from柱移动到to柱
 */
func Hanoi(n uint) {
	if n == 0 {
		return
	}
	move(n, "a", "c", "b")
}

func move(num uint, from, to, assist string) {
	if num == 1 {
		fmt.Printf("move %+v from %+v to %+v\n", 1, from, to)
		return
	}
	move(num-1, from, assist, to)
	fmt.Printf("move %+v from %+v to %+v\n", num, from, to)
	move(num-1, assist, to, from)
}
