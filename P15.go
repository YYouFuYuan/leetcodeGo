package main

import "sort"

func threeSum(nums []int) [][]int {
	var result [][]int //定义一个动态的slice
	sort.Ints(nums)    //从小到大排序
	for index, value := range nums {
		if value > 0 {
			break
		}
		if index > 0 && nums[index] == nums[index-1] {
			continue
		}
		left := index + 1
		right := len(nums) - 1
		//双指针
		for left < right {
			sum := value + nums[left] + nums[right]
			if sum == 0 {
				result = append(result, []int{nums[index], nums[left], nums[right]})
				for left+1 < right && nums[left] == nums[left+1] {
					left++
				}
				for right-1 > left && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if sum > 0 {
				right--
			} else {
				left++
			}
		}
	}
	return result
}
