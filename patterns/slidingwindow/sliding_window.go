package slidingwindow

/*
The Sliding Window pattern is a computational technique used to efficiently solve problems
that involve finding subarrays or substrings in arrays or strings that satisfy certain conditions.
This pattern is particularly useful when you need to track a subset of data in a larger dataset
and the subset changes as you iterate through the data.

How it works:
The core idea is to maintain a "window" (a contiguous subset of elements) that slides over
the data structure. Instead of recalculating the entire window from scratch for each position,
you incrementally update the window by adding new elements on one side and removing elements
from the other side.

The pattern typically involves:
1. Expanding the window by moving the right pointer and including new elements
2. Contracting the window by moving the left pointer and excluding elements
3. Maintaining window properties (like sum, count, or validity) during expansion/contraction
4. Tracking the optimal window (maximum, minimum, or target) throughout the process

Common Use Cases:
1. Fixed-size sliding window:
   - Finding the maximum sum of any contiguous subarray of size k
   - Finding the average of all contiguous subarrays of size k

2. Dynamic-size sliding window:
   - Finding the longest substring with at most k distinct characters
   - Finding the minimum window substring that contains all characters of a target string
   - Finding the longest substring without repeating characters

3. Two-pointer sliding window:
   - Finding a subarray with a given sum
   - Finding the smallest subarray with sum greater than a target

4. String pattern matching:
   - Finding all anagrams of a pattern in a string
   - Finding permutations of a string within another string

When to use this pattern:
- The problem asks for a contiguous subarray or substring
- The problem involves finding an optimal window (longest, shortest, maximum, minimum)
- You need to track elements in a specific range or window size
- The brute force approach would involve nested loops to check all possible subarrays

The sliding window pattern is efficient because it reduces time complexity from O(n²) or O(n³)
to O(n) in most cases, while maintaining O(1) to O(k) space complexity depending on the problem
requirements.
*/

// FindLongestSubstring finds the length of the longest substring without repeating characters.
func FindLongestSubstring(str string) int {
	// initiate the window start, window end, longest count, and last seen at map for each character
	windowStart, windowEnd, longest := 0, 0, 0
	lastSeenAt := make(map[byte]int)

	// iterate until window end equal to length of str
	for windowEnd < len(str) {
		// if the element in window end already appeared in current window then move the window start to last seen at + 1
		if lastSeenIdx, ok := lastSeenAt[str[windowEnd]]; ok && lastSeenIdx >= windowStart {
			windowStart = lastSeenIdx + 1
		}
		// update the last seen at of the window end character
		lastSeenAt[str[windowEnd]] = windowEnd

		// update longest based on the length of current window
		if windowEnd-windowStart+1 > longest {
			longest = windowEnd - windowStart + 1
		}
		// expand the window
		windowEnd++
	}

	return longest
}
