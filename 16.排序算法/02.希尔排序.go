package _6_排序算法

//希尔排序也是一种插入排序,是插入排序的改进版本
//希尔排序按增量分组.
//希尔排序不是稳定的排序
//速度: O(n^1.25)
func ShellSort(arr []int) {
	for gap := len(arr)/2; gap>0; gap /= 2 {//增量分组
		for i := gap; i < len(arr); i ++ { //插入排序
			tmp := arr[i]
			j := 0
			for j = i - gap; j >= 0 && arr[j] > tmp; j -= gap {
				arr[j+gap] = arr[j]
			}
			arr[j+gap] = tmp
		}
	}
}
//插入排序
func InsertInShellSort(arr []int, gap, i int){
	tmp := arr[i]
	j := 0
	for j = i-gap;j >= 0 && arr[j] > tmp; j -= gap {
		arr[j+gap] = arr[j]
	}
	arr[j+gap] = tmp
}
