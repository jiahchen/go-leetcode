package main

import "fmt"

func main() {
	input := 546736898
	ans := reverse(input)
	fmt.Println("Ans: ", ans)
}

func reverse(x int) int {
	tmpnum := x
	if x < 0 {
		tmpnum = 0 - x
	}
	numArray := itoa(tmpnum)
	ansArray := reverseArray(numArray)
	ans := atoi(ansArray)
	if x < 0 {
		ans = 0 - ans
	}
	if ans > 2147483647 || ans < -2147483648 {
		return 0
	}
	return ans
}

func itoa(num int) [10]int {
	var output [10]int
	i := 0
	for {
		if num == 0 {
			break
		}
		output[i] = num % 10
		num = num / 10
		i++
	}
	return output
}

func reverseArray(input [10]int) [10]int {
	var output [10]int
	firstNum := false
	j := 0
	for i := 9; i >= 0; i-- {
		if input[i] != 0 && firstNum == false {
			firstNum = true
		}
		if firstNum == true {
			output[j] = input[i]
			j++
		}
	}
	return output
}

func atoi(input [10]int) int {
	ans := 0
	for i := 0; i < 10; i++ {
		ans = ans + multiple(input[i], i)
	}
	return ans
}

func multiple(num int, times int) int {
	for i := 0; i < times; i++ {
		num = num * 10
	}
	return num
}
