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
	// use dummy node to keep track of head changes
	dummyNode := &structs.ListNode{Val: -1, Next: head}

	// set the head as the first iteration and dummy node as the previous node of head
	currNode := head
	prevNode := dummyNode

	// start iterating until the end or when the remaining node is only one (couldn't be swapped)
	for currNode != nil && currNode.Next != nil {
		// set the first and second node to be swapped
		firstNode := currNode
		secondNode := currNode.Next

		// swap the pointer between two adjacent nodes and set the previous node next as the second node
		// after the swap the second node will be the next of previous node
		prevNode.Next = secondNode
		firstNode.Next = secondNode.Next
		secondNode.Next = firstNode

		// update the previous node and current node for next iteration
		prevNode = firstNode
		currNode = firstNode.Next
	}

	return dummyNode.Next
}
