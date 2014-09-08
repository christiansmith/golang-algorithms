package main

// Doubly-linked list

type List struct {
	Head *Item
	Tail *Item
}

type Item struct {
	Value string
	Next  *Item
	Prev  *Item
}

func NewList(values ...string) *List {
	list := &List{}

	if len(values) > 0 {
		prev := &Item{values[0], nil, nil}
		list.Head = prev
		list.Tail = prev

		for i := 1; i < len(values); i++ {
			item := &Item{values[i], nil, nil}

			prev.Next = item
			item.Prev = prev
			prev = item

			if i == len(values)-1 {
				list.Tail = item
			}
		}
	}

	return list
}

func (list *List) Prepend(value string) *Item {
	head := list.Head
	item := &Item{value, head, nil}

	list.Head = item

	if head != nil {
		head.Prev = item
	}

	if list.Tail == nil {
		list.Tail = item
	}

	return item
}

func (list *List) Append(value string) *Item {
	tail := list.Tail
	item := &Item{value, nil, tail}

	list.Tail = item

	if tail != nil {
		tail.Next = item
	}

	if list.Head == nil {
		list.Head = item
	}

	return item
}

func (list *List) Shift() *Item {
	head := list.Head

	if head != nil {
		list.Head = head.Next
	}

	if list.Head == nil {
		list.Tail = nil
	}

	return head
}

func (list *List) Insert(before *Item, value string) *Item {
	var item, prev, next *Item

	prev = before.Prev
	next = before
	item = &Item{value, next, prev}

	if prev != nil {
		prev.Next = item
	}

	next.Prev = item

	return item
}

func (list *List) Remove(item *Item) {
	var prev, next, head, tail *Item

	head = list.Head
	tail = list.Tail

	if item != nil {
		if item == head && item == tail {
			list.Head = nil
			list.Tail = nil
		} else if item == head {
			next = head.Next
			next.Prev = nil
			list.Head = head.Next
		} else if item == tail {
			prev = tail.Prev
			prev.Next = nil
			list.Tail = prev
		} else {
			prev = item.Prev
			next = item.Next
			prev.Next = next
			next.Prev = prev
		}
	}

	return
}
