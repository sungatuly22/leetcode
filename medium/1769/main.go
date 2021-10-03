//link for the task ---> https://leetcode.com/problems/minimum-number-of-operations-to-move-all-balls-to-each-box

func minOperations(boxes string) []int {
	var i, j, k, ans int
	var res []int
	for i = 0; i < len(boxes); i++ {
		ans = 0
		for j = 0; j < i; j++ {
			if boxes[j] == '1' {
				ans += i - j
			}
		}
		for k = i + 1; k < len(boxes); k++ {
			if boxes[k] == '1' {
				ans += k - i
			}
		}
		res = append(res, ans)
	}
	return res
}