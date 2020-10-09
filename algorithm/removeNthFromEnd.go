package main

import "fmt"

type ListNode struct {
	Value int
	Next  *ListNode
}

/**
题目描述
给定一个链表，删除链表的倒数第n个节点并返回链表的头指针
例如，
 给出的链表为:1->2->3->4->5, n= 2.
 删除了链表的倒数第n个节点之后,链表变为1->2->3->5.
备注：
题目保证n一定是有效的
请给出请给出时间复杂度为\ O(n) O(n)的算法
*/
func main() {
	l := &ListNode{
		Value: 1,
		Next: &ListNode{
			Value: 2,
		},
	}
	head := removeNthFromEnd(l, 2)
	fmt.Printf("%+v", head)
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var pre *ListNode
	slow := head
	fast := head
	for i := n; i > 0; i-- {
		if fast == nil {
			return nil
		}
		fast = fast.Next
	}
	if fast == nil {
		return head.Next
	}
	for fast != nil {
		pre = slow
		slow = slow.Next
		fast = fast.Next
		if fast == nil {
			pre.Next = slow.Next
		}
	}
	return head
}

/**
 *
 * @param pListHead ListNode类
 * @param k int整型
 * @return ListNode类
 */
func FindKthToTail(pListHead *ListNode, k int) *ListNode {
	// write code here
	slow := pListHead
	fast := pListHead
	for i := k; i > 0; i-- {
		if fast == nil {
			return nil
		}
		fast = fast.Next
	}
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}
