package solution

import (
	"reflect"
	"testing"
)

func makeList(values []int) *ListNode {
	var head *ListNode

	for i := len(values) - 1; i >= 0; i-- {
		cur := &ListNode{Val: values[i], Next: head}
		head = cur
	}

	return head
}

func listToSlice(head *ListNode) []int {
	if head == nil {
		return nil
	}
	resp := []int{}
	for head != nil {
		resp = append(resp, head.Val)
		head = head.Next
	}

	return resp
}

func TestInvertLinkedList(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "empty list",
			input:    nil,
			expected: nil,
		},
		{
			name:     "single node",
			input:    []int{1},
			expected: []int{1},
		},
		{
			name:     "multiple nodes",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{5, 4, 3, 2, 1},
		},
		{
			name:     "with duplicates",
			input:    []int{1, 2, 2, 3},
			expected: []int{3, 2, 2, 1},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			head := makeList(tc.input)
			result := invertLinkedList(head)
			resultSlice := listToSlice(result)

			if !reflect.DeepEqual(resultSlice, tc.expected) {
				t.Fatalf("expected %v, got %v", tc.expected, resultSlice)
			}
		})
	}
}
