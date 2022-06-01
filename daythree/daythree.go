package daythree

import (
	"angular-avocados/tree"
	"bufio"
	"errors"
	"os"
	"strconv"
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

func binaryStringToDigits(s string, converter func(r rune) (int, error)) ([]int, error) {
	var digits []int
	for _, r := range s {
		digit, err := converter(r)
		if err != nil {
			return nil, err
		}
		digits = append(digits, digit)
	}
	return digits, nil
}

func digitsToBinaryStrings(digits []int, converter func(i int) (rune, rune, error)) (string, string, error) {
	var s []rune
	var inverted []rune
	for _, i := range digits {
		r0, r1, e := converter(i)
		if e != nil {
			return "", "", e
		}
		s = append(s, r0)
		inverted = append(inverted, r1)
	}
	return string(s), string(inverted), nil
}

func addIntArrays(a, b []int) ([]int, error) {
	if len(a) != len(b) {
		return nil, errors.New("arrays must be of equal length")
	}
	var c []int
	for i, v := range a {
		c = append(c, v+b[i])
	}
	return c, nil
}

func convertZeroToNegativeOne(r rune) (int, error) {
	if r == '0' {
		return -1, nil
	}
	i, err := strconv.Atoi(string(r))
	if err != nil {
		return 0, err
	}
	return i, nil
}

func toBinary(i int) (rune, rune, error) {
	if i < 0 {
		return '0', '1', nil
	} else if i > 0 {
		return '1', '0', nil
	} else {
		return ' ', ' ', errors.New("No tie breaker.")
	}
}

func PartOne(filename string) (int, error) {
	var zeroValue int
	ch, err := parseInput(filename)
	if err != nil {
		return zeroValue, err
	}
	digits, err := binaryStringToDigits(<-ch, convertZeroToNegativeOne)

	if err != nil {
		return zeroValue, err
	}
	for binaryStr := range ch {
		nextDigits, err := binaryStringToDigits(binaryStr, convertZeroToNegativeOne)
		if err != nil {
			return zeroValue, err
		}
		digits, err = addIntArrays(digits, nextDigits)
		if err != nil {
			return zeroValue, err
		}
	}

	a, aNaught, err := digitsToBinaryStrings(digits, toBinary)
	if err != nil {
		return zeroValue, err
	}

	x, e := strconv.ParseInt(a, 2, 64)
	if e != nil {
		return zeroValue, e
	}
	y, e := strconv.ParseInt(aNaught, 2, 64)
	if e != nil {
		return zeroValue, e
	}

	return int(x * y), nil
}

func PartTwo(filename string) (int, error) {
	var zeroValue int
	ch, err := parseInput(filename)
	if err != nil {
		return zeroValue, err
	}

	tree := &tree.Tree{}

	for binaryStr := range ch {
		tree.Insert(binaryStr)
	}

	a := tree.MaxByLeafCount()
	aNaught := tree.MinByLeafCount()

	x, e := strconv.ParseInt(a.Value, 2, 64)
	if e != nil {
		return zeroValue, e
	}
	y, e := strconv.ParseInt(aNaught.Value, 2, 64)
	if e != nil {
		return zeroValue, e
	}

	return int(x * y), nil
}
