package main

// Euclid's Algorithm
func GCD(u, v int) int {
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
