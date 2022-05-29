package main

func Swap(nums []int, i, j int) {
	if i != j {
		nums[i] = nums[i] ^ nums[j]
		nums[j] = nums[i] ^ nums[j]
		nums[i] = nums[i] ^ nums[j]
	}
}

//荷兰国旗问题
func Partition(nums []int, begin, end int) []int {
	left := begin - 1
	right := end + 1
	index := begin
	target := nums[index]
	for index < right {
		if nums[index] == target {
			index++
		} else if nums[index] > target {
			Swap(nums, index, right-1)
			right--
		} else {
			Swap(nums, index, left+1)
			left++
			index++
		}
	}
	return []int{left + 1, right - 1}
}

//快排
func QuickSort(nums []int, left, right int) {
	if left >= right {
		return
	}
	index := Partition(nums, left, right)
	QuickSort(nums, left, index[0]-1)
	QuickSort(nums, index[0]+1, right)
}

func sortArray(nums []int) []int {
	QuickSort(nums, 0, len(nums)-1)
	return nums

}
