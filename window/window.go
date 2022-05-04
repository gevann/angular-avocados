package window

import (
	"fmt"
)

// Generic Windower interface for generic data type T
type Windower[T any] interface {
	Len() int
	String() string
	Append(i T)
	Remove()
	Get(i uint) (T, error)
	Last() (T, bool)
	Reset()
	Data() []T
}

type Window[T any] struct {
	sum  int
	data []T
}

func (w *Window[T]) Data() []T {
	return w.data
}

func (w *Window[T]) Len() int {
	return len(w.data)
}

func (w *Window[T]) String() string {
	return fmt.Sprintf("%v", w.data)
}

func (w *Window[T]) Append(i T) {
	w.data = append(w.data, i)
}

func (w *Window[T]) Remove() {
	w.data = w.data[1:]
}

func (w *Window[T]) Get(i uint) (T, error) {
	if i > uint(w.Len()-1) {
		var zeroValue T
		return zeroValue, fmt.Errorf("Index out of range")
	}
	return w.data[i], nil
}

func (w *Window[T]) Last() (T, bool) {
	if w.Len() == 0 {
		var zeroValue T
		return zeroValue, false
	}
	return w.data[w.Len()-1], true
}

func (w *Window[T]) Reset() {
	w.data = nil
}

func New[T any](data ...T) Windower[T] {
	return &Window[T]{
		data: data,
	}
}
