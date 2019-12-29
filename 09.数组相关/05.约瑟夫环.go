package main

import (
	"fmt"
)


func main() {
	fmt.Println(Jos(80, 3))
	fmt.Println(Jos2(80,3))
	fmt.Println(Jos3(80, 3))

	fmt.Println(Jos(8, 5))
	fmt.Println(Jos2(8,5))
	fmt.Println(Jos3(8, 5))
}
//非递归方法求解约瑟夫环问题.
//递推方法,推算出公式:
//f(n)=(k+f(n-1))%n;
//f(1)=1;
func Jos(number, index int) int {
	//处理边界
	if number < 0 || index < 0 {
		return -1
	}
	last := 0
	for i := 2; i <= number; i ++ {
		last = (last + index) % i
	}
	return last+1
}
//模拟一个环, 直接剔出第一个人,然后再验证是否符合条件,
// 不符合追加到环尾部.直到环中最后一个元素
func Jos2(number, index int) int {
	//处理边界
	if number < 0 || index < 0 {
		return -1
	}
	//申请一个切片,模拟一个环.一个位置一个人.
	people := make([]int, number)
	for i := 0; i < number; i ++ {
		people[i] = i+1
	}
	i := 0 //编号从0开始
	p := 0
	for len(people) > 1 { //只要切片里的人数还大于1个人则继续数.
		i ++ //正式开始数
		head := people[0] //直接出列第一个人
		people = people[1:] //将从环中删除
		if i % index != 0 { //查看此人是不是中招的人,不是则重新加入环后面.
			people = append(people, head)
		}
		p ++
	}
	fmt.Println("p", p)
	return people[0]
}

//模拟一个环, 如果是目标则对环中的元素标记为0, 直到环中只有一个元素大于0
func Jos3(number, index int) int {
	if number < 0 || index < 0 {
		return -1
	}
	//申请一个切片,模拟一个环.一个位置一个人.
	people := make([]int, number)
	for i := 0; i < number; i ++ {
		people[i] = i+1
	}
	j := 0 //数组指针, 向后走
	l := 0 //记录退出的人数
	k := 0 //报数,1,2,3
	p := 0
	for l < number -1 {
		if people[j] != 0 { //未出局的人就开始数.
			k++
		}
		if k == index {
			people[j] = 0 //出局的人
			k = 0 //重新算
			l++ //记录已经出局1人
		}
		j ++ //继续移动数组指针
		if j == number {//如果移动到数组尾则从头算开始移动
			j = 0
		}
		p ++
	}
	fmt.Println("p", p)
	//大于0的则是目标值
	for i := 0; i < number; i ++ {
		if people[i] > 0 {
			return people[i]
		}
	}
	return -1
}