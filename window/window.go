package window

import (
	"fmt"
)

type IntWindow struct {
	Sum  int
	Data []int
}

func (w *IntWindow) Len() int {
	return len(w.Data)
}

func (w *IntWindow) String() string {
	return fmt.Sprintf("%v", w.Data)
}

func (w *IntWindow) Append(i int) {
	w.Data = append(w.Data, i)
	w.Sum += i
}

func (w *IntWindow) Remove() {
	datum := w.Data[0]
	w.Data = w.Data[1:]
	w.Sum -= datum
}

func (w *IntWindow) Get(i uint) (int, error) {
	if i > uint(w.Len()-1) {
		return 0, fmt.Errorf("Index out of range")
	}
	return w.Data[i], nil
}

func (w *IntWindow) Last() (int, bool) {
	if w.Len() == 0 {
		return 0, false
	}
	return w.Data[w.Len()-1], true
}

func (w *IntWindow) Reset() {
	w.Sum = 0
	w.Data = nil
}

func Window() *IntWindow {
	return &IntWindow{
		Sum:  0,
		Data: []int{},
	}
}
