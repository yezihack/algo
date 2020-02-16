/*
 * @Author: 百里
 * @Date: 2020-02-11 16:18:27
 * @LastEditTime: 2020-02-11 16:18:39
 * @LastEditors: 百里
 * @Description:
 * @FilePath: \leetcode\02.链表\02.头插法.go
 * @https://github.com/yezihack
 */
package linked
//定义一个链表结构
type linked struct {
	data int		//存储的值
	next *linked	//链表下一个节点
}
//头插法. 创建一个头插法链表
func HeadInsertLinked(arr []int) *linked {
	var head *linked			//声明一个空链表
	for i := 0; i < len(arr); i ++ {
		node := &linked{data:arr[i]} //创建一个节点
		node.next = head			//新节点的下一个节点存储是头节点,实现头插法的关键代码行.
		head = node					//设置新的头节点.实现下一步头插法.
	}
	return head //返回头节点 .
}
