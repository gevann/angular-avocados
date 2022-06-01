package main

import (
	"angular-avocados/dayone"
	"angular-avocados/daythree"
	"angular-avocados/daytwo"
	"flag"
	"fmt"
	"os"
	"strings"
)

func recovery() {
	if r := recover(); r != nil {
		fmt.Fprintln(os.Stderr, "Error:", r)
	}
}

type printable[T interface{}] struct {
	part     uint8
	filename string
	fn       func(s string) (T, error)
}

func print[T interface{}](day uint8, args []printable[T]) {
	printStr := fmt.Sprintf("\n################################ DAY %d #######################################\n", day)

	for _, p := range args {

		result, err := p.fn(p.filename)
		errString := ""
		if err != nil {
			errString = err.Error()
		}

		inputType := "Full"
		// check if filename containts a partial or full
		if strings.Contains(p.filename, "partial") {
			inputType = "Partial"
		}

		printStr += fmt.Sprintf("\tPart %d (%s):\n\t- Result:%v\n\t- Error:'%s'\n", p.part, inputType, result, errString)
	}
	fmt.Println(printStr)
}

var days = []uint{1, 2, 3}

func main() {
	defer recovery()
	fmt.Printf("\n\n~~~~~~~~~~~ Advent of Code 2021 ~~~~~~~~~~~~\n")

	var selectedDate uint
	flag.UintVar(&selectedDate, "days", days[len(days)-1], "Day to run")
	flag.Parse()

	switch selectedDate {
	case 1:
		dayOne()
	case 2:
		dayTwo()
	case 3:
		dayThree()
	default:
		dayOne()
		dayTwo()
		dayThree()
	}
}

func dayOne() {
	print(1, []printable[int]{
		{1, "dayone/d1p1_input.partial.txt", dayone.PartOne},
		{1, "dayone/d1p1_input.full.txt", dayone.PartOne},
		{2, "dayone/d1p1_input.partial.txt", dayone.PartTwo},
		{2, "dayone/d1p1_input.full.txt", dayone.PartTwo},
	})
}

func dayTwo() {
	print(2, []printable[int]{
		{1, "daytwo/d2p1_input.partial.txt", daytwo.PartOne},
		{1, "daytwo/d2p1_input.full.txt", daytwo.PartOne},
		{2, "daytwo/d2p1_input.partial.txt", daytwo.PartTwo},
		{2, "daytwo/d2p1_input.full.txt", daytwo.PartTwo},
	})
}

func dayThree() {
	print(3, []printable[int]{
		{1, "daythree/input.partial.txt", daythree.PartOne},
		{1, "daythree/input.full.txt", daythree.PartOne},
		{2, "daythree/input.partial.txt", daythree.PartTwo},
		{2, "daythree/input.full.txt", daythree.PartTwo},
	})
}
