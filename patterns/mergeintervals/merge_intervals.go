package mergeintervals

import (
	"sort"
)

/*
The Merge Intervals pattern is all about working with time ranges that might overlap.
Each time range, or interval, is defined by a start and an end point—for example, [10, 20] means the range starts at 10 and ends at 20.
The pattern works by identifying overlapping intervals and merging them into one. If there is no overlap, the interval is kept as it is.

Two intervals [a, b] and [c, d] overlap if:
- a <= c <= b (interval c starts within interval a)
- c <= a <= d (interval a starts within interval c)

Now, let’s see how the Merge Intervals pattern works:
1. Sort the intervals based on their start times (if they aren’t already sorted):
   Sorting ensures that for any two intervals, the start time of the second interval is always greater than or equal to the start time of the first.
   Therefore, we’ll only compare the second interval’s start time with the first interval’s end time to check for overlap.
2. Compare the intervals one by one:
   For each pair, check if the start time of the next interval is less than or equal to the end time of the current interval.
   If it is, the intervals overlap.
3. Merge overlapping intervals: When an overlap is found, merge the two intervals into one and add the merged interval to a new list.
   To merge the intervals, take the smallest start and largest end times from the two intervals.
   For example, if we have the intervals [1, 4] and [3, 6], the merged interval would be [1, 6], as we take the smallest start time (1) and the largest end time (6).
   If there is no overlap, add the current interval to the list as it is.

This pattern is used to solve problems where we need to merge overlapping ranges into one, insert a new range into an existing list,
or find the minimum number of ranges needed to cover everything. For example:
- Scheduling events without conflicts.
- Managing resource usage (like CPU or meeting rooms).
- Finding free or busy time slots.
- Combining overlapping data ranges.

Merging intervals across different fields helps eliminate redundancy, simplify planning, and make the overall system easier to manage.

When to use this pattern:
- The problem involves intervals, ranges, or time periods
- You need to merge overlapping or adjacent intervals
- The problem asks about scheduling, resource allocation, or time management
- You need to find intersections or unions of intervals
- The problem involves detecting conflicts or finding optimal arrangements

The merge intervals pattern typically has O(n log n) time complexity due to sorting,
and O(1) to O(n) space complexity depending on whether you modify in-place or create new arrays.
*/

// LeastTime calculate the least amount of time unit needed to process tasks
// with cooling down period n between identical task.
func LeastInterval(tasks []byte, n int) int {
	// first find the freqyency of each task, because task represented by uppercase letter we can use array of size 26
	taskFreqs := make([]int, 26)
	for _, task := range tasks {
		taskFreqs[task-'A'] += 1
	}

	// sort task frequency in descending order
	sort.Slice(taskFreqs, func(i int, j int) bool {
		return taskFreqs[i] > taskFreqs[j]
	})

	// determine the max frequency from the first index of sorted taskFreqs
	maxFreq := taskFreqs[0]

	// calculate the idle slots from the max frequency - 1 multiplied by cooling down period n
	// substraction by one is because the first task doesn't need cooling down period
	idleSlots := (maxFreq - 1) * n

	// substract the idle slots based on the other tasks that could be inserted to idle slots
	for i := 1; i < 26 && taskFreqs[i] > 0; i++ {
		idleSlots -= min(maxFreq-1, taskFreqs[i])
	}

	// idle slots should not be negative
	if idleSlots < 0 {
		idleSlots = 0
	}

	// the least interval is the number of tasks (one task need one time unit for processing) plus the remaining of idle slots
	return len(tasks) + idleSlots
}
