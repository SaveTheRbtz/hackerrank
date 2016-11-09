package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func sieveOfEratosthenes(N uint64) (primes []uint64) {
	// TODO(rbtz): convert to bitmask
	b := make([]bool, N)
	// TODO(rbtz): skip even
	// TODO(rbtz): up to sqrt(N)
	for i := uint64(2); i < N; i++ {
		if b[i] == true {
			continue
		}
		primes = append(primes, i)
		for k := i * i; k < N; k += i {
			b[k] = true
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

	//TODO(rbtz): cache?
	for _, n := range input {
		for i := lPrimes - 1; i > 0; i-- {
			if n%primes[i] == 0 {
				fmt.Println(primes[i])
				break
			}
		}
	}
}
