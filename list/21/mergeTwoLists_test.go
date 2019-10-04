package list

import (
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	f3:=&ListNode{Val: 4, Next:nil}
	f2 :=&ListNode{Val: 2, Next:f3}
	f1:=&ListNode{Val: 1, Next:f2}

	l3:=&ListNode{Val: 4, Next:nil}
	l2 :=&ListNode{Val: 3, Next:l3}
	l1:=&ListNode{Val: 1, Next:l2}
	list := mergeTwoLists(f1,l1)

	for ;list!=nil;list= Next {
		t.Log(list)
	}

	list = mergeTwoLists(nil,nil)
	for ;list!=nil;list= Next {
		t.Log(Val)
	}

	list = mergeTwoLists(f1,nil)
	for ;list!=nil;list= Next {
		t.Log(Val)
	}
}
