package _8_哨兵的应用

//不利用哨兵查找方法
func Find(arr []int, val int) int {
	if len(arr) == 0 {
		return -1
	}
	i := 0
	for i < len(arr) {
		if arr[i] == val {
			return i
		}
		i ++
	}
	return -1
}
//借用哨兵实现查询,减少一次比较
func FindBySentry(arr []int, val int) int {
	size := len(arr)
	if size == 0 {
		return -1
	}
	//先判断数组最后一位是否与查询值相等
	if arr[size-1] == val {
		return size -1
	}
	//将最后一位保存在临时变量里,方便后面恢复
	temp := arr[size - 1]
	//将查询的值做为哨兵放在数组最后一位.
	arr[size-1] = val
	i := 0
	//进行比较
	for arr[i] != val {
		i ++
	}
	//恢复值
	arr[size-1] = temp
	//如果i==最后一位的下标,说明没有找到
	if i == size-1 {
		return -1
	} else {
		return i
	}
}
