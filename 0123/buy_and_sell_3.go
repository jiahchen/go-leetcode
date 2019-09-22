package main

import (
	"fmt"
	"math"
)

func main() {
	input := []int{1, 2, 4, 2, 5, 7, 2, 4, 9, 0}
	ans := maxProfit(input)
	fmt.Println("Ans: ", ans)
}

func maxProfit(prices []int) int {
	firstbuy, firstsell, secbuy, secsell := math.MinInt32, 0, math.MinInt32, 0

	for i := 0; i < len(prices); i++ {
		firstbuy = max(firstbuy, -prices[i])
		firstsell = max(firstsell, prices[i]+firstbuy)
		secbuy = max(secbuy, firstsell-prices[i])
		secsell = max(secsell, prices[i]+secbuy)
	}
	return secsell
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
