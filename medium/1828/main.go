//link for the task ---> https://leetcode.com/problems/queries-on-number-of-points-inside-a-circle

func countPoints(points [][]int, queries [][]int) []int {
	var dist = 0.0
	ans := 0
	var res []int
	for _, query := range queries {
		ans = 0
		for _, point := range points {
			dist = math.Sqrt(float64((query[0]-point[0])*(query[0]-point[0]) + (query[1]-point[1])*(query[1]-point[1])))
			if dist <= float64(query[2]) {
				ans++
			}
		}
		res = append(res, ans)
	}
	return res
}