package LeetCode

import (
	"fmt"
	"github.com/yezihack/algo/00.src"
	"testing"
)

//https://leetcode-cn.com/problems/first-missing-positive
// 缺失的第一个正数
//给定一个未排序的整数数组，找出其中没有出现的最小的正整数。
//注: 你的算法的时间复杂度应为O(n)，并且只能使用常数级别的空间
//解题思路: 正整数是从1到无穷大. 所以我们从1开始查找.如果数组的值-1与下标不相等则直接下标+1,就是结果.
//问题来了, 首先我们要对数组进行一次桶排序.这里的桶排序有个要求, 不能使用额外的空间.所以我们只能原地排序.
//方法就是下标与值-1相匹配, 不相等则交换位置.

func MissingFirstNumber(nums []int) int {
	//计算长度.
	size := len(nums)
	if size == 0 {
		return 1
	}
	//对数组进行一次桶排序.而且是原地桶排序
	for i := 0; i < size; i ++ {
		//值必须在size范围内才进行桶排序,否则无意义.
		//然后再判断值与值对应的下标-1是否匹配.不匹配才需要交换.
		for nums[i] > 0 && nums[i] <= size && nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	fmt.Println(nums)
	//通过上面的原地桶排序之后,我们将数组全部看一眼
	// 也就是进行值与下标+1比较一下,如果下标与值不匹配则下标加1则是最小正整数.
	for i := 0; i < size; i ++ {
		if nums[i] - 1 != i {
			return i + 1
		}
	}
	//否则直接返回数组长度+1,即是最小正整数.
	return size + 1
}

func TestMissingFirstNumber(t *testing.T) {
	src.Asset(1, t, 3, MissingFirstNumber([]int{1, 2, 0}))
	src.Asset(2, t, 2, MissingFirstNumber([]int{3,4,-1,1}))
	src.Asset(3, t, 1, MissingFirstNumber([]int{7,8,9,11,12}))
}
//使用map实现
//实现思路: 先将nums全部存入hashmap里, 然后1~n遍历一次,在map里查找有没有此元素.如果没有则是最小正整数
//如果map没有的话,则是长度+1
func MissingByMap(nums []int) int {
	//申请map空间, 其实不符合题目要求. 我们权当练习.
	set := make(map[int]struct{})
	for _, num := range nums {
		set[num] = struct{}{}
	}
	fmt.Println(set)
	//然后1~n遍历一次.查看一下map里是否有此值.如果没有则就是当前的i值.
	for i := 1; i <= len(nums); i ++ {
		if _, ok := set[i]; !ok {
			return i
		}
	}
	fmt.Println(nums)
	return len(nums) + 1
}

func TestMissingByMap(t *testing.T) {
	//src.Asset(1, t, 3, MissingByMap([]int{1, 2, 0}))
	//src.Asset(2, t, 2, MissingByMap([]int{3,4,-1,1}))
	src.Asset(3, t, 1, MissingByMap([]int{7,8,9,11,12}))
}

func MissingNumber(nums []int) int {
	length := len(nums)

	for i := 0; i < length; i ++ {
		fmt.Printf("%d, %d\n", nums[i], nums[i]-1)
		for nums[i] > 0 && nums[i] <= length && nums[nums[i]-1] != nums[i] {
			fmt.Println("swap")
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	fmt.Println("nums", nums)
	//// [1, -1, 3, 4]
	for i := 0; i < length; i ++ {
		if nums[i] != i + 1{
			fmt.Println("return", i, nums[i])
			return i + 1
		}
	}
	return length + 1
}
func TestMissingNumber(t *testing.T) {
	src.Asset(1, t, 2, MissingNumber([]int{1, -1, 3, 4}))
	src.Asset(2, t, 1, MissingNumber([]int{7,8,9,11,12}))
	src.Asset(3, t, 3, MissingNumber([]int{1,2,0}))
	src.Asset(4, t, 2, MissingNumber([]int{3,4,-1,1}))
}