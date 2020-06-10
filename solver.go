package sudokuSolver

import (
	"context"
	"math"
)

func isValid(grid [][]int, x, y, i int) bool {
	for n := 0; n < 9; n++ {
		if grid[y][n] == i || grid[n][x] == i{
			return false
		}
	}
	topX, topY := int(3 * (math.Floor(float64(x/3)))), int(3 * (math.Floor(float64(y/3))))
	for x := topX; x < topX + 3; x++ {
		for y := topY; y < topY + 3; y++ {
			if grid[y][x] == i {
				return false
			}
		}
	}
	return true
}

func findNextEmptyCell(grid [][]int) (int,int) {
	for y := 0; y < 9; y++ {
		for x := 0; x < 9; x++ {
			if grid[y][x] == 0 {
				return x, y
			}
		}
	}
	return -1, -1
}

func SolveSudoku(grid [][]int, x, y int, ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return false
	default:
	}
	x, y = findNextEmptyCell(grid)
	if x == -1 && y == -1 {
		return true
	}
	for i := 1; i <= 9; i++ {
		if isValid(grid, x, y, i) {
			grid[y][x] = i
			if SolveSudoku(grid, x, y, ctx) {
				return true
			}
			select {
			case <-ctx.Done():
				return false
			default:
			}
			grid[y][x] = 0
		}
	//time.Sleep(100000000)
	}
	return false
}