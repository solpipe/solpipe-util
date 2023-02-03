package tree

import (
	"encoding/binary"
	"errors"
)

type Tree[T any] struct {
	root *Node[T]
}

func Create[T any](list []T) (*Tree[T], error) {

	t := &Tree[T]{}
	t.grow(list)
	return t, nil
}

type Node[T any] struct {
	parent     *Node[T]
	childLeft  *Node[T]
	childRight *Node[T]
	value      T
}

func (t *Tree[T]) grow(list []T) {
	n := &Node[T]{}
	t.root = n
	n.grow(list)

}

func split(n int) (a int, b int, r int) {
	if n%2 == 0 {
		a = n / 2
		b = n / 2
		r = 0
		return
	} else {
		a = (n + 1) / 2
		b = n - a
		r = a - b
		return
	}
}

func spaceOut[T any](list []T) (left []T, right []T) {
	a, b, r := split(len(list))
	left = list[0:a]
	spacer := list[a:] // should be len(right)=b
	if 0 < r {
		right = make([]T, b+r)
		for i := 0; i < b; i++ {
			right[i] = spacer[i]
		}
		// copy last element in list
		for i := 0; i < r; i++ {
			right[i+b] = list[len(list)-1]
		}

	} else {
		right = list[a:]
	}
	return
}

func (n *Node[T]) grow(list []T) {
	if len(list) == 1 {
		n.value = list[0]
		return
	} else if len(list) == 0 {
		panic("should not be here")
	}
	leftList, rightList := spaceOut(list)

	left := &Node[T]{parent: n}
	left.grow(leftList)
	n.childLeft = left
	right := &Node[T]{parent: n}
	right.grow(rightList)
	n.childRight = right
}

func (t *Tree[T]) Find(b []byte) (ans T, err error) {
	if len(b) < 4 {
		err = errors.New("byte array too short")
		return
	}
	ans = t.find(binary.BigEndian.Uint32(b[0:4])).value
	return
}

func (t *Tree[T]) find(x uint32) *Node[T] {
	n := t.root.find(x)
	if n == nil {
		panic("failed to find node")
	}
	return n
}

func (n *Node[T]) find(x uint32) *Node[T] {
	if n.childLeft == nil {
		return n
	}
	if x%2 == 0 {
		return n.childLeft.find(x / 2)
	} else {
		return n.childRight.find((x - 1) / 2)
	}
}
