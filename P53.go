package main

import "math"

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func maxSubArray(nums []int) int {
	result := math.MinInt
	cur := 0
	for _, value := range nums {
		if cur < 0 {
			cur = value
		} else {
			cur += value
		}
		result = max(result, cur)
	}
	return result
}
