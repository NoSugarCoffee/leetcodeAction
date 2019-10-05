package list

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

//234.回文链表
//请判断一个链表是否为回文链表。
//输入: 1->2
//输出: false
//输入: 1->2->2->1
//输出: true

func isPalindrome(head *ListNode) bool {
	var (
		nodeMap map[int]*ListNode
		todo    *ListNode
		count   int
		index   int
	)
	nodeMap = make(map[int]*ListNode,0)
	for todo = head;todo != nil;todo = todo.Next {
		nodeMap[count] = todo
		count++
	}

	index = count - 1
	for i:=0;i < count / 2;i++ {
		if nodeMap[i].Val != nodeMap[index - i].Val {
			return false
		}
	}
	return true
}


