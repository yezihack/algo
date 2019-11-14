package _6_排序算法

/*
插入排序:
1. 原始数据越接近有序,速度越快.
2. 时间复杂度:最坏,平均O(n^2),最好O(n)
3. 属于稳定排序
	提高排序效率.
	1. 减少元素的比较次数
	2. 减少元素的移动次数

*/


//插入排序,不使用哨兵方法
//采用升序排序
func InsertSort(arr []int) {
	i, j := 0, 0
	for i = 1; i < len(arr); i ++ {//从第2个位置开始循环
		if arr[i] < arr[i-1] {//判断前面元素与后面元素的大小.如果前大于后则进行交换
			tmp := arr[i] //将后面元素存储临时变量里
			for j = i-1; j >= 0 && tmp < arr[j]; j--{//从i-1位置开始,将元素逐一往后移动
				arr[j+1] = arr[j]
			}
			arr[j+1] = tmp
		}
	}
}
//插入排序,哨兵版本
//减少比较次数
func InsertSortSentry(arr []int) {
	i, j := 0, 0
	for i = 2; i < len(arr); i ++ {
		if arr[i] < arr[i-1] { //只有后面小于前面的数字才进行交换
			arr[0] = arr[i] //设置哨兵
			for j = i-1;arr[0] < arr[j]; j--{//因为有哨兵,无需防止数组越界.
				arr[j+1] = arr[j]//移动元素
			}
			arr[j+1] = arr[0]
		}
	}
}
//折半插入排序
func BInsertSort(arr []int)  {
	for i := 1; i < len(arr); i ++ {
		if arr[i] < arr[i-1] { //无序的元素才进行排序操作
			tmp := arr[i] //设置一个临时变量
			low, high := 0, i-1
			for low <= high {//折半操作
				mid := (low + high) >> 1
				if tmp > arr[mid] {
					low = mid + 1
				} else {
					high = mid - 1
				}
			}
			for j := i-1;j >= high + 1; j -- {//移动元素
				arr[j+1] = arr[j]
			}
			arr[high+1] = tmp//将元素插入到合适的位置
		}
	}
}
//折半哨兵版本插入排序
func BSentryInsertSort(arr []int) {
	for i := 2; i < len(arr); i ++ {
		if arr[i] < arr[i-1] {
			arr[0] = arr[i]//设置哨兵
			//使用折半找到插入的位置
			low, high := 0, i-1
			for low <= high {//折半查找
				mid := (low + high) >> 1
				if arr[0] > arr[mid] {
					low = mid + 1
				} else {
					high = mid - 1
				}
			}
			for j := i-1; j >= high + 1;j-- {//移动大于等于折半查找到的位置+1
				arr[j+1] = arr[j]//移动元素
			}
			arr[high+1] = arr[0]//将哨兵填入折半位置
		}
	}
	arr[0] = 0 //哨兵位置归0
}