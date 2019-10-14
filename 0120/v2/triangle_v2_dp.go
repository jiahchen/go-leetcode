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
	if triangle == nil || triangle[0] == nil {
		return -1
	}
	makeMinPath(triangle)
	return findMin(triangle[len(triangle)-1])
}

func makeMinPath(triangle [][]int) {
	for i := 1; i < len(triangle); i++ {
		triangle[i][0] += triangle[i-1][0]
		for j := 1; j < i; j++ {
			triangle[i][j] += min(triangle[i-1][j-1], triangle[i-1][j])
		}
		triangle[i][i] += triangle[i-1][i-1]
		fmt.Println(triangle)
	}
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func findMin(input []int) int {
	min := input[0]
	for i := 1; i < len(input); i++ {
		if input[i] < min {
			min = input[i]
		}
	}
	return min
}
