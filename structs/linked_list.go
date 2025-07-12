package structs

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewLinkedListNode(val int, next *ListNode) *ListNode {
	node := new(ListNode)
	node.Val = val
	node.Next = next
	return node
}

func InitLinkedListNode(val int) *ListNode {
	node := new(ListNode)
	node.Val = val
	node.Next = nil
	return node
}

func ReverseLinkedList(head *ListNode) *ListNode {
	var prev, next, curr *ListNode = nil, nil, head

	for curr != nil {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}

type LinkedList struct {
	Head *ListNode
}

// InsertNodeAtHead method will insert a LinkedListNode at head of a linked list.
func (l *LinkedList) InsertNodeAtHead(node *ListNode) {
	if l.Head == nil {
		l.Head = node
	} else {
		node.Next = l.Head
		l.Head = node
	}
}

// CreateLinkedList method will create the linked list using the given integer array with the help of InsertAtHead method.
func (l *LinkedList) CreateLinkedList(lst []int) {
	for i := len(lst) - 1; i >= 0; i-- {
		newNode := InitLinkedListNode(lst[i])
		l.InsertNodeAtHead(newNode)
	}
}

// CreateCycleOnTail method will create a cycle in the linked list from the tail to the given cycle entry node.
func (l *LinkedList) CreateCycleOnTail(cycleEntry int) {
	if cycleEntry == -1 {
		return
	}

	if l.Head == nil {
		return
	}

	var cycleEntryNode *ListNode

	node := l.Head
	for range cycleEntry {
		node = node.Next
	}
	cycleEntryNode = node

	node = l.Head
	for node.Next != nil {
		node = node.Next
	}
	node.Next = cycleEntryNode
}

// DisplayLinkedList method will display the elements of linked list.
func (l *LinkedList) DisplayLinkedList() {
	nodeVisited := make(map[*ListNode]bool)

	curr := l.Head
	fmt.Print("[")
	for curr != nil {
		fmt.Print(curr.Val)
		nodeVisited[curr] = true
		curr = curr.Next
		if nodeVisited[curr] {
			break
		}
		if curr != nil {
			fmt.Print(", ")
		}
	}
	fmt.Print("]")
}

func (l *LinkedList) String() string {
	nodeVisited := make(map[*ListNode]bool)

	str := "["
	curr := l.Head
	for curr != nil {
		str += fmt.Sprintf("%d", curr.Val)
		nodeVisited[curr] = true
		curr = curr.Next
		if nodeVisited[curr] {
			break
		}
		if curr != nil {
			str += ", "
		}
	}
	str += "]"
	return str
}
