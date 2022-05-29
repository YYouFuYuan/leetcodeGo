package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	//反转链表，每次指向前一个指针
	var prev *ListNode = nil
	temp := head
	for temp != nil {
		next := temp.Next
		temp.Next = prev
		prev = temp
		temp = next
	}
	return prev
}
