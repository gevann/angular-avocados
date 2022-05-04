package main

import (
	"angular-avocados/dayone"
	"fmt"
	"os"
)

func recovery() {
	if r := recover(); r != nil {
		fmt.Fprintln(os.Stderr, "Error:", r)
	}
}

func main() {
	defer recovery()
	partial, full := "dayone/d1p1_input.partial.txt", "dayone/d1p1_input.full.txt"
	fmt.Printf("Day One, Part One (Partial): %d\n", dayone.PartOne(partial))
	fmt.Printf("Day One, Part One (Full): %d\n", dayone.PartOne(full))
	fmt.Printf("Day One, Part Two (Partial): %d\n", dayone.PartTwo(partial))
	fmt.Printf("Day One, Part Two (Full): %d\n", dayone.PartTwo(full))
} // end main
