package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func sieveOfEratosthenes(N uint64) (primes []uint64) {
	// TODO(rbtz): convert to bitmask
	primes = append(primes, 2)

	b := make([]bool, N)
	for i := uint64(3); i <= uint64(math.Sqrt(float64(N))); i += 2 {
		if b[i] == true {
			continue
		}
		for k := i * i; k < N; k += i {
			b[k] = true
		}
	}

	for i := uint64(3); i < N; i += 2 {
		if b[i] == false {
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
		for i := lPrimes - 1; i > 0; i-- {
			// TODO(rbtz): skip primes bigger than n
			if n%primes[i] == 0 {
				fmt.Println(primes[i])
				break
			}
		}
	}
}
