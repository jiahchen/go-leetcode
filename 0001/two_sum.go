package main

import "fmt"

func main() {
	input := []int{1, 3, 5, 7}
	target := 12
	fmt.Println("Ans: ", twoSum(input, target))
}

func twoSum(nums []int, target int) []int {
	match := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		var tmp = target - nums[i]
		if match[nums[i]] > 0 {
			return setAns(match[nums[i]]-1, i)
		}
		match[tmp] = i + 1
	}
	return setAns(0, 0)
}

func setAns(first int, sec int) []int {
	return []int{first, sec}
}
