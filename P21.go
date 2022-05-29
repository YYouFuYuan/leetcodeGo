package main

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	head := &ListNode{
		Val:  0,
		Next: nil,
	}
	temp := head
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			temp.Next = list1
			temp = temp.Next
			list1 = list1.Next
		} else {
			temp.Next = list2
			temp = temp.Next
			list2 = list2.Next
		}
	}
	if list1 != nil {
		temp.Next = list1
	}
	if list2 != nil {
		temp.Next = list2
	}
	return head.Next
}
