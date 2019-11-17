package leet_code

import (
	"fmt"
	"sort"
	"strconv"
)
//提交超时,311 / 313 个通过测试用例
func ThreeSum(nums []int) [][]int{
	result := make([][]int, 0)
	repeatMap := make(map[string]struct{})
	for i := 0; i < len(nums) -2; i ++ {
		for j := i + 1; j < len(nums) -1; j ++ {
			for k := j + 1; k < len(nums); k ++ {
				a, b, c := nums[i], nums[j], nums[k]
				abc := strconv.FormatInt(int64(a), 10) + strconv.FormatInt(int64(b), 10) + strconv.FormatInt(int64(c), 10)
				if  _, ok := repeatMap[abc]; a + b + c == 0 && !ok {
					result = append(result, []int{a, b, c})
					repeatMap[abc] = struct{}{}
				}
			}
		}
	}
	return result
}
//执行用时：508 ms, 内存消耗: 22.8 MB, 313 / 313 个通过测试用例
func ThreeSumMap(nums []int) [][]int{
	result := make([][]int, 0)
	repeatMap := make(map[string]struct{})
	for i := 0; i < len(nums) - 2; i ++ {
		if i >= 1 && nums[i] == nums[i-1] {
			continue
		}
		m := make(map[int]struct{})
		for j := i + 1; j <len(nums); j ++ {
			a, b, c := nums[i], nums[j], -nums[i]-nums[j]
			if _, ok := m[b]; !ok {
				m[c] = struct{}{} //这次将c存入map里.
			} else {
				abc := fmt.Sprint(a, b, c)
				if _, ok := repeatMap[abc]; !ok {
					result = append(result, []int{a, c, b})
					repeatMap[abc] = struct{}{}
				}
			}
		}
	}
	return result
}
//执行用时：40 ms, 内存消耗 : 7.9 MB, 313 / 313 个通过测试用例,
func ThreeSumPointer(nums []int) [][]int {
	result := make([][]int, 0)
	sort.Ints(nums)
	if len(nums) < 3 { //小于3个元素,直接返回
		return result
	}
	size := len(nums) //记录长度
	for i := 0; i < size; i ++ { //循环
		if nums[i] > 0 {//第一个元素大于0,后面相加不可能等于0
			return result
		}
		if i > 0 && nums[i] == nums[i - 1] { //两者相等,再加一个数,永远不会=0的.
			continue
		}
		//创建一左一边指针.当三数相加不为0,
		// 分二种情况:
		// 大于0则移动右指针,即--
		// 小于0则移动左指针,即++
		l, r := i + 1, size - 1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			//三数之和等于0,则有解.
			if sum == 0 {
				//加入切片
				result = append(result, []int{nums[i], nums[l], nums[r]})
				for l < r && nums[l] == nums[l+1] { //过滤两数相等情况.
					l ++
				}
				for l < r && nums[r] == nums[r-1] {//过滤两数相等情况.
					r --
				}
				//同时移动左右指针
				l ++
				r --
			} else if sum > 0 {//大于0则移动右指针,即--
				r --
			} else if sum < 0 {// 小于0则移动左指针,即++
				l ++
			}
		}
	}
	return result
}
