package LeetCode

import . "github.com/yezihack/algo/00.src"

/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || p == nil || q == nil {
		return nil
	}
	node := root
	for node != nil {
		val := node.Val
		if  p.Val > val && q.Val > val {
			node = node.Right
		} else if p.Val < val && q.Val < val {
			node = node.Left
		} else {
			return node
		}
	}
	return nil
}
func lowestCommonAncestorByReserve(root, p, q *TreeNode) *TreeNode {
	if root == nil || p == nil || q == nil {
		return nil
	}
	node := root
	for node != nil {
		val := node.Val
		if  p.Val > val && q.Val > val {
			return lowestCommonAncestorByReserve(node.Right, p, q)
		} else if p.Val < val && q.Val < val {
			return lowestCommonAncestorByReserve(node.Left, p, q)
		} else {
			return node
		}
	}
	return nil
}
