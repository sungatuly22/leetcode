// link for the task ---> https://leetcode.com/problems/rotate-array

func rotate(nums []int, k int) {
	rotatedArr := make([]int, len(nums))
	for i, v := range nums {
		rotatedArr[(i+k)%len(nums)] = v
	}
	copy(nums, rotatedArr)
}