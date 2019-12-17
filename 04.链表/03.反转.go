package linked

import "fmt"

//打印链表
func Print(head *SingleNode) {
	p := head.next
	format := ""
	for p != nil {
		format += fmt.Sprint(p.GetData())
		p = p.next
		if p != nil {
			format += "->"
		}
	}
	fmt.Println(format)
}
//反转链表,非递归方法
func Reverse(head *SingleNode) *SingleNode {
	if head == nil || head.next == nil{
		return nil
	}
	//申请一个前驱指针
	var pre *SingleNode
	//设置一个当前指针
	cur := head.next
	for cur != nil { //
		tmp := cur.next //1 定义一个临时指针,指向下一个指针
		cur.next = pre //3 当前的下一个指向前驱指针
		pre = cur //4 前驱指针指向当前
		cur = tmp //2 当前指向下一个指针
	}
	head.next = pre //前驱指向头的下一个结点
	return head
}

