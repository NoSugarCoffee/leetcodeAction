package list

import "testing"

func Test_hasCycle(t *testing.T) {
	f3:=&ListNode{Val: 4, Next:nil}
	f2 :=&ListNode{Val: 2, Next:f3}
	f1:=&ListNode{Val: 1, Next:f2}

	var (
		l1,l2,l3,l4 *ListNode
	)
	l4=&ListNode{Val:-4,Next:nil}
	l3=&ListNode{Val:0,Next:l4}
	l2=&ListNode{Val:2,Next:l3}
	l1=&ListNode{Val:3,Next:l2}
	l4.Next=l2
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"nocycle",
		args{f1},
		false,
		},
		{
		"cycle",
		args{l1},
			true,
		},
		{
			"single",
			args{f3},
				false,
		},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasCycle(tt.args.head); got != tt.want {
				t.Errorf("hmasCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
