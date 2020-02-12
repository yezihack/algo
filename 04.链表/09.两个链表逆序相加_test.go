package linked

import (
	"fmt"
	"strconv"
	"testing"
)
//定义节点结构.
type ListNode struct {
	 Val int 		//值域
	 Next *ListNode //下一个节点指针
 }

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := new(ListNode) //创建一个头节点.(1)
	tail := head  //头节点赋值给尾指针
	carry := 0 //声明一个进位变量. 大于10则 除以10, 得到最高位. 然后取模得到最底位.
	for l1 != nil || l2 != nil { //两个链表不空
		x := 0 //提取链表里的变量
		if l1 != nil {x = l1.Val}
		y := 0 //提取链表里的变量
		if l2 != nil {y = l2.Val}
		sum := x + y + carry //x, y相加,必须还得加上进位变量.
		carry = sum / 10     //除以10, 得到最高位
		tail.Next = &ListNode{Val:sum % 10} //取模得到最底位
		fmt.Printf("sum:%d, carry:%d, mod:%d, x:%d, y:%d\n", sum, carry, sum % 10, x, y)
		tail = tail.Next //进行尾插法.
		if l1 != nil {l1 = l1.Next} //判断链表是否为空
		if l2 != nil {l2 = l2.Next} //判断链表是否为空
	}
	if carry > 0 {//如果进位不为0的话,再创建一个节点
		tail.Next = &ListNode{Val:carry}
	}
	return head.Next //返回下一个节点.因为(1)我们创建了一个空节点.
}
//反转链表
func ReserveLinked(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var pre *ListNode //前驱链表
	for head != nil {
		cur := head.Next
		head.Next = pre
		pre = head
		head = cur
	}
	return pre
}
func TestAddTwoNumbers(t *testing.T) {
	tailInsert := func(arr []int) *ListNode {
		head := new(ListNode)
		tail := head
		for _, v := range arr {
			node := &ListNode{Val:v}
			tail.Next = node
			tail = node
		}
		return head.Next
	}
	//[1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1]
	//[5,6,4]
	s1 := tailInsert([]int{1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1})
	s2 := tailInsert([]int{5,6,4})
	printListLinked(addTwoNumbers(s1, s2))
}
func printListLinked(head *ListNode) {
	str := ""
	for head != nil {
		str += strconv.Itoa(head.Val)
		head = head.Next
		if head != nil {
			str += "=>"
		}
	}
	println(str)
}
//头插法, 意思就是每次插在头部, 已有的元素的前面. 如1,2,3, 头插法结果就是: 3, 2, 1
//头插法的关键在于: 建一个头指针, 循环时,每次创建一个新节点, 将新节点的后续节点指针头指针.以次类推
//func HeadInsert(arr []int) *ListNode {
//	length := len(arr) //计算数组的长度
//	var head *ListNode //创建一个空头指针
//	for i := 0; i < length; i ++ { //循环数组
//		node := &ListNode{Val:arr[i]} //创建一个节点.用于插到头部位置的节点
//		node.Next = head //将之前头指针连接到新的节点尾部. 这样就是头插法的关键.
//		head = node  //将上一步连接好的远整链表赋值给头指针. 待下次循环,继续头插法.
//	}
//	return head //返回头指针.
//}
func HeadInsert(arr []int) *ListNode {
	var head *ListNode
	for i := 0; i < len(arr); i ++ {
		node := &ListNode{Val:arr[i]}
		node.Next = head
		head = node
	}
	return head
}

func TestHeadInsert(t *testing.T){
	//创建一组数列的链表.
	printListLinked(HeadInsert([]int{1,2,3}))
}
//func TailInsert(arr []int) *ListNode {
//	head := new(ListNode) //初始一个空指针节点,即头指针节点.
//	end := head //再创建一个尾指针
//	for i := 0; i < len(arr); i ++  { //遍历数组
//		node := &ListNode{Val:arr[i]} //创建一个新的节点.
//		end.Next = node //把新的节点赋值给尾指针的next指针.实现尾指针操作.
//		end = node      //新的节点指向尾指针.让下一次操作,继续指向尾部.
//	}
//	return head.Next //返回指针的next. 因为初始指针时,是一个空节点.不返回.
//}
func TailInsert(arr []int) *ListNode {
	head := new(ListNode)
	tail := head
	for i := 0; i < len(arr);i ++ {
		node := &ListNode{Val:arr[i]}
		tail.Next = node
		tail = node
	}
	return head.Next
}

func TestTailInsert(t *testing.T){
	printListLinked(TailInsert([]int{1,2,3,4,5}))
}

func TestNumber(t *testing.T) {
	num := 123
	for num > 0 {
		n := num % 10
		num = num / 10
		fmt.Printf("n:%d, num:%d\n", n, num)
	}
}