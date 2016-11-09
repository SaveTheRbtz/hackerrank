package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Get input len
	reader := bufio.NewReader(os.Stdin)
	tStr, _ := reader.ReadString('\n')
	t, _ := strconv.Atoi(strings.TrimSpace(tStr))

	// Iterate over input
	input := make([]int, t)
	for i := 0; i < t; i++ {
		nStr, _ := reader.ReadString('\n')
		n, _ := strconv.Atoi(strings.TrimSpace(nStr))
		input[i] = n
	}

	for _, m := range input {
		var result uint64

		for i := 0; i < m; i++ {
			if i%3 == 0 || i%5 == 0 {
				result += uint64(i)
			}
		}
		fmt.Println(result)
	}
}
