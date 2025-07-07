package twopointers_test

import (
	"reflect"
	"testing"

	"github.com/adyanf/go-dsa-patterns/patterns/twopointers"
	"github.com/adyanf/go-dsa-patterns/structs"
)

func TestRemoveNthLastNode(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		delete   int
		expected []int
	}{
		{
			name:     "Case 1",
			nums:     []int{23, 28, 10, 5, 67, 39, 70, 28},
			delete:   2,
			expected: []int{23, 28, 10, 5, 67, 39, 28},
		},
		{
			name:     "Case 2",
			nums:     []int{34, 53, 6, 95, 38, 28, 17, 63, 16, 76},
			delete:   3,
			expected: []int{34, 53, 6, 95, 38, 28, 17, 16, 76},
		},
		{
			name:     "Case 3",
			nums:     []int{288, 224, 275, 390, 4, 383, 330, 60, 193},
			delete:   4,
			expected: []int{288, 224, 275, 390, 4, 330, 60, 193},
		},
		{
			name:     "Case 4",
			nums:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			delete:   1,
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8},
		},
		{
			name:     "Case 5",
			nums:     []int{69, 8, 49, 106, 116, 112},
			delete:   6,
			expected: []int{8, 49, 106, 116, 112},
		},
	}

	for _, tc := range testCases {
		ll := &structs.LinkedList{}
		ll.CreateLinkedList(tc.nums)

		got := twopointers.RemoveNthLastNode(ll.Head, tc.delete)

		llResult := &structs.LinkedList{}
		llResult.InsertNodeAtHead(got)

		llExpected := &structs.LinkedList{}
		llExpected.CreateLinkedList(tc.expected)

		if llResult.String() != llExpected.String() {
			t.Errorf("RemoveNthLastNode(%v, %v) = %v, expected %v", tc.nums, tc.delete, llResult.String(), llExpected.String())
		}
	}
}

func TestSortColors(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		expected []int
	}{
		{
			name:     "Case 1",
			nums:     []int{0, 1, 0},
			expected: []int{0, 0, 1},
		},
		{
			name:     "Case 2",
			nums:     []int{1},
			expected: []int{1},
		},
		{
			name:     "Case 3",
			nums:     []int{2, 2},
			expected: []int{2, 2},
		},
		{
			name:     "Case 4",
			nums:     []int{1, 1, 0, 2},
			expected: []int{0, 1, 1, 2},
		},
		{
			name:     "Case 5",
			nums:     []int{2, 1, 1, 0, 0},
			expected: []int{0, 0, 1, 1, 2},
		},
	}

	for _, tc := range testCases {
		twopointers.SortColors(tc.nums)

		if !reflect.DeepEqual(tc.nums, tc.expected) {
			t.Errorf("SortColors(%v) = %v, expected %v", tc.nums, tc.nums, tc.expected)
		}
	}
}
