package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

// There is a better solution with O(1) in memory usage
func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	seen := map[*ListNode]bool{}

	for head.Next != nil {
		seen[head] = true

		_, ok := seen[head.Next]

		if ok {
			return head.Next
		}

		head = head.Next
	}

	return nil
}
