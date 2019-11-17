package leet_code

import "fmt"

func TwoSum(nums []int, target int) []int {
	hashMap := make(map[int]int)
	for i := 0;i < len(nums); i ++ {
		hashMap[nums[i]] = i
	}
	fmt.Println(hashMap)
	for i := 0; i < len(nums); i ++ {
		diff := target - nums[i]
		fmt.Printf("diff:%d, i:%d, nums:%d\n", diff, i, nums[i])
		if idx, ok := hashMap[diff]; idx != i && ok {
			return []int{i, idx}
		}
	}
	return []int{0, 0}
}

//思路: 利用相减的结果,错位存储map, 也就是说把相减的结果做key存入map里,待下次相减的结果在map里找到,并返回
func FindTwoSum(list []int, target int) (int, int) {
	checked := make(map[int]int)
	for i := 0; i < len(list); i++ {
		diff := target - list[i]        //相减获取差值
		if _, ok := checked[diff]; ok { //判断是否存在map里
			return checked[diff], i //返回map的值,与当前i为下标
		}
		checked[list[i]] = i //将相减的结果做key存储.
	}
	return -1, -1
}

func TwoSum2(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		diff := target - num
		if j, ok := m[num]; ok {
			return []int{j, i}
		}
		m[diff] = i
	}
	return []int{0, 0}
}
func TwoSum3(nums []int, target int) []int{
	result := make([]int, 0)
	set := make(map[int]int)
	for k, v := range nums {
		set[target - v] = k
	}
	for k, v := range nums {
		if i, ok := set[v];ok {
			result = append(result, k, i)
			break
		}
	}
	return result
}