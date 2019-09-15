package main

import "fmt"

func main() {
	input := []int{7, 6, 4, 3, 1}
	ans := maxProfit(input)
	fmt.Println("Ans: ", ans)
}

func maxProfit(prices []int) int {
	if len(prices) <= 0 {
		return 0
	}
	buyPoint := prices[0]
	ans := 0

	for i := 1; i < len(prices); i++ {
		if prices[i] < buyPoint {
			buyPoint = prices[i]
		}
		if prices[i]-buyPoint > ans {
			ans = prices[i] - buyPoint
		}
	}
	return ans
}
