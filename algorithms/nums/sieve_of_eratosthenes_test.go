package nums

import (
	"reflect"
	"testing"
)

func TestSieveOfEratosthenes(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		expected map[int][]int
	}{
		{
			name:     "n=0 returns empty map",
			n:        0,
			expected: map[int][]int{},
		},
		{
			name:     "n=1 returns empty map",
			n:        1,
			expected: map[int][]int{},
		},
		{
			name: "n=2 returns only prime 2",
			n:    2,
			expected: map[int][]int{
				2: {2},
			},
		},
		{
			name: "n=10 returns correct prime factorizations",
			n:    10,
			expected: map[int][]int{
				2:  {2},
				3:  {3},
				4:  {2},
				5:  {5},
				6:  {2, 3},
				7:  {7},
				8:  {2},
				9:  {3},
				10: {2, 5},
			},
		},
		{
			name: "primes up to 20 have single-element factor slices equal to themselves",
			n:    20,
			expected: map[int][]int{
				2:  {2},
				3:  {3},
				4:  {2},
				5:  {5},
				6:  {2, 3},
				7:  {7},
				8:  {2},
				9:  {3},
				10: {2, 5},
				11: {11},
				12: {2, 3},
				13: {13},
				14: {2, 7},
				15: {3, 5},
				16: {2},
				17: {17},
				18: {2, 3},
				19: {19},
				20: {2, 5},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := sieveOfEratosthenes(tt.n)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("sieveOfEratosthenes(%d) = %v, want %v", tt.n, result, tt.expected)
			}
		})
	}
}

func TestSieveOfEratosthenesPrimesOnly(t *testing.T) {
	primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}
	result := sieveOfEratosthenes(30)

	for _, p := range primes {
		factors, ok := result[p]
		if !ok {
			t.Errorf("prime %d not found in result", p)
			continue
		}
		if len(factors) != 1 || factors[0] != p {
			t.Errorf("prime %d has factors %v, want [%d]", p, factors, p)
		}
	}
}

func TestSieveOfEratosthenesMapSize(t *testing.T) {
	n := 15
	result := sieveOfEratosthenes(n)
	// every number from 2..n should be a key
	if len(result) != n-1 {
		t.Errorf("expected %d entries, got %d", n-1, len(result))
	}
	for i := 2; i <= n; i++ {
		if _, ok := result[i]; !ok {
			t.Errorf("missing key %d in result", i)
		}
	}
}
