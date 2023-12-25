package hw04lrucache

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestList(t *testing.T) {
	t.Run("empty list", func(t *testing.T) {
		l := NewList()

		require.Equal(t, 0, l.Len())
		require.Nil(t, l.Front())
		require.Nil(t, l.Back())
	})

	t.Run("complex", func(t *testing.T) {
		l := NewList()

		l.PushFront(10) // [10]
		l.PushBack(20)  // [10, 20]
		l.PushBack(30)  // [10, 20, 30]
		require.Equal(t, 3, l.Len())

		middle := l.Front().Next // 20
		l.Remove(middle)         // [10, 30]
		require.Equal(t, 2, l.Len())

		for i, v := range [...]int{40, 50, 60, 70, 80} {
			if i%2 == 0 {
				l.PushFront(v)
			} else {
				l.PushBack(v)
			}
		} // [80, 60, 40, 10, 30, 50, 70]

		require.Equal(t, 7, l.Len())
		require.Equal(t, 80, l.Front().Value)
		require.Equal(t, 70, l.Back().Value)

		l.MoveToFront(l.Front()) // [80, 60, 40, 10, 30, 50, 70]
		l.MoveToFront(l.Back())  // [70, 80, 60, 40, 10, 30, 50]

		elems := make([]int, 0, l.Len())
		for i := l.Front(); i != nil; i = i.Next {
			elems = append(elems, i.Value.(int))
		}
		require.Equal(t, []int{70, 80, 60, 40, 10, 30, 50}, elems)
	})

	t.Run("custom student test", func(t *testing.T) {
		l := NewList()

		a := l.PushBack(10)
		b := l.PushFront(9)
		c := l.PushFront(8)
		d := l.PushFront(7)
		e := l.PushFront(6)
		f := l.PushBack(5)

		require.Equal(t, 6, l.Front().Value)

		l.MoveToFront(f)
		require.Equal(t, 5, l.Front().Value)

		l.MoveToFront(a)
		require.Equal(t, 10, l.Front().Value)

		l.MoveToFront(b)
		require.Equal(t, 9, l.Front().Value)

		l.MoveToFront(c)
		require.Equal(t, 8, l.Front().Value)

		l.MoveToFront(d)
		require.Equal(t, 7, l.Front().Value)

		l.MoveToFront(e)
		require.Equal(t, 6, l.Front().Value)

		l.Remove(e)
		require.Equal(t, 7, l.Front().Value)
	})

	t.Run("remove item from with one element", func(t *testing.T) {
		l := NewList()
		i := l.PushFront(10)
		l.Remove(i)

		// nothing must happen
		l.MoveToFront(i)
	})

	t.Run("move one of the elements when two in list", func(t *testing.T) {
		l := NewList()
		a := l.PushFront(10)
		b := l.PushBack(20)

		l.MoveToFront(a)
		require.Equal(t, 10, l.Front().Value)

		l.MoveToFront(b)
		require.Equal(t, 20, l.Front().Value)
	})
}
