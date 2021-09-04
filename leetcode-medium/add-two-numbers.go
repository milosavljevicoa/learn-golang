package leetcodemedium

import (
	"fmt"
)

//https://leetcode.com/problems/add-two-numbers/
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var start *ListNode = nil
	currList1, currList2 := l1, l2
	curr := start

	for ; currList1 != nil && currList2 != nil; currList1, currList2 = currList1.Next, currList2.Next {
		newNode := &ListNode{Val: currList1.Val + currList2.Val, Next: nil}
		if start == nil {
			start = newNode
		} else {
			curr.Next = newNode
		}
		curr = newNode
	}

	appendListNodes := func(appendTo *ListNode, from *ListNode) {
		for currList1 := from; currList1 != nil; currList1 = currList1.Next {
			newNode := &ListNode{Val: currList1.Val, Next: nil}
			appendTo.Next = newNode
			appendTo = newNode
		}

	}

	if currList1 != nil {
		appendListNodes(curr, currList1)
	} else if currList2 != nil {
		appendListNodes(curr, currList2)
	}

	return start
}

func printList(l1 *ListNode) {
	for curr := l1; curr != nil; curr = curr.Next {
		fmt.Println(curr.Val)
	}
}

func RunAddTwoNumbers() {
	l1, l2 :=
		ListNode{Val: 2,
			Next: &ListNode{Val: 3,
				Next: &ListNode{Val: 4,
					Next: &ListNode{Val: 10,
						Next: &ListNode{Val: 1,
							Next: nil,
						},
					},
				},
			},
		},
		ListNode{Val: 3,
			Next: &ListNode{Val: 4,
				Next: &ListNode{Val: 5,
					Next: nil,
				},
			},
		}

	printList(addTwoNumbers(&l1, &l2))
}
