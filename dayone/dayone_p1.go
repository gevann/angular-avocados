package dayone

import (
	"angular-avocados/window"
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func parse(s string) (int, error) {
	i, e := strconv.Atoi(strings.Trim(s, " \t\r\n,"))
	return i, e
}

func generator(r io.Reader, windowLen uint) <-chan window.Windower[int] { // returns receive-only channel
	scanner := bufio.NewScanner(r)
	ch := make(chan window.Windower[int])

	go func() { // anonymous goroutine
		defer close(ch)
		windowPrev := window.New[int]()
		windowCurr := window.New[int]()

		// iterate windowLen times
		for i := uint(0); i < windowLen; i++ {
			scanner.Scan()
			datum, err := parse(scanner.Text())
			if err != nil {
				panic(err)
			}
			windowPrev.Append(datum)
		}

		ch <- windowPrev

		for scanner.Scan() {
			i, err := parse(scanner.Text())
			if err != nil {
				panic(err)
			}
			for _, datum := range windowPrev.Data()[1:] {
				windowCurr.Append(datum)
			}

			windowCurr.Append(i)
			if windowCurr.Len() == int(windowLen) {
				ch <- windowCurr
				windowPrev = windowCurr
				windowCurr = window.New[int]()
			} else {
				break
			}
		}
	}()
	return ch
}

func increases(ch <-chan window.Windower[int], comparator func(w window.Windower[int]) bool) int {
	count := 0
	for window := range ch {
		if comparator(window) {
			count++
		}
	}
	return count
}

func lastGreaterThanFirst(w window.Windower[int]) bool {
	first, err := w.Get(0)
	if err != nil {
		panic(err)
	}
	if last, ok := w.Last(); ok {
		if last > first {
			return true
		}
	}
	return false
}

func PartOne(filePath string) int {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	ch := generator(file, 2)
	increases := increases(ch, lastGreaterThanFirst)
	fmt.Println(increases)
	return increases
}
