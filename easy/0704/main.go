//link for the task ---> https://leetcode.com/problems/binary-search

func search(nums []int, target int) int {
	var l, r, mid int
	r = len(nums) - 1
	for l <= r {
		mid = l + (r-l)/2
		if nums[mid] > target {
			r = mid - 1
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			return mid
		}
	}
	return -1
}