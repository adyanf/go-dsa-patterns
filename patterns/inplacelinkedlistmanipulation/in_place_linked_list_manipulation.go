package inplacelinkedlistmanipulation

import "github.com/adyanf/go-dsa-patterns/structs"

/*
The in-place manipulation of a linked list pattern allows us to modify a linked list without using any additional memory.
In-place refers to an algorithm that processes or modifies a data structure using only the existing memory space,
without requiring additional memory proportional to the input size. This pattern is best suited for problems
where we need to modify the structure of the linked list, i.e., the order in which nodes are linked together.
For example, some problems require a reversal of a set of nodes in a linked list which can extend to reversing the whole linked list.
Instead of making a new linked list with reversed links, we can do it in place without using additional memory.

The naive approach to reverse a linked list is to traverse it and produce a new linked list with every link reversed.
The time complexity of this algorithm is O(n) while consuming O(n) extra space. How can we implement the in-place reversal of nodes so that no extra space is used?
We iterate over the linked list while keeping track of three nodes: the current node, the next node, and the previous node.
Keeping track of these three nodes enables us to efficiently reverse the links between every pair of nodes.
This in-place reversal of a linked list works in O(n) time and consumes only O(1) space.

Key characteristics:
- Modifies the original linked list structure
- Uses O(1) extra space (excluding input)
- Often involves pointer manipulation and reversal techniques
- Common operations include reversing, rotating, and rearranging nodes

Common techniques:
1. Reverse a linked list or part of it
2. Use two or multiple pointers to traverse and modify
3. Break and reconnect links between nodes
4. Swap nodes or change their positions

Time complexity: Usually O(n) where n is the number of nodes
Space complexity: O(1) - constant extra space

Examples of problems:
- Reverse a linked list
- Reverse nodes in k-groups
- Rotate a linked list
- Remove duplicates from sorted list
- Merge two sorted lists in-place
- Rearrange linked list (odd-even positions)
*/

// SwapPairs swaps every two adjacent nodes of the linked list and return the head
func SwapPairs(head *structs.ListNode) *structs.ListNode {
	// iterate from the head of linked list
	curr := head

	// set the previous node as nil
	var prev *structs.ListNode

	// start iterating until the end or when the remaining node is only one (couldn't be swapped)
	for curr != nil && curr.Next != nil {
		// store the next and nextNext pointer
		next := curr.Next
		nextNext := curr.Next.Next

		// swap the pointer between two adjacent nodes
		next.Next = curr
		curr.Next = nextNext

		// if the previous node exist, change it to point to the next node (because it will be positioned in current node place after swap)
		if prev != nil {
			prev.Next = next
		}
		// if current node is head, then change the head into the next node (because it will be positioned in the head place after swap)
		if curr == head {
			head = next
		}

		// update the previous node and current node for next iteration
		prev = curr
		curr = nextNext
	}

	return head
}
