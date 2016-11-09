package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	resultMoulus = 1000000007
)

type countTable map[int]map[int]map[int]uint64

func recurseDigits(t countTable, i1, i2, n int) uint64 {
	if n == 0 {
		return 1
	}

	if _, ok := t[i1]; !ok {
		t[i1] = make(map[int]map[int]uint64)
	}
	if _, ok := t[i1][i2]; !ok {
		t[i1][i2] = make(map[int]uint64)
	}

	if t[i1][i2][n] == 0 {
		for i := 0; i < 10-i2-i1; i++ {
			t[i1][i2][n] += recurseDigits(t, i2, i, n-1)
		}
	}

	return t[i1][i2][n] % resultMoulus
}

func main() {
	var result uint64

	// Get input
	reader := bufio.NewReader(os.Stdin)
	mStr, _ := reader.ReadString('\n')
	m, _ := strconv.Atoi(strings.TrimSpace(mStr))
	if m > 100 || m < 3 {
		panic("m is out of bounds")
	}

	// Recursively compute num of digits
	t := make(countTable)
	for i1 := 1; i1 < 10; i1++ {
		for i2 := 0; i2 < 10-i1; i2++ {
			result += recurseDigits(t, i1, i2, m-2)
			result %= resultMoulus
		}
	}

	fmt.Println(result)
}
