package solution

import "testing"

func buildLinkedListWithCycle(values []int, pos int) (*ListNode, *ListNode) {
	if len(values) == 0 {
		return nil, nil
	}

	nodes := make([]*ListNode, len(values))
	for i, v := range values {
		nodes[i] = &ListNode{Val: v}
		if i > 0 {
			nodes[i-1].Next = nodes[i]
		}
	}

	var cycleStart *ListNode
	if pos >= 0 && pos < len(nodes) {
		cycleStart = nodes[pos]
		nodes[len(nodes)-1].Next = cycleStart
	}

	return nodes[0], cycleStart
}

func TestDetectCycle(t *testing.T) {
	tests := []struct {
		name   string
		values []int
		pos    int
	}{
		{
			name:   "example 1 - cycle at index 1",
			values: []int{3, 2, 0, -4},
			pos:    1,
		},
		{
			name:   "example 2 - cycle at index 0",
			values: []int{1, 2},
			pos:    0,
		},
		{
			name:   "example 3 - no cycle",
			values: []int{1},
			pos:    -1,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			head, expectedNode := buildLinkedListWithCycle(tc.values, tc.pos)
			got := detectCycle(head)

			if got != expectedNode {
				t.Fatalf("detectCycle() returned %p, expected %p", got, expectedNode)
			}
		})
	}
}
