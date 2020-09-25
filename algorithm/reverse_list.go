package main

import "fmt"

type ListNode struct {
	Value int
	Next  *ListNode
}

//翻转链表
func main() {
	pHead := &ListNode{
		Value: 1,
		Next:  nil,
	}
	head := ReverseList(pHead)
	fmt.Println(head)
}
func ReverseList(pHead *ListNode) *ListNode {
	var pre *ListNode = nil
	var cur = pHead
	var nex *ListNode = nil
	for cur != nil {
		nex = cur.Next
		cur.Next = pre
		pre = cur
		cur = nex
	}
	return pre
}
