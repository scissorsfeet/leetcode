package IslandPerimeter

import (
    "testing"
    "fmt"
)

func TestA(t *testing.T) {
    var grid [][]int
    grid = append(grid, []int{0, 1, 0, 0})
    grid = append(grid, []int{1, 1, 1, 0})
    grid = append(grid, []int{0, 1, 0, 0})
    grid = append(grid, []int{1, 1, 0, 0})
    ret := islandPerimeter(grid)
    if ret != 16 {
        t.Error(fmt.Sprintf("wrong len:%d", ret))
    }else{
        t.Log(ret)
    }
}

func TestB(t *testing.T) {
    var grid [][]int
    ret := islandPerimeter(grid)
    if ret != 0 {
        t.Error(fmt.Sprintf("wrong len:%d", ret))
    }else{
        t.Log(ret)
    }
}