package main

// Euclid's Algorithm

func GCDIterative(u, v int) int {
	var t int
	for u > 0 {
		if u < v {
			t = u
			u = v
			v = t
		}
		u = u - v
	}
	return v
}

func GCDRecursive(p, q int) int {
	if q == 0 {
		return p
	}

	r := p % q
	return GCDRecursive(q, r)
}
