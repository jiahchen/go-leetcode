package main

import "fmt"

func main() {
	input := 21120
	ans := isPalindrome(input)
	fmt.Println("Ans: ", ans)
}

func isPalindrome(x int) bool {
	if x < 0 || (x > 10 && x%10 == 0) {
		return false
	}
	tmpReverse := 0
	for i := x; i >= 0; i /= 10 {
		if tmpReverse > i {
			if tmpReverse/10 == i {
				return true
			}
			break
		}
		if i == tmpReverse {
			return true
		}
		tmpReverse = tmpReverse*10 + (i % 10)
	}
	return false
}
