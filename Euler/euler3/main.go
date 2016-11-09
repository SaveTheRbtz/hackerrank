package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	size = 64
)

type bits []uint64

// BitSet is a set of bits that can be set, cleared and queried.
type BitSet struct {
	bits     bits
	len      uint64
	capacity uint64
}

// Set ensures that the given bit is set in the BitSet.
func (s *BitSet) Set(i uint64) {
	s.bits[i/size] |= 1 << (i % size)
}

// Clear ensures that the given bit is cleared (not set) in the BitSet.
func (s *BitSet) Clear(i uint64) {
	s.bits[i/size] &^= 1 << (i % size)
}

// IsSet returns true if the given bit is set, false if it is cleared.
func (s *BitSet) IsSet(i uint64) bool {
	return s.bits[i/size]&(1<<(i%size)) != 0
}

// PrevUnset returns previous unset bit.
func (s *BitSet) PrevUnset(current uint64) uint64 {
	for i := current; i > 0; i-- {
		if !s.IsSet(i) && i <= s.len {
			return i
		}
	}
	return 0 // XXX should never happen
}

func sieveOfEratosthenes(N uint64) (b BitSet) {
	bCap := N/size + 1
	b.bits = make(bits, bCap)
	b.capacity = bCap
	b.len = N - 1

	// TODO(rbtz): skip even bits
	b.Set(1)
	for i := uint64(2); i <= uint64(math.Sqrt(float64(N))); i++ {
		if b.IsSet(i) {
			continue
		}
		for k := i * i; k < N; k += i {
			b.Set(k)
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

	// Initialize bitset
	bitset := sieveOfEratosthenes(biggest + 1)
	// TODO(rbtz): run in parallel
	for _, n := range input {
		current := n
		for {
			p := bitset.PrevUnset(current)
			if n%p == 0 {
				fmt.Println(p)
				break
			}
			current = p - 1
		}
	}
}
