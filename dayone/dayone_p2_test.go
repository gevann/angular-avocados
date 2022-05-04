package dayone

import (
	"angular-avocados/window"
	"reflect"
	"strings"
	"testing"
)

func Test_sum(t *testing.T) {
	type args struct {
		data []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "returns the correct sum",
			args: args{
				data: []int{1, 2, 3},
			},
			want: 6,
		},
		{
			name: "returns 0 when the slice is empty",
			args: args{
				data: []int{},
			},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sum(tt.args.data); got != tt.want {
				t.Errorf("sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sumGenerator(t *testing.T) {
	type args struct {
		chIn <-chan window.Windower[int]
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "converts the channel of windows to another channel of windows",
			args: args{
				chIn: generator(strings.NewReader(partialInput), 3),
			},
			want: [][]int{
				{607, 618},
				{618, 618},
				{618, 617},
				{617, 647},
				{647, 716},
				{716, 769},
				{769, 792},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			generator := sumGenerator(tt.args.chIn)
			for _, want := range tt.want {
				got := <-generator
				if !reflect.DeepEqual(got.Data(), want) {
					t.Errorf("sumGenerator() = %v, want %v", got, want)
				}
			}
		})
	}
}
