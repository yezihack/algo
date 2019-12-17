/*
BF算法,即简单匹配算法,也称暴力算法(brute-force)
解题思路:
1.利用两个指针一一比较
2.假如T字符是否匹配S字符
3.匹配步骤: T的索引j, S(模式串)的索引i, 不相等时则j回到0位置,i回到上次位置加1的位置
3.回朔公式:i = i - j + 1
*/
package string_string

//s是主串, t子串
//返回子串出现在主串的位置,如果未匹配到则返回-1
//时间复杂度: O(M*N)
func BFSearch(s, t string) int {
	i, j := 0, 0
	//声明两个指针变量,用于判断字符是否相等.
	//使用一个for进行循环比较.直到一个指针超出某一字符长度.则跳出
	for i < len(s) && j < len(t) {
		//比较
		if s[i] == t[j] {
			i++
			j++
		} else {
			i = i - j + 1 //回溯,回到上一次位置+1的新位置 <重点,必须在j=0之前,否则回不到之前的位置啦,什么为j=0>
			j = 0         //回到原位置
		}
	}
	if j >= len(t) {
		return i - len(t)
	}
	return -1
}

//BF算法Brute Force暴力匹配算法,又叫朴素匹配算法
func BF(s, t string) int {
	n, m := 0, 0 //申请二个指针, n是s主串的指针, m是t模式串的指针
	for n < len(s) && m < len(t) {
		if s[n] == t[m] { //判断指针位置的字符是否相等,如果相等则指针向前移动
			n ++
			m ++
		} else {//否则主串位置回到上次匹配位置的前一位, 子串位置回到0位置,
			n = n - m + 1
			m = 0
		}
	}
	if m >= len(t) {//m指针移动大于等于子串长度,说明匹配成功
		return n-len(t) //返回n指针减去子串t的长度就是匹配子串的开始位置
	}
	return -1 //否则匹配不成功
}

func Bfs(s, t string) int {
	n, m := 0, 0
	for n < len(s) && m < len(t) {
		if s[n] == t[m] {
			n ++
			m ++
		} else {
			n = n - m + 1
			m = 0
		}
	}
	if m >= len(t) {
		return n - len(t)
	}
	return -1
}
func main() {

}
