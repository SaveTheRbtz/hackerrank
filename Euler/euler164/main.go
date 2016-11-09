package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
)

const (
	resultMoulus = 1000000000 + 7
	charDiff     = 48 * 3
)

var (
	wg        sync.WaitGroup
	goMaxProc = runtime.GOMAXPROCS(0)
)

func genMDigitNumbers(m int, c chan string, quit chan struct{}, numWorkers int) {
	for i := uint64(math.Pow10(m - 1)); i < uint64(math.Pow10(m)); i++ {
		fmt.Println(i)
		c <- strconv.FormatUint(i, 10)
	}
	for i := 0; i < numWorkers; i++ {
		quit <- struct{}{}
	}
}

func worker(
	quit chan struct{},
	numbers chan string,
	wg *sync.WaitGroup,
	result *uint64) {

	defer wg.Done()

NumberSelect:
	for {
		select {
		case number := <-numbers:
			numLen := len(number)
			// TODO(rbtz): table lookup
			for i := 0; i < numLen-2; i++ {
				if number[i]+number[i+1]+number[i+2]-charDiff > 9 {
					continue NumberSelect
				}
			}
			atomic.AddUint64(result, 1)
		case <-quit:
			return
		}
	}
}

func main() {
	var result uint64

	// Spawn workers
	numbers := make(chan string, goMaxProc)
	quit := make(chan struct{})
	for i := 0; i < goMaxProc; i++ {
		wg.Add(1)
		go worker(quit, numbers, &wg, &result)
	}

	// Produce data
	reader := bufio.NewReader(os.Stdin)
	mStr, _ := reader.ReadString('\n')
	m, _ := strconv.Atoi(strings.TrimSpace(mStr))
	if m > 100 || m < 3 {
		panic("m is out of bounds")
	}
	go genMDigitNumbers(m, numbers, quit, goMaxProc)

	wg.Wait()
	fmt.Println(result % resultMoulus)
}
