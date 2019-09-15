package main

import "fmt"

func main() {
	input := []int{}
	ans := maxProfit(input)
	fmt.Println("Ans: ", ans)
}

func maxProfit(prices []int) int {
	if len(prices) <= 0 {
		return 0
	}
	buyPoint := prices[0]
	sellPoint := 0
	ans := 0

	for i := 1; i < len(prices); i++ {
		if prices[i-1] > prices[i] {
			sellPoint = prices[i-1]
			ans = ans + sellPoint - buyPoint
			buyPoint = prices[i]
		}
		if i == len(prices)-1 {
			sellPoint = prices[i]
			ans = ans + sellPoint - buyPoint
		}
	}
	return ans
}
