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
	dayone.PartOne("dayone/d1p1_input.partial.txt")
	dayone.PartOne("dayone/d1p1_input.full.txt")
} // end main
