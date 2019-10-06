package main

import "fmt"

func main() {
	input := []int{1, 3, 2, 8, 4, 9}
	fee := 2
	ans := maxProfit(input, fee)
	fmt.Println("Ans: ", ans)
}

func maxProfit(prices []int, fee int) int {
	if len(prices) <= 0 {
		return 0
	}
	buyPoint := prices[0]
	ans := 0

	for i := 1; i < len(prices); i++ {
		if prices[i] < buyPoint {
			buyPoint = prices[i]
		} else {
			if prices[i]-buyPoint > fee {
				ans = ans + prices[i] - buyPoint - fee
				buyPoint = prices[i] - fee
			}
		}
	}
	return ans
}
