package nums

// sieveOfEratosthenes returns a map where the keys are prime numbers up to n and the values are slices of their multiples.
func sieveOfEratosthenes(n int) map[int][]int {
	factors := make(map[int][]int, n)
	for i := 2; i <= n; i++ {
		if _, exists := factors[i]; !exists {
			for j := i; j <= n; j += i {
				factors[j] = append(factors[j], i)
			}
		}
	}
	return factors
}
