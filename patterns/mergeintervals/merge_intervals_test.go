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
		name      string
		intervals [][]int
		expected  [][]int
	}{
		{
			name:      "Case 1",
			intervals: [][]int{{4, 6}, {3, 7}, {1, 5}},
			expected:  [][]int{{1, 7}},
		},
		{
			name:      "Case 2",
			intervals: [][]int{{1, 5}, {4, 6}, {11, 15}, {6, 8}},
			expected:  [][]int{{1, 8}, {11, 15}},
		},
		{
			name:      "Case 3",
			intervals: [][]int{{1, 5}},
			expected:  [][]int{{1, 5}},
		},
		{
			name:      "Case 4",
			intervals: [][]int{{1, 9}, {3, 8}, {4, 4}},
			expected:  [][]int{{1, 9}},
		},
		{
			name:      "Case 5",
			intervals: [][]int{{1, 2}, {8, 8}, {3, 4}},
			expected:  [][]int{{1, 2}, {3, 4}, {8, 8}},
		},
	}

	for _, tc := range testCases {
		got := mergeintervals.MergeIntervals(tc.intervals)
		if !reflect.DeepEqual(got, tc.expected) {
			t.Errorf("MergeIntervals(%v) = %v, expected %v", tc.intervals, got, tc.expected)
		}
	}
}

func TestInsertIntervals(t *testing.T) {
	testCases := []struct {
		name              string
		existingIntervals [][]int
		newInterval       []int
		expected          [][]int
	}{
		{
			name:              "Case 1",
			existingIntervals: [][]int{{1, 2}, {3, 4}, {5, 8}, {9, 15}},
			newInterval:       []int{2, 5},
			expected:          [][]int{{1, 8}, {9, 15}},
		},
		{
			name:              "Case 2",
			existingIntervals: [][]int{{1, 6}, {8, 9}, {10, 15}, {16, 18}},
			newInterval:       []int{9, 10},
			expected:          [][]int{{1, 6}, {8, 15}, {16, 18}},
		},
		{
			name:              "Case 3",
			existingIntervals: [][]int{{1, 2}, {3, 4}, {5, 8}, {9, 15}},
			newInterval:       []int{16, 17},
			expected:          [][]int{{1, 2}, {3, 4}, {5, 8}, {9, 15}, {16, 17}},
		},
		{
			name:              "Case 4",
			existingIntervals: [][]int{{1, 4}, {5, 6}, {7, 8}, {9, 10}},
			newInterval:       []int{1, 5},
			expected:          [][]int{{1, 6}, {7, 8}, {9, 10}},
		},
		{
			name:              "Case 5",
			existingIntervals: [][]int{{1, 3}, {4, 6}, {7, 8}, {9, 10}},
			newInterval:       []int{1, 10},
			expected:          [][]int{{1, 10}},
		},
	}

	for _, tc := range testCases {
		got := mergeintervals.InsertInterval(tc.existingIntervals, tc.newInterval)
		if !reflect.DeepEqual(got, tc.expected) {
			t.Errorf("InsertInterval(%v, %v) = %v, expected %v", tc.existingIntervals, tc.newInterval, got, tc.expected)
		}
	}
}
