package singlylist

import (
	"fmt"
)

type Node[T comparable] struct {
	val  T
	next *Node[T]
}

type SinglyLinkedList[T comparable] struct {
	head *Node[T]
	len  uint
}

func NewSinglyLinkedList[T comparable]() *SinglyLinkedList[T] {
	return &SinglyLinkedList[T]{
		head: nil,
		len:  0,
	}
}
func (l *SinglyLinkedList[T]) Append(e T) {
	l.InsertAtEnd(e)
}
func (l *SinglyLinkedList[T]) InsertAtHead(e T) {
	head := l.head
	l.head = &Node[T]{e, head}

	l.len++
}
func (l *SinglyLinkedList[T]) InsertAtEnd(e T) {
	endNode := &Node[T]{e, nil}

	if l.len == 0 {
		l.head = endNode
		l.len++
		return
	}
	cur := l.head
	for i := uint(0); i < l.len; i++ {
		if cur.next == nil {
			cur.next = endNode
			l.len++
			return
		}
		cur = cur.next
	}
}

func (l *SinglyLinkedList[T]) InsertAt(v T, i uint) bool {
	if i > l.len {
		return false
	}
	if i == 0 {
		l.InsertAtHead(v)
		return true
	}
	cur := l.head
	for j := uint(0); j < l.len; j++ {
		if i == j {
			cur.next = &Node[T]{v, cur.next}
			break
		}
		cur = cur.next
	}
	l.len++
	return true
}

func (l *SinglyLinkedList[T]) DeleteHead() bool {
	return l.DeleteAt(0)
}

func (l *SinglyLinkedList[T]) DeleteVal(t T) bool {
	var r bool
	cur := l.head
	if l.len == 0 {
		fmt.Println("List is empty")
		return r
	}

	for i := uint(0); i < l.len; i++ {
		if cur.val == t {
			if l.head == cur {
				l.head = cur.next
			} else {
				pre := l.getAt(i - 1)
				pre.next = cur.next
			}
			i--
			l.len--
			r = true
		}

		cur = cur.next
	}
	return r
}

func (l *SinglyLinkedList[T]) DeleteAt(i uint) bool {
	if i > l.len {
		return false
	}
	head := l.head
	if i == 0 {
		l.head = head.next
		head = nil
		l.len--
		return true
	}
	p := l.getAt(i - 1)
	d := p.next
	p.next = d.next
	l.len--
	d = nil
	return true
}

func (l *SinglyLinkedList[T]) Get(index uint) (v T, i bool) {
	var t T
	if index > l.len-1 {
		v, i = t, false
	}
	cur := l.head
	for x := uint(0); x < index; x++ {
		cur = cur.next
	}
	v, i = cur.val, true
	return
}
func (l *SinglyLinkedList[T]) getAt(i uint) *Node[T] {
	if i >= l.len {
		return nil
	}
	if i == 0 {
		return l.head
	}

	cur := l.head
	for j := uint(0); j < l.len; j++ {
		if i == j {
			return cur
		}
		cur = cur.next
	}
	return nil
}

func (l *SinglyLinkedList[T]) Size() uint {
	return l.len
}

func (l *SinglyLinkedList[T]) Print() {
	if l.len == 0 {
		fmt.Println("No nodes in list")
	}
	ptr := l.head
	for i := uint(0); i < l.len; i++ {
		fmt.Printf(" Node value: %v, node add: %v \n", ptr, &ptr)
		ptr = ptr.next
	}
	fmt.Println()
}

func (l *SinglyLinkedList[T]) Traverse(f func(T)) {
	if l.len == 0 {
		fmt.Println("No nodes in list")
	}
	ptr := l.head
	for i := uint(0); i < l.len; i++ {
		f(ptr.val)
		ptr = ptr.next
	}
}
