package linked

import "fmt"

//单链表只有一个节点, 即下一个节点.
//单链表分头节点, 尾节点.
//单链表特点就是每一个节点有一个指向下一个的节点, 除尾节点为nil.

//定义一个通用类型
type Object interface{}

//创建一个单链表的节点
type SingleLinkNode struct {
	Next  *SingleLinkNode
	Value Object
}

//实例一个节点
func NewSingleLinkNode(v Object) *SingleLinkNode {
	return &SingleLinkNode{
		Value: v,
	}
}

//创建一个单链表结构
//特点: 有头节点, 尾节点, 长度
type SingleLink struct {
	head *SingleLinkNode //头节点
	tail *SingleLinkNode //尾节点
	len  int             // 链表长度
}

//创建一个空链表
//一个单链表具有的功能:
//1. 向头部或尾部添加节点
//2. 删除头部或尾部节点
//3. 获取头部或尾部节点
//4. 返回链表长度.
//5. 打印整个链表
func NewSingleLink() *SingleLink {
	return &SingleLink{}
}

//打印.只要获取头节点, 然后判断下一个节点是否为空, 如果不为空则继续循环读取数据.
func (s *SingleLink) Print() {
	//判断链表是否为空
	if s.len == 0 {
		fmt.Println("链表为空")
		return
	}
	p := s.head
	str := ""
	for p != nil {
		str += fmt.Sprintf("%s", p.Value)
		p = p.Next
		if p != nil {
			str += fmt.Sprintf("=>")
		}
	}
	fmt.Println("链表结果:", str)
}

//1. 向头部添加节点
func (s *SingleLink) AddHead(v Object) {
	//判断当前是否是空链表
	node := NewSingleLinkNode(v)
	if s.len == 0 {
		s.head = node
		s.tail = node
	} else {
		_node := s.head
		s.head = node
		s.head.Next = _node //将新的节点下一个节点指针旧头部节点
	}
	s.len++ //长度相加
}


//1. 向尾部添加节点
func (s *SingleLink) Append(v Object) {
	//判断当前是否是空链表
	node := NewSingleLinkNode(v)
	if s.len == 0 {
		s.head = node
		s.tail = node
	} else {
		oldTail := s.tail   //取出旧尾部节点
		oldTail.Next = node //将旧尾部节点的下一个节点设置成新节点
		s.tail = node       //设置新的尾部节点,即新加入的节点
	}
	s.len++
}

//2. 删除头部节点
func (s *SingleLink) DelHead() Object {
	//判断当前是否是空链表
	if s.len == 0 {
		return ""
	}
	node := s.head //取出当前头部节点
	//需要判断这个是不是最后一个节点
	if node.Next == nil {
		s.head = nil
		s.tail = nil
	} else {
		s.head = node.Next //将下一个节点设置为当前头部.
	}
	s.len--
	return node.Value
}

//2. 删除尾部节点,这个比较麻烦,因为每一个节点没有上一个节点,删除了最后一个节点,你无法再设置新的尾部节点.
//只有一个办法,就是使用循环找到倒数第二个节点
//时间复杂度为O(n).主要原因没有上一个节点的指针,所有只有一个办法就是从头找到尾部.
func (s *SingleLink) DelTail() Object {
	//判断当前是否是空链表
	if s.len == 0 {
		return ""
	}
	node := s.tail //取出当前尾部节点
	//判断整个链表是不是只有一个节点啦
	if s.len == 1 {
		s.head = nil
		s.tail = nil
	} else if s.len == 2 { //判断节点是否只有2个节点
		s.tail = s.head
		s.head.Next = nil
		return node.Value
	} else { //判断节点大于2个节点以上,不含2个元素
		//找到倒数第二个元素
		p := s.head                   //获取头部节点,就可以找到整个链表数据.也就可以找到倒数第二个元素啦
		var _node2 *SingleLinkNode    //定义一个空链表节点
		i := 1                        //从1开始, 因为上面的len也是从1开始, 这样好计算.
		for p != nil && i < s.len-1 { //循环到len-1,即倒数第二个元素的位置.
			p = p.Next //继续向下找
			_node2 = p //赋值给临时变量
			i++
		}
		//找到倒数第二个元素
		s.tail = _node2   //赋值给尾部.
		s.tail.Next = nil //将尾部下一个节点设置为空.
	}
	s.len--
	return node.Value
}
//func main() {
//	link := NewSingleLink()
//	link.Append("a")
//	link.Append("b")
//	link.Append("c")
//	link.Append("d")
//	link.Print()                        //print result
//	fmt.Println("删除头部", link.DelHead()) //删除头部
//	link.Print()                        //print result
//	link.Append("e")
//	link.Append("f")
//	link.Print()                        //print result
//	fmt.Println("删除尾部", link.DelTail()) //删除尾部
//	link.Print()                        //print result
//}
