package dayone

import (
	"angular-avocados/window"
	"io"
	"reflect"
	"strings"
	"testing"
)

const partialInput = `199,
200,
208,
210,
200,
207,
240,
269,
260,
263
`

func Test_parse(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
		{"positives", args{"123"}, 123, false},
		{"negatives", args{"-123"}, -123, false},
		{"whitespaces", args{" \t 123 \n \r  "}, 123, false},
		{"commas", args{",123"}, 123, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parse(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("parse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_generator(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"returns windows with the correct Data", args{strings.NewReader(partialInput)}, [][]int{
			{199, 200},
			{200, 208},
			{208, 210},
			{210, 200},
			{200, 207},
			{207, 240},
			{240, 269},
			{269, 260},
			{260, 263},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			generator := generator(tt.args.r, 2)
			for _, want := range tt.want {
				got := <-generator
				if !reflect.DeepEqual(got.Data(), want) {
					t.Errorf("Generator() = %v, want %v", got, want)
				}
			}
		})
	}
}

func Test_increases(t *testing.T) {
	type args struct {
		ch <-chan window.Windower[int]
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"partial", args{generator(strings.NewReader(partialInput), 2)}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := increases(tt.args.ch, lastGreaterThanFirst); got != tt.want {
				t.Errorf("increases() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lastGreaterThanFirst(t *testing.T) {
	type args struct {
		w window.Windower[int]
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"returns true when the last element is greater than the first",
			args{window.New(1, 2, 3)},
			true,
		},
		{
			"returns false when the last element is not greater than the first",
			args{window.New(1, 2, 1)},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lastGreaterThanFirst(tt.args.w); got != tt.want {
				t.Errorf("lastGreaterThanFirst() = %v, want %v", got, tt.want)
			}
		})
	}
}
