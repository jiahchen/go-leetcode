package main

import "fmt"

func main() {
	input := []int{2, 7, 11, 5}
	target := 9
	fmt.Println("Ans: ", twoSum(input, target))
}

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return setAns(i, j)
			}
		}
	}
	return setAns(0, 0)
}

func setAns(first int, sec int) []int {
	return []int{first, sec}
}
