package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type indexedValue struct {
	index int
	value int
}

type resultValue struct {
	result uint64
	input  indexedValue
}

type byValue struct{ values []indexedValue }

func (s byValue) Len() int           { return len(s.values) }
func (s byValue) Swap(i, j int)      { s.values[i], s.values[j] = s.values[j], s.values[i] }
func (s byValue) Less(i, j int) bool { return s.values[i].value < s.values[j].value }

type byInputIndex struct{ results []resultValue }

func (s byInputIndex) Len() int           { return len(s.results) }
func (s byInputIndex) Swap(i, j int)      { s.results[i], s.results[j] = s.results[j], s.results[i] }
func (s byInputIndex) Less(i, j int) bool { return s.results[i].input.index < s.results[j].input.index }

func main() {
	// Get input len
	reader := bufio.NewReader(os.Stdin)
	tStr, _ := reader.ReadString('\n')
	t, _ := strconv.Atoi(strings.TrimSpace(tStr))

	// Iterate over input
	input := make([]indexedValue, t)
	for i := 0; i < t; i++ {
		nStr, _ := reader.ReadString('\n')
		n, _ := strconv.Atoi(strings.TrimSpace(nStr))
		input[i] = indexedValue{i, n}
	}

	// Sort input by value
	sort.Sort(byValue{input})

	// Init cache
	resultCache := make([]resultValue, t)

	for _, m := range input {
		var result, maxR uint64
		var maxV int

		// Lookup in cache
		for _, c := range resultCache {
			if c.result > maxR {
				maxR = c.result
				maxV = c.input.value
			}
		}
		result = maxR

		for i := maxV; i < m.value; i++ {
			if i%3 == 0 || i%5 == 0 {
				result += uint64(i)
			}
		}
		resultCache[m.index] = resultValue{result, m}
	}

	// Sort input by value
	sort.Sort(byInputIndex{resultCache})

	// Print results
	for _, r := range resultCache {
		fmt.Println(r.result)
	}
}
