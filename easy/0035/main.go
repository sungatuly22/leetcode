//link for the task ---> https://leetcode.com/problems/search-insert-position

func searchInsert(nums []int, target int) int {
	var l, r, mid, ans int
	r = len(nums) - 1
	ans = -1
	for l <= r {
		mid = l + (r-l)/2
		if nums[mid] < target {
			ans = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return ans + 1
}