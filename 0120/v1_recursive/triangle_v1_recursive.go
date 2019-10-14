package main

import "fmt"

func main() {
	input := [][]int{
		{2},
		{3, 4},
		{6, 5, 7},
		{4, 1, 8, 3},
	}
	ans := minimumTotal(input)
	fmt.Println(ans)
}

func minimumTotal(triangle [][]int) int {
	return findMinPath(0, 0, triangle)
}

func findMinPath(rooti int, rootj int, triangle [][]int) int {
	if rooti == len(triangle)-1 {
		return triangle[rooti][rootj]
	}
	fmt.Println(min(findMinPath(rooti+1, rootj, triangle),
		findMinPath(rooti+1, rootj+1, triangle)) + triangle[rooti][rootj])
	return min(findMinPath(rooti+1, rootj, triangle),
		findMinPath(rooti+1, rootj+1, triangle)) + triangle[rooti][rootj]
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
