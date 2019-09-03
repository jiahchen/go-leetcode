package main

import "fmt"

func main() {
	input := [][]int{
		{1, 1, 0},
		{1, 1, 1},
		{1, 1, 1},
		{1, 1, 1},
	}
	ans := largest1BorderedSquare(input)
	fmt.Println(ans)
}

func largest1BorderedSquare(grid [][]int) int {
	square := 0
	maxsize := checkMaxSize(grid)

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				if square == 0 {
					square = 1
				}
				for k := square; k < maxsize; k++ {
					if checkBoarder(grid, k, i, j) {
						square = k + 1
					}
				}
			}
		}
	}
	return square * square
}

func checkMaxSize(grid [][]int) int {
	maxsize := len(grid)

	for i := 0; i < len(grid); i++ {
		if len(grid[i]) > maxsize {
			maxsize = len(grid[i])
		}
	}
	return maxsize
}

func checkBoarder(grid [][]int, size int, row int, col int) bool {
	endrow := row + size
	endcol := col + size

	// check endpoint exist
	if endrow > (len(grid) - 1) {
		return false
	}
	if endcol > (len(grid[endrow]) - 1) {
		return false
	}

	// if the endpoint is 0, return folse
	if grid[endrow][endcol] != 1 {
		return false
	}

	// check row all 1 from start/end point
	for i := 0; i <= size; i++ {
		if grid[row][col+i] != 1 || grid[endrow][endcol-i] != 1 {
			return false
		}
	}

	for j := 0; j <= size; j++ {
		if grid[row+j][col] != 1 || grid[endrow-j][endcol] != 1 {
			return false
		}
	}
	return true
}
