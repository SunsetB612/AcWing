package main

func searchRange(nums []int, target int) []int {
	result := []int{lowerBound(nums, target), upperBound(nums, target)}
	return result
}

func lowerBound(nums []int, target int) int {
	result := -1
	low, high := 0, len(nums)-1

	for low <= high {
		mid := (low + high) / 2
		if nums[mid] == target {
			result = mid
			high = mid - 1
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return result
}

func upperBound(nums []int, target int) int {
	result := -1
	low, high := 0, len(nums)-1

	for low <= high {
		mid := (low + high) / 2
		if nums[mid] == target {
			result = mid
			low = mid + 1
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return result
}
func main() {

}
