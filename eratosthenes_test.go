package main

import "testing"

func isPrime(n int) bool {
	var i int
	for i = 2; i < n; i++ {
		if n%i == 0 && i != n {
			return false
		}
	}
	return true
}

func TestEratosthenes(t *testing.T) {
	primes := SieveOfEratosthenes()
	for i, p := range primes {
		primality := isPrime(i)
		if i > 1 && p != primality {
			t.Error("Expected primality of", i, "to be", primality, "got", p)
		}
	}
}

func BenchmarkEratosthenes(b *testing.B) {
	for n := 0; n < b.N; n++ {
		SieveOfEratosthenes()
	}
}
