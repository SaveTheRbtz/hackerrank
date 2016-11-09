package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const size = 64

type bits uint64

// BitSet is a set of bits that can be set, cleared and queried.
type BitSet []bits

// Set ensures that the given bit is set in the BitSet.
func (s *BitSet) Set(i uint64) {
	(*s)[i/size] |= 1 << (i % size)
}

// Clear ensures that the given bit is cleared (not set) in the BitSet.
func (s *BitSet) Clear(i uint64) {
	(*s)[i/size] &^= 1 << (i % size)
}

// IsSet returns true if the given bit is set, false if it is cleared.
func (s *BitSet) IsSet(i uint64) bool {
	return (*s)[i/size]&(1<<(i%size)) != 0
}

func sieveOfEratosthenes(N uint64) (primes []uint64) {
	primes = append(primes, 2)

	b := make(BitSet, N)
	for i := uint64(3); i <= uint64(math.Sqrt(float64(N))); i += 2 {
		if b.IsSet(i) {
			continue
		}
		for k := i * i; k < N; k += i {
			b.Set(k)
		}
	}

	// TODO(rbtz): use bitmask directly without copying it
	for i := uint64(3); i < N; i += 2 {
		if !b.IsSet(i) {
			primes = append(primes, i)
		}
	}
	return
}

func main() {
	// Get input len
	reader := bufio.NewReader(os.Stdin)
	tStr, _ := reader.ReadString('\n')
	t, _ := strconv.Atoi(strings.TrimSpace(tStr))

	// Iterate over input
	input := make([]uint64, t)
	for i := 0; i < t; i++ {
		nStr, _ := reader.ReadString('\n')
		n, _ := strconv.ParseUint(strings.TrimSpace(nStr), 10, 100)
		input[i] = n
	}

	// Find max
	biggest := uint64(0)
	for _, n := range input {
		if biggest < n {
			biggest = n
		}
	}

	primes := sieveOfEratosthenes(biggest + 1)
	lPrimes := len(primes)

	for _, n := range input {
		// Approximate number of primes below n
		maxIndex := uint64(lPrimes - 1)
		approxIndex := uint64(float64(n) / (math.Log(float64(n)) - 1.1))
		if approxIndex < maxIndex {
			maxIndex = approxIndex
		}
		// Move right until the end or prime that is >= n
		for ; maxIndex < uint64(lPrimes)-1 && n >= primes[maxIndex]; maxIndex++ {
		}
		// Find the biggest factors out of all primes
		for i := maxIndex; i > 0; i-- {
			p := primes[i]
			if n%p == 0 {
				fmt.Println(p)
				break
			}
		}
	}
}
