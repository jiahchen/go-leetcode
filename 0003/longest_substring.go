package main

import "fmt"

func main() {
	s := "aaps"
	ans := lengthOfLongestSubstring(s)
	fmt.Println("Ans: ", ans)
}

func lengthOfLongestSubstring(s string) int {
	maxAns := 1
	tmpAns := 1
	start := 0

	if len(s) == 0 {
		return 0
	}
	for end := 0; end < len(s); end++ {
		if start != end {
			checkPoint := checkDuplicatePoint(s, start, end)
			if checkPoint > -1 {
				if tmpAns > maxAns {
					maxAns = tmpAns
				}
				tmpAns = end - checkPoint
				start = checkPoint + 1
			} else {
				tmpAns++
				if end == len(s)-1 {
					if tmpAns > maxAns {
						maxAns = tmpAns
					}
				}
			}
		}
	}
	return maxAns
}

func checkDuplicatePoint(s string, start int, end int) int {
	for i := start; i < end; i++ {
		if s[end] == s[i] {
			return i
		}
	}
	return -1
}
