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

func TestFindRepeatedDnaSequences(t *testing.T) {
	testCases := []struct {
		name     string
		s        string
		expected map[string]bool
	}{
		{
			name: "Case 1",
			s:    "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT",
			expected: map[string]bool{
				"AAAAACCCCC": true,
				"CCCCCAAAAA": true,
			},
		},
		{
			name: "Case 2",
			s:    "AAAAAAAAAAAAA",
			expected: map[string]bool{
				"AAAAAAAAAA": true,
			},
		},
		{
			name: "Case 3",
			s:    "ACGTACGTACGTACGTACGTACGTACGTACGT",
			expected: map[string]bool{
				"ACGTACGTAC": true,
				"CGTACGTACG": true,
				"GTACGTACGT": true,
				"TACGTACGTA": true,
			},
		},
		{
			name: "Case 4",
			s:    "GGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGG",
			expected: map[string]bool{
				"GGGGGGGGGG": true,
			},
		},
		{
			name:     "Case 5",
			s:        "GTACGTACGTACGCCCCCCCCGGGGG",
			expected: map[string]bool{},
		},
	}

	for _, tc := range testCases {
		got := slidingwindow.FindRepeatedDnaSequences(tc.s)

		for _, value := range got {
			if !tc.expected[value] {
				t.Errorf("FindRepeatedDnaSequences(%v) = %v, expected %v", tc.s, got, tc.expected)
				break
			}
		}
	}
}

func TestFindRepeatedDnaSequencesKMPAlgorithm(t *testing.T) {
	testCases := []struct {
		name     string
		s        string
		expected map[string]bool
	}{
		{
			name: "Case 1",
			s:    "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT",
			expected: map[string]bool{
				"AAAAACCCCC": true,
				"CCCCCAAAAA": true,
			},
		},
		{
			name: "Case 2",
			s:    "AAAAAAAAAAAAA",
			expected: map[string]bool{
				"AAAAAAAAAA": true,
			},
		},
		{
			name: "Case 3",
			s:    "ACGTACGTACGTACGTACGTACGTACGTACGT",
			expected: map[string]bool{
				"ACGTACGTAC": true,
				"CGTACGTACG": true,
				"GTACGTACGT": true,
				"TACGTACGTA": true,
			},
		},
		{
			name: "Case 4",
			s:    "GGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGG",
			expected: map[string]bool{
				"GGGGGGGGGG": true,
			},
		},
		{
			name:     "Case 5",
			s:        "GTACGTACGTACGCCCCCCCCGGGGG",
			expected: map[string]bool{},
		},
	}

	for _, tc := range testCases {
		got := slidingwindow.FindRepeatedDnaSequencesKMPAlgorithm(tc.s)

		for _, value := range got {
			if !tc.expected[value] {
				t.Errorf("FindRepeatedDnaSequencesKMPAlgorithm(%v) = %v, expected %v", tc.s, got, tc.expected)
				break
			}
		}
	}
}
