package greedy

import "math"

//贪心算法 对问题的求解,只做出当前的最好的选择

//找零钱 13
func FindGreedy(v int) []int {
	a := []int{10, 8, 5, 3, 2, 1}
	result := make([]int, 0)
	for i := 0; i < len(a); i++ {
		cnt := math.Floor(float64(v) / float64(a[i]))
		if cnt > 0 && v > 0 && v >= a[i]*int(cnt) {
			v -= a[i] * int(cnt)
			for j := 0; j < int(cnt); j++ {
				result = append(result, a[i])
			}
		}
	}
	return result
}
