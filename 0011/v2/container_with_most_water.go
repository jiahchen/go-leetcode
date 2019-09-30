package main

import "fmt"

func main() {
	input := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println("Ans: ", maxArea(input))
}

func maxArea(height []int) int {
	tmp := 0
	ans := 0
	i, j := 0, len(height)-1
	for {
		if height[i] > height[j] {
			tmp = height[j] * (j - i)
			j--
		} else {
			tmp = height[i] * (j - i)
			i++
		}
		if tmp > ans {
			ans = tmp
		}
		if j-i == 0 {
			break
		}
	}
	return ans
}
