package main

type Node struct {
	Key  int
	Next *Node
}

// N people have decided to commit mass suicide by arranging themselves in a
// circle and killing the Mth person around the circle, closing ranks as each
// person drops out of the circle. The problem is to find out which person is
// the last to die, or more  generally, to find the order in which the people
// are executed.
//
// Uncomment fmt statements to see a log of each step in the procedure.
func Josephus(n, m int) *Node {
	var i int
	var member, next *Node

	// Initialize the first node in the list
	member = &Node{Key: 1}
	next = member

	// initialize the remaining nodes
	for i = 2; i <= n; i++ {
		member.Next = &Node{Key: i}
		//fmt.Printf("%d is linked to %d\n", member.Key, member.Next.Key)
		member = member.Next
	}

	// complete the cirle
	member.Next = next
	//fmt.Printf("%d is linked to %d\n", member.Key, member.Next.Key)
	//fmt.Println()

	//fmt.Printf("%d is the first member\n", member.Key)

	// while there is still more than one member of the circle
	for member != member.Next {

		// count M members ahead in the circle to find the next to remove
		for i = 1; i < m; i++ {
			member = member.Next
			//fmt.Printf("skipped %d\n", member.Key)
		}
		//fmt.Printf("%d is next to die\n", member.Next.Key)

		// in C, we would need to remember the address of
		// next, so that we could explicitly free it from memory
		next = member.Next
		//fmt.Printf("%d is out of the circle\n", next.Key)

		// here we remove a member by changing the link
		member.Next = member.Next.Next
		//fmt.Printf("%d is now linked to %d\n", member.Key, member.Next.Key)

		// here we ignore "next" and let Go's garbage collector do it's thing
	}

	//fmt.Printf("%d is the last man standing\n", member.Key)
	return member
}
