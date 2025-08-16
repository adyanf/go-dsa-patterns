package mergeintervals_test

import (
	"reflect"
	"testing"

	"github.com/adyanf/go-dsa-patterns/patterns/mergeintervals"
)

func TestLeastInterval(t *testing.T) {
	testCases := []struct {
		name     string
		tasks    []byte
		n        int
		expected int
	}{
		{
			name:     "Case 1",
			tasks:    []byte{'A', 'A', 'B', 'B'},
			n:        2,
			expected: 5,
		},
		{
			name:     "Case 2",
			tasks:    []byte{'A', 'A', 'A', 'B', 'B', 'C', 'C'},
			n:        1,
			expected: 7,
		},
		{
			name:     "Case 3",
			tasks:    []byte{'S', 'I', 'V', 'U', 'W', 'D', 'U', 'X'},
			n:        0,
			expected: 8,
		},
		{
			name:     "Case 4",
			tasks:    []byte{'A', 'K', 'X', 'M', 'W', 'D', 'X', 'B', 'D', 'C', 'O', 'Z', 'D', 'E', 'Q'},
			n:        3,
			expected: 15,
		},
		{
			name:     "Case 5",
			tasks:    []byte{'A', 'B', 'C', 'O', 'Q', 'C', 'Z', 'O', 'X', 'C', 'W', 'Q', 'Z', 'B', 'M', 'N', 'R', 'L', 'C', 'J'},
			n:        10,
			expected: 34,
		},
	}

	for _, tc := range testCases {

		got := mergeintervals.LeastInterval(tc.tasks, tc.n)
		if got != tc.expected {
			t.Errorf("LeastInterval(%v, %v) = %v, expected %v", tc.tasks, tc.n, got, tc.expected)
		}
	}
}

func TestMergeIntervals(t *testing.T) {
	testCases := []struct {
		name     string
		tasks    [][]int
		expected [][]int
	}{
		{
			name:     "Case 1",
			tasks:    [][]int{{4, 6}, {3, 7}, {1, 5}},
			expected: [][]int{{1, 7}},
		},
		{
			name:     "Case 2",
			tasks:    [][]int{{1, 5}, {4, 6}, {11, 15}, {6, 8}},
			expected: [][]int{{1, 8}, {11, 15}},
		},
		{
			name:     "Case 3",
			tasks:    [][]int{{1, 5}},
			expected: [][]int{{1, 5}},
		},
		{
			name:     "Case 4",
			tasks:    [][]int{{1, 9}, {3, 8}, {4, 4}},
			expected: [][]int{{1, 9}},
		},
		{
			name:     "Case 5",
			tasks:    [][]int{{1, 2}, {8, 8}, {3, 4}},
			expected: [][]int{{1, 2}, {3, 4}, {8, 8}},
		},
	}

	for _, tc := range testCases {
		got := mergeintervals.MergeIntervals(tc.tasks)
		if !reflect.DeepEqual(got, tc.expected) {
			t.Errorf("MergeIntervals(%v) = %v, expected %v", tc.tasks, got, tc.expected)
		}
	}
}
