// link for the task ---> https://leetcode.com/problems/squares-of-a-sorted-array

func sortedSquares(nums []int) []int {
	ans := make([]int, len(nums))
	for i, k, j := 0, len(nums)-1, len(ans)-1; i <= j; k-- {
		fmt.Println(i, k, j)
		if nums[i]*nums[i] > nums[j]*nums[j] {
			ans[k] = nums[i] * nums[i]
			i++
		} else {
			ans[k] = nums[j] * nums[j]
			j--
		}
	}
	return ans
}