package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func sumOfMultiples(d uint64, max uint64) uint64 {
	N := (max - 1) / d
	return d * N * (N + 1) / 2
}

func main() {
	// Get input len
	reader := bufio.NewReader(os.Stdin)
	tStr, _ := reader.ReadString('\n')
	t, _ := strconv.Atoi(strings.TrimSpace(tStr))

	// Iterate over input
	for i := 0; i < t; i++ {
		nStr, _ := reader.ReadString('\n')
		n, _ := strconv.ParseUint(strings.TrimSpace(nStr), 10, 100)
		result := sumOfMultiples(3, n) + sumOfMultiples(5, n) - sumOfMultiples(15, n)
		fmt.Println(result)
	}
}
