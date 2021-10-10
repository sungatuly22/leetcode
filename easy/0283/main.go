// link for the task ---> https://leetcode.com/problems/move-zeroes

func moveZeroes(nums []int) {
	ans := make([]int, len(nums))
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			ans[j] = nums[i]
			j++
		}
	}
	for i := j; i < len(nums); i++ {
		ans[i] = 0
	}
	copy(nums, ans)
}