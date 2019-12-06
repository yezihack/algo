package greedy

import "strconv"

//问题描述: 给你一个整数 n，使得从 n 中删除 k 个数字之后的数字最大。
//如果 1432219 整数, 删除3个数, 得到4329最大值

//使用贪心算法求得,时间复杂度O(n*m)
func GetMaxNumber(n, count int) int {
	//先将数字转整形数组
	s := strconv.Itoa(int(n))
	b := make([]int, len(s))
	for k, v := range s {//转换成数字数组
		b[k], _ = strconv.Atoi(string(v))
	}
	for i := 0; i < count; i ++ {
		min := b[0]
		k := 0
		for key, v := range b {
			if v < min {//查找最小值
				min = v//最小值
				k = key //记录下标
			}
		}
		if k >= 0 {
			b  = append(b[:k], b[k+1:]...)//最小值从下标中删除掉
		}
	}
	//数字数组转换成整数
	result := ""
	for i := 0; i < len(b); i++ {
		result += strconv.Itoa(b[i])
	}
	v, _ := strconv.Atoi(result)
	return v
}

