// link for the task ---> https://leetcode.com/problems/find-peak-element

func findPeakElement(nums []int) int {
	var l, r, mid int
	r = len(nums) - 1
	for l < r {
		mid = l + (r-l)/2
		if nums[mid] > nums[mid+1] {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}