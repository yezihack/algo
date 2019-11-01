package leet_code_test

import (
	"fmt"
	"sort"
	"testing"
)

func Test3Sum(t *testing.T) {
	nums := []int{-1, 0, 1, 2, -1, -4}
	result := threeSum(nums)
	fmt.Println(result)
}
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	fmt.Println(nums)
	sum := make([][]int, 0)
	for k := 0; k < len(nums); k++ {
		fix := nums[k]
		if nums[k] > 0 {
			break
		}
		if k > 0 && nums[k] == nums[k-1] { //两者=0,再加一个数,永远不会=0的.
			continue
		}
		i, j := k+1, len(nums)-1
		for i < j {
			if nums[i]+nums[j] == -fix {
				if i == k+1 || j == len(nums)-1 {
					sum = append(sum, []int{nums[i], nums[j], nums[k]})
					i++
					j--
				} else if nums[i] == nums[i-1] {
					i++
				} else if nums[j] == nums[j+1] {
					j--
				} else {
					sum = append(sum, []int{nums[i], nums[j], nums[k]})
					i++
					j--
				}
			} else if nums[i]+nums[j] < -fix {
				i++
			} else {
				j--
			}
		}
	}
	return sum
}
