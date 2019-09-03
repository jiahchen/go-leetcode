package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	num1 := &ListNode{9, nil}
	num1 = addNum(9, num1)

	num2 := &ListNode{9, nil}

	ans := addTwoNumbers(num1, num2)
	printNum(ans)
}

func addNum(num int, numList *ListNode) *ListNode {
	newNum := &ListNode{num, nil}
	if numList == nil {
		return numList
	}
	for n := numList; n != nil; n = n.Next {
		if n.Next == nil {
			n.Next = newNum
			return numList
		}
	}
	return numList
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	n1 := l1
	n2 := l2
	ans := l1
	checkCarry := 0

	for {
		if n1 == nil {
			if n2 == nil {
				// all digit finish.
				if checkCarry == 1 {
					ans = addNum(1, ans)
				}
				break
			} else {
				// l1 finish, l2 not (l2 > l1 case), change ans to l2
				ans = l2
				n2.Val = n2.Val + checkCarry
				checkCarry = 0
				if n2.Val >= 10 {
					n2.Val = n2.Val % 10
					checkCarry = 1
				}
				n2 = n2.Next
			}
		} else {
			if n2 == nil {
				// l1 not, l2 finish (l1 > l2 case)
				n1.Val = n1.Val + checkCarry
				checkCarry = 0
				if n1.Val >= 10 {
					n1.Val = n1.Val % 10
					checkCarry = 1
				}
				n1 = n1.Next
			} else {
				// l1, l2 not finish
				n1.Val = n1.Val + n2.Val + checkCarry
				checkCarry = 0
				if n1.Val >= 10 {
					n1.Val = n1.Val % 10
					checkCarry = 1
				}
				n2.Val = n1.Val
				n1 = n1.Next
				n2 = n2.Next
			}
		}
	}
	return ans
}

func printNum(numList *ListNode) {
	for n := numList; n != nil; n = n.Next {
		fmt.Println(n.Val)
	}
}
