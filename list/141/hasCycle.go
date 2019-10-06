package list

type ListNode struct {
	Val int
	Next *ListNode
}

//141.环形链表
//输入：head = [3,2,0,-4], pos = 1
//输出：true
//解释：链表中有一个环，其尾部连接到第二个节点
//
//type ListNode struct {
//	Val int
//	Next *ListNode
//}
//
//输入：head = [1,2], pos = 0
//输出：true
//解释：链表中有一个环，其尾部连接到第一个节点。

func hasCycle(head *ListNode) bool {
	m:=make(map[*ListNode]int)
	for todo:=head;todo!=nil;todo=todo.Next {
		if m[todo] != 0 {
			return true
		}else {
			m[todo] = 1
		}
	}
	return false
}
