/*
 * @Author: 百里
 * @Date: 2019-12-03 11:59:33
 * @LastEditTime : 2019-12-28 17:22:24
 * @LastEditors  : 百里
 * @Description: 
 * @FilePath: \yezihack\algo\04.链表\05.回文判断.go
 * @https://github.com/yezihack
 */
package linked

/*******************
题目:　判断字符串是否是回文
如：abba, 从前读到尾和从尾读取前都是一样的.
abcba, aviva 都是回文字

提供三种解题思路:
1.第一种是借用一个栈, 将一半的数据存入栈中, 后半部分与出栈的元素比较.
2.第二种二个指针向中间逼近
3.第三种也是使用二个指针,先从中间开始向左右扩散
********************/
//使用栈判断,思路:将前一半入栈,后一半与栈里的元素比较
func IsPalindromeByLinked(l *SingleLinked) bool {
	size := l.length
	if size == 0 {
		return false
	}
	if size == 1 {
		return true
	}
	//申请一个数组,存储链表前一半部分
	s := make([]string, 0, size/2)
	cur := l.head
	for i := uint(0); i < size; i ++ {
		cur = cur.next
		//如果是奇数,也就是中间的数,跳过
		if size % 2 != 0 && i == size/2 {
			continue
		}
		if i < size/2 {//前一半入栈
			s = append(s, cur.GetData().(string))
		} else {//后一半则与栈进行判断比较
			if s[size-i-1] != cur.GetData().(string) {
				return false
			}
		}
	}
	return true
}
//解题思路:
// 借用两个指针,即左右指针向中间移动比较
func CheckPalindrome(s string) bool {
	if s == "" {
		return false
	}
	if len(s) == 1 {
		return true
	}
	r := []rune(s)//转换成ac
	left, right := 0, len(r)-1
	for left < right {
		if r[left] != r[right] {
			return false
		}
		left ++
		right--
	}
	return true
}

//解题思路:
// 与左右指针向中间逼近的思路刚好相反的思路
// 先从中间开始向两边扩散逐一比较字符
func CheckPalindrome2(s string) bool {
	if s == "" {
		return false
	}
	if len(s) == 1 {
		return true
	}
	r := []rune(s)
	mid := len(s) >> 1
	left, right := 0, 0
	if len(s) % 2 == 0 {//偶数则left=mid-1
		left, right = mid-1, mid
	} else {
		left, right = mid, mid
	}
	//从中间向两边扩散
	for left >= 0 && right < len(r) {
		if r[left] != r[right] {
			return false
		}
		left -- //左边的--
		right++ //右边的++
	}
	return true
}
