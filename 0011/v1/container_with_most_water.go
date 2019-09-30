package main

import "fmt"

func main() {
	input := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println("Ans: ", maxArea(input))
}

func maxArea(height []int) int {
	tmp := 0
	ans := 0
	for i := 0; i < len(height)-1; i++ {
		for j := i + 1; j < len(height); j++ {
			if height[i] > height[j] {
				tmp = height[j] * (j - i)
			} else {
				tmp = height[i] * (j - i)
			}
			if tmp > ans {
				ans = tmp
			}
		}
	}
	return ans
}
