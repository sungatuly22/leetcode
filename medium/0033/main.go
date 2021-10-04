//link for the task ---> https://leetcode.com/problems/search-in-rotated-sorted-array

func search(nums []int, target int) int {
	var l, r, mid int
	r = len(nums) - 1
	for l <= r {
		mid = l + (r-l)/2
		if nums[mid] == target {
			return mid
		}
		if nums[l] <= nums[mid] {
			if nums[mid] < target || nums[l] > target {
				l = mid + 1
			} else {
				r = mid - 1
			}
		} else {
			if nums[r] < target || nums[mid] > target {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
	}
	return -1
}