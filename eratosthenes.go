package main

const N = 1000

func SieveOfEratosthenes() [N + 1]bool {
	var i, j int
	var a [N + 1]bool

	for a[1], i = false, 2; i <= N; i++ {
		a[i] = true
	}

	for i = 2; i <= N/2; i++ {
		for j = 2; j <= N/i; j++ {
			a[i*j] = false
		}
	}

	return a
}
