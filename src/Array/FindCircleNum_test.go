package Array

import "testing"

func TestFindCircleNum(t *testing.T) {
	M := [][]int{
		{1,0,0,1},
		{0,1,1,0},
		{0,1,1,1},
		{1,0,1,1}}
	t.Log(findCircleNum(M))
}
