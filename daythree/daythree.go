package daythree

import (
	"bufio"
	"os"
)

func parseInput(filename string) (<-chan string, error) {
	file, err := os.Open(filename)

	if err != nil {
		return nil, err
	}

	ch := make(chan string)

	go func() {
		defer file.Close()
		defer close(ch)

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			val := scanner.Text()
			ch <- val
		}
	}()

	return ch, nil
}
