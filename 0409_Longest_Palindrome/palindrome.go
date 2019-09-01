package main

import "fmt"

func main() {
	input := "aaa"
	palindromeLen := longestPalindrome(input)
	fmt.Println("Ans: ", palindromeLen)
}

func longestPalindrome(s string) int {
	palinlen := 0
	single := 0
	var tmp [65]int

	for i := 0; i < len(s); i++ {
		tmp[s[i]-65]++
	}
	for i := 0; i < len(tmp); i++ {
		if tmp[i] == 1 {
			single = 1
		}
		if tmp[i] > 1 {
			if (tmp[i] % 2) == 1 {
				palinlen = palinlen + tmp[i] - 1
				single = 1
			} else {
				palinlen = palinlen + tmp[i]
			}
		}
	}
	return palinlen + single
}
