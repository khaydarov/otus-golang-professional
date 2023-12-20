package hw04lrucache

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
	head *ListItem
	tail *ListItem

	length int
}

func (l *list) Len() int {
	return l.length
}

func (l *list) Front() *ListItem {
	return l.head
}

func (l *list) Back() *ListItem {
	return l.tail
}

func (l *list) PushFront(v interface{}) *ListItem {
	newItem := &ListItem{Value: v}

	newItem.Next = l.head.Next
	newItem.Prev = l.head

	l.head.Next.Prev = newItem
	l.head.Next = newItem

	l.length++
	return newItem
}

func (l *list) PushBack(v interface{}) *ListItem {
	newItem := &ListItem{Value: v}

	newItem.Next = l.tail
	newItem.Prev = l.tail.Prev

	l.tail.Prev.Next = newItem
	l.tail.Prev = newItem

	l.length++
	return l.tail
}

func (l *list) Remove(i *ListItem) {
	i.Prev.Next = i.Next
	i.Next.Prev = i.Prev

	l.length--
}

func (l *list) MoveToFront(i *ListItem) {
	i.Prev.Next = i.Next
	i.Next.Prev = i.Prev

	i.Next = l.head.Next
	i.Prev = l.head

	l.head.Next.Prev = i
	l.head.Next = i

}

func NewList() List {
	return new(list)
}
