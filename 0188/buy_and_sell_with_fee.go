package main

import "fmt"

func main() {
	input := []int{3, 2, 6, 5, 0, 3, 5, 7, 3, 4, 6, 2, 8}
	time := 2
	ans := maxProfit(input, time)
	fmt.Println("Ans: ", ans)
}

func maxProfit(prices []int, k int) int {
	if len(prices) <= 0 {
		return 0
	}
	buy := prices[0]
	t := [][]int{}
	t1 := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	t2 := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	t3 := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	t = append(t, t1)
	t = append(t, t2)
	t = append(t, t3)

	for i := 1; i < k+1; i++ {
		buy = prices[0]
		for j := 1; j < len(prices); j++ {
			t[i][j] = max(t[i][j-1], prices[j]-buy)
			buy = min(buy, prices[j]-t[i-1][j-1])
			fmt.Println(t[i][j], buy)
		}
		fmt.Println(t[i])
		fmt.Println("======================")
	}
	return t[k][len(prices)-1]
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
