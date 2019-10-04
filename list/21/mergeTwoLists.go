package list

type ListNode struct {
	Val int
	Next *ListNode
}

// 21.合并两个有序链表
// 将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
//输入：1->2->4, 1->3->4
//输出：1->1->2->3->4->4

// 使用改变当前结点的方式
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	//暂时没法与下列合并，做特殊判断
	if l1 == nil && l2 ==nil {
		return nil
	}
	//创建新的结点
	var (
		// 创建的新结点
		newNode *ListNode
		// 跟踪新链表的第一个结点
		head *ListNode
		// 改名
		l1Todo *ListNode = l1
		l2Todo *ListNode = l2
	)
	newNode = new(ListNode)
	head = newNode

	for  ;l1Todo != nil && l2Todo != nil;{
		if l1Todo.Val <= l2Todo.Val {
			// 值 copy
			*newNode = *l1Todo
			l1Todo = l1Todo.Next
		}else {
			*newNode =*l2Todo
			l2Todo = l2Todo.Next
		}
		// 开辟新结点
		newNode.Next = new(ListNode)
		newNode = newNode.Next
	}

	// 尾巴
	if (l1Todo != nil){
		*newNode = *l1Todo
	}
	if (l2Todo != nil){
		*newNode = *l2Todo
	}
	return head
}


