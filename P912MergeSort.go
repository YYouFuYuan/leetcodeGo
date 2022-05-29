package main

func MergeSort(nums []int, left, right int) {
	if left >= right {
		return
	}
	middle := left + ((right - left) >> 1)
	MergeSort(nums, left, middle)
	MergeSort(nums, middle+1, right)
	Merge(nums, left, middle, right)
}

//合并两个有序数组
func Merge(nums []int, left int, middle int, right int) {
	arr := make([]int, right-left+1)
	index := 0
	a1 := left
	a2 := middle + 1
	for a1 <= middle && a2 <= right {
		if nums[a1] <= nums[a2] {
			arr[index] = nums[a1]
			index++
			a1++
		} else {
			arr[index] = nums[a2]
			index++
			a2++
		}
	}
	for a1 <= middle {
		arr[index] = nums[a1]
		index++
		a1++
	}
	for a2 <= right {
		arr[index] = nums[a2]
		index++
		a2++
	}

	for i := 0; i < right-left+1; i++ {
		nums[left+i] = arr[i]
	}
}
