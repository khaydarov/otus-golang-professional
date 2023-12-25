package hw04lrucache

import "sync"

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	front *ListItem
	back  *ListItem

	length int
	mu     sync.Mutex
}

func (l *list) Len() int {
	l.mu.Lock()
	defer l.mu.Unlock()

	return l.length
}

func (l *list) Front() *ListItem {
	l.mu.Lock()
	defer l.mu.Unlock()

	return l.front
}

func (l *list) Back() *ListItem {
	l.mu.Lock()
	defer l.mu.Unlock()

	return l.back
}

func (l *list) PushFront(v interface{}) *ListItem {
	l.mu.Lock()
	defer l.mu.Unlock()

	newItem := &ListItem{Value: v}

	// list is empty -> front and back points to the new item
	// list is not empty -> new item points to the current front and becomes front
	if l.front == nil && l.back == nil {
		l.front = newItem
		l.back = newItem
	} else {
		newItem.Next = l.front

		l.front.Prev = newItem
		l.front = newItem
	}

	l.length++
	return newItem
}

func (l *list) PushBack(v interface{}) *ListItem {
	l.mu.Lock()
	defer l.mu.Unlock()

	newItem := &ListItem{Value: v}

	// list is empty -> front and back points to the new item
	// list is not empty -> new item points to the current back and becomes back
	if l.front == nil && l.back == nil {
		l.front = newItem
		l.back = newItem
	} else {
		newItem.Prev = l.back

		l.back.Next = newItem
		l.back = newItem
	}

	l.length++
	return newItem
}

func (l *list) Remove(i *ListItem) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.front == l.back {
		l.front = nil
		l.back = nil

		return
	}

	switch i {
	case l.front:
		l.front = l.front.Next
		l.front.Prev = nil
	case l.back:
		l.back = l.back.Prev
		l.back.Next = nil
	default:
		i.Prev.Next = i.Next
		i.Next.Prev = i.Prev
	}

	l.length--
}

func (l *list) MoveToFront(i *ListItem) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.front == l.back {
		return
	}

	if l.back == i {
		l.back = l.back.Prev

		i.Prev.Next = nil
		i.Prev = nil
		i.Next = l.front

		l.front.Prev = i
		l.front = i
	} else if i != l.front {
		i.Prev.Next = i.Next
		i.Next.Prev = i.Prev

		l.front.Prev = i
		l.front = i
	}
}

func NewList() List {
	return new(list)
}
