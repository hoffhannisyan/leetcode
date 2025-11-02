package tasks

import (
	"reflect"
	"testing"
)

func sliceToList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0]}
	current := head
	for _, v := range nums[1:] {
		current.Next = &ListNode{Val: v}
		current = current.Next
	}
	return head
}

func listToSlice(node *ListNode) []int {
	result := []int{}
	for node != nil {
		result = append(result, node.Val)
		node = node.Next
	}
	return result
}

func TestAddTwoNumbers(t *testing.T) {
	tests := []struct {
		name     string
		l1       []int
		l2       []int
		expected []int
	}{
		{"Example 1", []int{2, 4, 3}, []int{5, 6, 4}, []int{7, 0, 8}},
		{"Example 2", []int{0}, []int{0}, []int{0}},
		{"Example 3", []int{9, 9, 9, 9, 9, 9, 9}, []int{9, 9, 9, 9}, []int{8, 9, 9, 9, 0, 0, 0, 1}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			l1 := sliceToList(test.l1)
			l2 := sliceToList(test.l2)
			result := AddTwoNumbers(l1, l2)
			got := listToSlice(result)

			if !reflect.DeepEqual(got, test.expected) {
				t.Errorf("AddTwoNumbers(%v, %v) = %v; want %v", test.l1, test.l2, got, test.expected)
			}
		})
	}
}
