package main

import "testing"

func TestNewListEmpty(t *testing.T) {
	list := NewList()
	if list.Head != nil {
		t.Error(`Expected head to be nil, got`, list.Head)
	}

	if list.Tail != nil {
		t.Error(`Expected tail to be nil, got`, list.Tail)
	}
}

func TestNewList(t *testing.T) {
	list := NewList("Chapter Two")
	head := list.Head
	tail := list.Tail

	if head.Value != "Chapter Two" {
		t.Error("Expected \"Chapter Two\", got", head.Value)
	}

	if head.Next != nil {
		t.Error(`Expected next to be nil`)
	}

	if tail != head {
		t.Error(`Expected head to equal tail`)
	}
}

func BenchmarkNewList(b *testing.B) {
	for n := 0; n < b.N; n++ {
		NewList("First Item")
	}
}

func TestPrependEmptyList(t *testing.T) {
	list := NewList()
	zero := list.Prepend("Zero")

	if list.Head != zero {
		t.Error("Expected zero to be the head of the list, got", list.Head)
	}

	if list.Tail != zero {
		t.Error("Expected zero to be the tail of the list, got", list.Tail)
	}
}

func TestPrependList(t *testing.T) {
	list := NewList("One", "Two", "Three")
	zero := list.Prepend("Zero")

	if list.Head != zero {
		t.Error("Expected zero to be the head of the list, got", list.Head)
	}
}

func BenchmarkPrependList(b *testing.B) {
	list := NewList()
	for n := 0; n < b.N; n++ {
		list.Prepend("Item" + string(n))
	}
}

func TestAppendEmptyList(t *testing.T) {
	list := NewList()
	last := list.Append("last")

	if list.Head != last {
		t.Error("Expected zero to be the head of the list, got", list.Head)
	}

	if list.Tail != last {
		t.Error("Expected zero to be the tail of the list, got", list.Tail)
	}
}

func TestAppendList(t *testing.T) {
	list := NewList("One", "Two")
	three := list.Append("Three")

	if list.Tail != three {
		t.Error("Expected three to be the tail of the list, got", list.Tail)
	}
}

func BenchmarkAppendList(b *testing.B) {
	list := NewList()
	for n := 0; n < b.N; n++ {
		list.Append("Item" + string(n))
	}
}

func TestShiftLastFromList(t *testing.T) {
	list := NewList("Last")
	list.Shift()

	if list.Head != nil {
		t.Error(`Expected head to be nil, got`, list.Head)
	}

	if list.Tail != nil {
		t.Error(`Expected tail to be nil, got`, list.Tail)
	}
}

func TestShiftEmptyList(t *testing.T) {
	list := NewList()
	empty := list.Shift()

	if empty != nil {
		t.Error(`Expected empty to be nil, got`, nil)
	}
}

func TestShiftList(t *testing.T) {
	list := NewList("One", "Two", "Three")
	next := list.Shift()

	if next.Value != "One" {
		t.Error("Expected next to be \"One\", got", next)
	}

	if list.Head.Value != "Two" {
		t.Error(`Expected head to be "Two", got`, list.Head)
	}
}

func BenchmarkShiftList(b *testing.B) {
	list := BigList(b.N)
	//b.ResetTimer()

	for n := 0; n < b.N; n++ {
		list.Shift()
	}
}

func TestInsertList(t *testing.T) {
	list := NewList("One", "Two", "Four")
	three := list.Insert(list.Tail, "Three")

	if list.Head.Next != three.Prev {
		t.Error(`Expected three to be linked to two`)
	}

	if list.Tail != three.Next {
		t.Error(`Expected three to be linked to four`)
	}

	if three.Value != "Three" {
		t.Error(`Expected the value of item to be "Three", got`, three.Value)
	}
}

func BenchmarkInsertList(b *testing.B) {
	list := NewList("Last")
	for n := 0; n < b.N; n++ {
		list.Insert(list.Tail, "Item "+string(n))
	}
}

func TestRemoveFromList(t *testing.T) {
	list := NewList("One", "Two", "Three", "Four")
	list.Remove(list.Tail.Prev)

	if list.Tail.Prev.Value != "Two" {
		t.Error(`Expected three to be removed`)
	}
}

func TestRemoveLastFromList(t *testing.T) {
	list := NewList("Last")
	list.Remove(list.Head)

	if list.Head != nil {
		t.Error(`Expected head to be nil, got`, list.Head)
	}

	if list.Tail != nil {
		t.Error(`Expected tail to be nil, got`, list.Tail)
	}
}

func TestRemoveFromListTail(t *testing.T) {
	list := NewList()
	a := list.Append("a")
	b := list.Append("b")
	c := list.Append("c")

	list.Remove(c)

	if list.Tail != b {
		t.Error(`Expected tail to be "b", got`, list.Tail)
	}
	if list.Tail.Next != nil {
		t.Error(`Expected next of tail to be nil, got`, list.Tail.Next)
	}
	if list.Tail.Prev != a {
		t.Error(`Expected prev of tail to be "a", got`, list.Tail.Prev)
	}
}

func TestRemoveFromListHead(t *testing.T) {
	list := NewList()
	a := list.Append("a")
	b := list.Append("b")
	c := list.Append("c")

	list.Remove(a)

	if list.Head != b {
		t.Error(`Expected head to be "b", got`, list.Head)
	}
	if list.Head.Next != c {
		t.Error(`Expected next of head to be "c", got`, list.Head.Next)
	}
	if list.Head.Prev != nil {
		t.Error(`Expected prev of head to be nil, got`, list.Head.Prev)
	}
}

func TestRemoveFromListMiddle(t *testing.T) {
	list := NewList()
	a := list.Append("a")
	b := list.Append("b")
	c := list.Append("c")

	list.Remove(b)

	if a.Next != c {
		t.Error(`Expected next of "a" to be "c", got`, a.Next)
	}

	if c.Prev != a {
		t.Error(`Expected prev pf "c" to be "a", got`, c.Prev)
	}
}

func TestRemoveFromListOfTwo(t *testing.T) {
	list := NewList()
	a := list.Append("a")
	b := list.Append("b")

	list.Remove(b)

	if list.Head != list.Tail {
		t.Error(`Expected head and tail to be equal, got`, list.Head, list.Tail)
	}

	if list.Head != a {
		t.Error(`Expected head to be "a", got`, list.Head)
	}
}

func TestRemoveLastListItem(t *testing.T) {
	list := NewList("only")
	list.Remove(list.Head)

	if list.Head != nil {
		t.Error()
	}

	if list.Tail != nil {
		t.Error()
	}
}

func BenchmarkRemoveFromList(b *testing.B) {
	list := BigList(b.N)
	//b.ResetTimer()

	for n := 0; n < b.N; n++ {
		if list.Tail.Prev != nil {
			list.Remove(list.Tail.Prev)
		}
	}
}

// Generate a large list for benchmarks
func BigList(n int) *List {
	list := NewList()
	for i := 0; i <= n; i++ {
		list.Append("Item " + string(i))
	}
	return list
}
