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

// DisplayLinkedList method will display the elements of linked list.
func (l *LinkedList) DisplayLinkedList() {
	temp := l.Head
	fmt.Print("[")
	for temp != nil {
		fmt.Print(temp.Val)
		temp = temp.Next
		if temp != nil {
			fmt.Print(", ")
		}
	}
	fmt.Print("]")
}

func (l *LinkedList) String() string {
	str := "["
	temp := l.Head
	for temp != nil {
		str += fmt.Sprintf("%d", temp.Val)
		temp = temp.Next
		if temp != nil {
			str += ", "
		}
	}
	str += "]"
	return str
}
