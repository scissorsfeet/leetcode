package Array

import "testing"

func TestNumIslands(t *testing.T) {
	nums := [][]byte{
		{1,1,1,1,0},
		{1,1,0,1,0},
		{1,1,0,0,0},
		{0,0,0,0,0},
	}
	t.Log(numIslands(nums))
}
