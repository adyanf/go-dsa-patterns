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

func TestCircularArrayLoop(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		expected bool
	}{
		{
			name:     "Case 1",
			nums:     []int{1, 3, -2, -4, 1},
			expected: true,
		},
		{
			name:     "Case 2",
			nums:     []int{2, 1, -1, -2},
			expected: false,
		},
		{
			name:     "Case 3",
			nums:     []int{5, 4, -2, -1, 3},
			expected: false,
		},
		{
			name:     "Case 4",
			nums:     []int{1, 2, -3, 3, 4, 7, 1},
			expected: true,
		},
		{
			name:     "Case 5",
			nums:     []int{3, 3, 1, -1, 2},
			expected: true,
		},
		{
			name:     "Case 6",
			nums:     []int{-375, -2743, -2, -1566, -3648, -3555, -1299, -272, -1179, -1026, -2964, -4658, -4167, -3256, -1729, -1037, -3402, -4794, -620, -376, -2077, -1190, -458, -3497, -992, -720, -1495, -3439, -2710, -4605, -3824, -1221, -828, -3523, -2944, -519, -4979, -3997, -4696},
			expected: false,
		},
	}

	for _, tc := range testCases {
		got := fastslowpointers.CircularArrayLoop(tc.nums)

		if got != tc.expected {
			t.Errorf("CircularArrayLoop(%v) = %v, expected %v", tc.nums, got, tc.expected)
		}
	}
}
