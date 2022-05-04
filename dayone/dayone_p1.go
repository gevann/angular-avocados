package dayone

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type IntTuple struct {
	A int
	B int
}

func parse(s string) (int, error) {
	i, e := strconv.Atoi(strings.Trim(s, " \t\r\n,"))
	return i, e
}

func generator(r io.Reader) <-chan IntTuple { // returns receive-only channel
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	ch := make(chan IntTuple)

	init := scanner.Text()
	a, err := parse(init)
	if err != nil {
		panic("Could not parse initial value")
	}

	go func() { // anonymous goroutine
		for scanner.Scan() {
			b, err := parse(scanner.Text())
			if err != nil {
				panic("Could not value")
			}
			tuple := IntTuple{a, b} // create a tuple
			ch <- tuple             // send tuple to channel
			a = b                   // update a
		}
		close(ch) // close channel
	}()
	return ch
}

func increases(ch <-chan IntTuple) int {
	count := 0
	for tuple := range ch {
		if tuple.A < tuple.B {
			count++
		}
	}
	return count
}

func Main(filePath string) int {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	ch := generator(file)
	increases := increases(ch)
	fmt.Println(increases)
	return increases
} // end main
