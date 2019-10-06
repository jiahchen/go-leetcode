package main

import "fmt"

func main() {
	nums1 := []int{3}
	nums2 := []int{1, 2}
	fmt.Println("Ans: ", findMedianSortedArrays(nums1, nums2))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	preAns := 0.0
	ans := 0.0
	i, j := 0, 0

	if len(nums1) == 0 {
		return findMedian(nums2)
	}

	if len(nums2) == 0 {
		return findMedian(nums1)
	}

	for n := ((len(nums1) + len(nums2)) / 2); n >= 0; n-- {
		preAns = ans

		if i == len(nums1) {
			ans = float64(nums2[j])
			j++
			continue
		}
		if j == len(nums2) {
			ans = float64(nums1[i])
			i++
			continue
		}

		if nums1[i] <= nums2[j] {
			ans = float64(nums1[i])
			i++
		} else {
			ans = float64(nums2[j])
			j++
		}
		fmt.Println(preAns, ans)
	}

	if (len(nums1)+len(nums2))%2 == 1 {
		return ans
	}
	return (preAns + ans) / 2
}

func findMedian(nums []int) float64 {
	length := len(nums)
	if length%2 == 1 {
		return float64(nums[length/2])
	}
	return (float64(nums[length/2]) + float64(nums[length/2-1])) / 2
}
