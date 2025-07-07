package twopointers

import "github.com/adyanf/go-dsa-patterns/structs"

/*
Two Pointers Pattern

The two pointers pattern is a versatile technique used in problem-solving to efficiently
traverse or manipulate sequential data structures, such as arrays, strings, or linked lists.
It involves maintaining two pointers that traverse the data structure in a coordinated manner,
typically starting from different positions or moving in opposite directions.
These pointers dynamically adjust based on specific conditions or criteria, allowing for the efficient
exploration of the data and enabling solutions with optimal time and space complexity.
Whenever there’s a requirement to find two data elements in an array that satisfy a certain condition,
the two pointers pattern should be the first strategy to come to mind.

The implementation depends on the problem, but common variations include:
1. Opposite Ends: One pointer starts at the beginning (left) and the other at the
   end (right). They move towards each other until they meet or cross. This is
   useful for problems involving sorted arrays, like finding a pair with a specific sum.
2. Same Direction (Fast & Slow): Both pointers start at or near the beginning, but
   one (the "fast" pointer) moves ahead of the other (the "slow" pointer). This is
   often used for cycle detection in linked lists, finding the middle element, or
   processing elements in a window.
3. Different Start, Same Direction: Pointers might start at different positions
   (e.g., one at index 0, one at index k) and move in the same direction to compare
   or process elements that are a fixed distance apart.

Use this pattern when all of these conditions are fulfilled:
1. Linear data structure: The input data can be traversed in a linear fashion, such as an array, linked list, or string.
2. Process pairs: Process data elements at two different positions simultaneously.
3. Dynamic pointer movement: Both pointers move independently of each other according to certain conditions or criteria.
   In addition, both pointers might move along the same or two different data structures.

Common use cases include:
- Finding a pair in a sorted array that sums to a target value.
- Reversing an array or string in-place.
- Checking for palindromes.
- Removing duplicates from a sorted array.
- Cycle detection in linked lists (Floyd's Tortoise and Hare algorithm).
- Finding the longest substring without repeating characters.

The primary benefit of the two pointers pattern is its efficiency. It can often
reduce the time complexity of a problem from a brute-force O(n^2) (using nested loops)
down to O(n) (using a single pass with two pointers). It also typically uses O(1)
extra space, making it very memory-efficient.
*/

// RemoveNthLastNode removes the nth last node from a linked list and returns the modified list head.
// The input `head` is the starting node of the linked list, and `n` represents the position from the end to remove.
func RemoveNthLastNode(head *structs.ListNode, n int) *structs.ListNode {
	// Initialize both pointers at the head
	left := head
	right := head

	// Advance the right pointer n nodes ahead
	// This creates a gap of n nodes between left and right
	for range n {
		right = right.Next
	}

	// If right pointer reaches nil, it means we need to remove the head
	// (this happens when n equals the length of the linked list)
	if right == nil {
		return head.Next
	}

	// Move both pointers until right reaches the last node
	// When right reaches the end, left will be at the node just before the one to be removed
	for right.Next != nil {
		right = right.Next
		left = left.Next
	}

	// Remove the nth node from the end by updating the next reference
	// Skip the node that comes after left
	left.Next = left.Next.Next

	// Return the original head as it wasn't removed
	return head
}

// SortColors sorts the array in place so that the elements of the same color are adjacent.
// And the final order is: red (0), then white (1), and then blue (2).
func SortColors(colors []int) []int {
	// mark for the next index to store red (0) element
	redPartitionIndex := 0
	// mark for the next index to store blue (2) element
	bluePartitionIndex := len(colors) - 1

	// example
	// 0, 0, 1 (current, redPartitionIndex), 2, 0, 1 (bluePartitionIndex), 2, 2, 2

	// iterate until current pointer exceeds the bluePartitionIndex pointer
	current := 0
	for current <= bluePartitionIndex {
		switch colors[current] {
		case 0:
			// if the current element is red (0), then swap the element with the red partition index
			// this ensured that red elements is placed at the beginning of the array
			colors[current], colors[redPartitionIndex] = colors[redPartitionIndex], colors[current]
			// move the current and red partition index pointers one position forward
			current++
			redPartitionIndex++
		case 1:
			// if the current element is white (1), then just move the current pointer one position forward
			// the white (1) elements is placed at the middle of the array
			current++
		case 2:
			// if the current element is blue (2), then swap the element with the blue partition index
			// this ensured that blue elements is placed at the end of the array
			colors[current], colors[bluePartitionIndex] = colors[bluePartitionIndex], colors[current]
			// move the blue partition index pointer one position backward
			bluePartitionIndex--
		}
	}

	return colors
}
