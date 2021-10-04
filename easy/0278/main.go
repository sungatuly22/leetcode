//link for the task ---> https://leetcode.com/problems/first-bad-version

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
	var l, r, mid, ans int
	l = 1
	r = n
	for l <= r {
		mid = l + (r-l)/2
		if !isBadVersion(mid) {
			l = mid + 1
		} else {
			ans = mid
			r = mid - 1
		}
	}
	return ans
}