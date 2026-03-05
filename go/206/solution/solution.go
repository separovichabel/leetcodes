package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func invertLinkedList(head *ListNode) *ListNode {
	var prev *ListNode
	cur := head

	for cur != nil {
		next := cur.Next

		cur.Next = prev

		prev = cur
		cur = next
	}

	return prev
}
