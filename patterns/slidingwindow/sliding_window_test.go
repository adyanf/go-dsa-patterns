package slidingwindow_test

import (
	"testing"

	"github.com/adyanf/go-dsa-patterns/patterns/slidingwindow"
)

func TestFindLongestSubstring(t *testing.T) {
	testCases := []struct {
		name     string
		str      string
		expected int
	}{
		{
			name:     "Case 1",
			str:      "abcdbea",
			expected: 5,
		},
		{
			name:     "Case 2",
			str:      "aba",
			expected: 2,
		},
		{
			name:     "Case 3",
			str:      "abccabcabcc",
			expected: 3,
		},
		{
			name:     "Case 4",
			str:      "aaaabaaa",
			expected: 2,
		},
		{
			name:     "Case 5",
			str:      "bbbbb",
			expected: 1,
		},
	}

	for _, tc := range testCases {
		got := slidingwindow.FindLongestSubstring(tc.str)

		if got != tc.expected {
			t.Errorf("FindLongestSubstring(%v) = %v, expected %v", tc.str, got, tc.expected)
		}
	}
}

func TestLongestRepeatingCharacterReplacement(t *testing.T) {
	testCases := []struct {
		name     string
		s        string
		k        int
		expected int
	}{
		{
			name:     "Case 1",
			s:        "aaacbbbaabab",
			k:        2,
			expected: 6,
		},
		{
			name:     "Case 2",
			s:        "aaacbbbaabab",
			k:        1,
			expected: 4,
		},
		{
			name:     "Case 3",
			s:        "dippitydip",
			k:        4,
			expected: 6,
		},
		{
			name:     "Case 4",
			s:        "coollooc",
			k:        2,
			expected: 6,
		},
		{
			name:     "Case 5",
			s:        "aaaaaaaaaa",
			k:        2,
			expected: 10,
		},
	}

	for _, tc := range testCases {
		got := slidingwindow.LongestRepeatingCharacterReplacement(tc.s, tc.k)

		if got != tc.expected {
			t.Errorf("LongestRepeatingCharacterReplacement(%v, %v) = %v, expected %v", tc.s, tc.k, got, tc.expected)
		}
	}
}
