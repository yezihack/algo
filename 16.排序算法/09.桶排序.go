package _6_排序算法

//思路:申请最大值的数量的桶, 然后把数值存放在桶的下标里.再循环出>0的下标

func BucketSort(arr []int) {
	max := getMax(arr)
	bucket := make([]int, max + 1)
	for i := range arr {
		bucket[arr[i]] ++
	}
	i := 0
	for k, v := range bucket {
		for j := 0; j < v; j ++ {
			arr[i] = k
			i ++
		}
	}
}
func getMax(arr []int) int {
	max := arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	return max
}
