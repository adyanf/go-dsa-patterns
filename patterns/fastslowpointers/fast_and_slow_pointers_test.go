package fastslowpointers_test

import (
	"testing"

	"github.com/adyanf/go-dsa-patterns/patterns/fastslowpointers"
)

func TestFindDuplicate(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "Case 1",
			nums:     []int{3, 4, 4, 4, 2},
			expected: 4,
		},
		{
			name:     "Case 2",
			nums:     []int{1, 1},
			expected: 1,
		},
		{
			name:     "Case 3",
			nums:     []int{1, 3, 4, 2, 2},
			expected: 2,
		},
		{
			name:     "Case 4",
			nums:     []int{1, 3, 6, 2, 7, 3, 5, 4},
			expected: 3,
		},
		{
			name:     "Case 5",
			nums:     []int{1, 2, 2},
			expected: 2,
		},
	}

	for _, tc := range testCases {
		got := fastslowpointers.FindDuplicate(tc.nums)

		if got != tc.expected {
			t.Errorf("FindDuplicate(%v) = %v, expected %v", tc.nums, got, tc.expected)
		}
	}
}
