package IslandPerimeter

/*
You are given a map in form of a two-dimensional integer grid where 1 represents land and 0 represents water.
一个二维表格，其中1表示陆地，0表示海水.
Grid cells are connected horizontally/vertically (not diagonally).
每个格子水平\垂直相连，对角不算相连
The grid is completely surrounded by water, and there is exactly one island (i.e., one or more connected land cells).
表格被海水完全保卫，只有一块陆地
The island doesn't have "lakes" (water inside that isn't connected to the water around the island).
陆地内部不会出现河水
One cell is a square with side length 1. The grid is rectangular, width and height don't exceed 100.
表格的宽高都不超过100，且每个格子长度为1
Determine the perimeter of the island.
计算陆地的周长
*/

func islandPerimeter(grid [][]int) int {
    height := len(grid)
    if height == 0 {
        return 0
    }
    width := len(grid[0])
    var cells []*Cell
    for i:=0; i<height; i++{
        for j:=0; j<width; j++ {
            var left, right, above, bottom int
            if j == 0 {
                left = 0
            }else{
                left = grid[i][j-1]
            }
            if j == width - 1 {
                right = 0
            }else{
                right = grid[i][j+1]
            }
            if i == 0{
                above = 0
            }else{
                above = grid[i-1][j]
            }
            if i == height - 1 {
                bottom = 0
            }else{
                bottom = grid[i+1][j]
            }

            pCell := NewCell(grid[i][j], above, bottom, left, right)
            cells = append(cells, pCell)
        }
    }

    pGrid := NewGrid(width, height, cells)
    return pGrid.Perimeter()
}

type Grid struct {
    Width int
    Height int
    Cells []*Cell
}

func NewGrid(width, height int, cells []*Cell) (pGrid *Grid){
    return &Grid{Width:width, Height:height, Cells:cells}
}

func (this *Grid) Perimeter() (perimeter int) {
    for i, _ := range this.Cells {
        perimeter += this.Cells[i].Perimeter()
    }

    return
}

type Cell struct {
    Value int
    Above int
    Bottom int
    Left int
    Right int
}

func NewCell(value, above, bottom, left, right int) (pCell *Cell) {
    return &Cell{Value:value, Above:above, Bottom:bottom, Left:left, Right:right}
}

func(this *Cell) Perimeter() (perimeter int){
    if this.Value == 0 {
        perimeter = 0
    }else {
        perimeter = 4//默认长度
        if this.Above == 1 {
            perimeter -= 1
        }
        if this.Bottom == 1 {
            perimeter -= 1
        }
        if this.Left == 1 {
            perimeter -= 1
        }
        if this.Right == 1 {
            perimeter -= 1
        }
    }
    return
}