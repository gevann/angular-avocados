package dayone

import (
	"angular-avocados/window"
	"os"
)

func sum(data []int) int {
	sum := 0
	for _, i := range data {
		sum += i
	}
	return sum
}

func sumGenerator(chIn <-chan window.Windower[int]) <-chan window.Windower[int] {
	chOut := make(chan window.Windower[int])
	go func() {
		defer close(chOut)
		windowPrev := <-chIn

		for windowCurr := range chIn {
			sumPrev := sum(windowPrev.Data())
			sumCurr := sum(windowCurr.Data())
			chOut <- window.New(sumPrev, sumCurr)
			windowPrev = windowCurr
		}
	}()
	return chOut
}

func PartTwo(filePath string) (int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return 0, err
	}
	windowLen := uint(3)
	ch := sumGenerator(generator(file, windowLen))
	increases := increases(ch, lastGreaterThanFirst)
	return increases, nil
}
