package _3_动态规划

import "fmt"

type Goods struct {
	name string //名称
	weight int //重量
	value int //价值
}
//包背问题.
//包背限定m大小, 有不同的物品列表.需要装载最大价值的物品列表.
func GetMaxValueGoodsList(list []*Goods, w int) {
	//初始矩阵大小
	rowLen := len(list)+1
	colLen := w
	matrix := make([][]*Goods, rowLen)
	for row := 0; row < rowLen; row ++ {
		matrix[row] = make([]*Goods, colLen)
	}
	for i := 0; i < rowLen; i ++ {
		for j := 0; j < colLen; j ++ {
			matrix[i][j] = new(Goods)
		}
	}
	for i := 0; i < len(list); i ++ {
		matrix[i][0].name = list[i].name
	}
	//设置一个哨兵,防止溢出,减少判断次数
	printGoods(list)
	printMatrix2(w, matrix)
	return
	//对矩阵进行赋值
	for row := 1; row <= len(list); row ++ {
		for c := 1; c <= w; c ++ {
			if list[row].weight >= c { //背包是否能装载得下.如果能,则涉及到两个问题.1偷,2不偷
				//1.如果不偷的话,价值即是上一次的坐, 即上一个单元格的值
				noSteal := matrix[row-1][c].value
				//2.如果偷的话,价值就是自身, 然后再将总背包容量减去当前物品占用的空间大小
				steal := list[row].value//当前物品的价值
				currWeight := list[row].weight //当前物品的重量
				surWeight := c - currWeight //计算剩余背包容量的大小,看看还能装着啥
				if surWeight > 0 {//剩余容量大小
					//查看之前容量大小的最大价值
					steal += matrix[row-1][surWeight].value
				}
				if noSteal > steal {
					matrix[row][c].value = noSteal
				} else {
					matrix[row][c].value = steal
				}
			}
		}
	}
	printMatrix2(w, matrix)
}
func printGoods(list []*Goods)  {
	for k, item := range list {
		fmt.Printf("ID:%d,name:%s, weight:%d, value:%d\n", k, item.name, item.weight, item.value)
	}
}
func printMatrix2(weight int, list [][]*Goods)  {
	fmt.Print("   ")
	for i := 1; i <= weight; i ++ {
		fmt.Print(i, " ")
	}
	fmt.Println()
	for k, rows := range list {
		for kk, item := range rows {
			if k == 0 {
				continue
			}
			if kk == 0{
				fmt.Print(item.name)
			}
			fmt.Print(item.value, " ")
		}
		if k > 0 {
			fmt.Println()
		}
	}
	fmt.Println()
}