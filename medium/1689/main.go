//link for the task ---> https://leetcode.com/problems/partitioning-into-minimum-number-of-deci-binary-numbers

func minPartitions(n string) int {
	mx := 0
	for i := 0; i < len(n); i++ {
		if mx < int(n[i]-48) {
			mx = int(n[i] - 48)
		}
	}
	return mx
}