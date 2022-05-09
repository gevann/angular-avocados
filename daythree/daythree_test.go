package daythree

import (
	"errors"
	"reflect"
	"testing"
)

func Test_parseInput(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "parse partial input file",
			args: args{
				filename: "input.partial.txt",
			},
			want: []string{
				"00100",
				"11110",
				"10110",
				"10111",
				"10101",
				"01111",
				"00111",
				"11100",
				"10000",
				"11001",
				"00010",
				"01010",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			channel, err := parseInput(tt.args.filename)

			if (err != nil) != tt.wantErr {
				t.Errorf("parseInput() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			for _, wanted := range tt.want {
				got := <-channel
				if !reflect.DeepEqual(got, wanted) {
					t.Errorf("parseInput() = %v, want %v", got, wanted)
				}
			}
		})
	}
}

func Test_binaryStringToDigits(t *testing.T) {
	type args struct {
		s         string
		converter func(r rune) (int, error)
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr bool
	}{
		{
			name: "parse binary string with converter",
			args: args{
				s: "00100",
				converter: func(r rune) (int, error) {
					switch r {
					case '0':
						return -1, nil
					case '1':
						return 1, nil
					default:
						return 0, errors.New("invalid character")
					}
				},
			},
			want:    []int{-1, -1, 1, -1, -1},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := binaryStringToDigits(tt.args.s, tt.args.converter)
			if (err != nil) != tt.wantErr {
				t.Errorf("binaryStringToDigits() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("binaryStringToDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_digitsToBinaryString(t *testing.T) {
	type args struct {
		digits    []int
		converter func(i int) (rune, rune, error)
	}
	tests := []struct {
		name         string
		args         args
		want         string
		wantInverted string
		wantErr      bool
	}{
		{
			name: "convert digits to string, with converter",
			args: args{
				digits: []int{-1, -1, 1, -1, -1},
				converter: func(i int) (rune, rune, error) {
					switch i {
					case -1:
						return '0', '1', nil
					case 1:
						return '1', '0', nil
					default:
						return ' ', ' ', errors.New("invalid digit")
					}
				},
			},
			want:         "00100",
			wantInverted: "11011",
			wantErr:      false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, inverted, err := digitsToBinaryStrings(tt.args.digits, tt.args.converter)
			if (err != nil) != tt.wantErr {
				t.Errorf("digitsToBinaryString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("digitsToBinaryString() = %v, want %v", got, tt.want)
			}
			if inverted != tt.wantInverted {
				t.Errorf("digitsToBinaryString() = %v, want %v", inverted, tt.wantInverted)
			}
		})
	}
}

func Test_addIntArrays(t *testing.T) {
	type args struct {
		a []int
		b []int
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr bool
	}{
		{
			name: "add two arrays",
			args: args{
				a: []int{1, 2, 3},
				b: []int{4, 5, 6},
			},
			want:    []int{5, 7, 9},
			wantErr: false,
		},
		{
			name: "errors when arrays are of different length",
			args: args{
				a: []int{1, 2, 3},
				b: []int{4, 5},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := addIntArrays(tt.args.a, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("addIntArrays() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addIntArrays() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convertZeroToNegativeOne(t *testing.T) {
	type args struct {
		r rune
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "convert zero to negative one",
			args: args{
				r: '0',
			},
			want:    -1,
			wantErr: false,
		},
		{
			name: "attempt Atoi on non-zero character",
			args: args{
				r: '1',
			},
			want: 1,
		},
		{
			name: "return error on invalid character",
			args: args{
				r: ' ',
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := convertZeroToNegativeOne(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("convertZeroToNegativeOne() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("convertZeroToNegativeOne() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_toBinary(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name       string
		args       args
		want       rune
		wantInvert rune
		wantErr    bool
	}{
		{
			name: "convert negatives to 0 and 1",
			args: args{
				i: -1,
			},
			want:       '0',
			wantInvert: '1',
			wantErr:    false,
		},
		{
			name: "convert positives to 1 and 0",
			args: args{
				i: 1,
			},
			want:       '1',
			wantInvert: '0',
			wantErr:    false,
		},
		{
			name: "return error when i is 0",
			args: args{
				i: 0,
			},
			want:       ' ',
			wantInvert: ' ',
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := toBinary(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("toBinary() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("toBinary() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.wantInvert {
				t.Errorf("toBinary() got1 = %v, want %v", got1, tt.wantInvert)
			}
		})
	}
}
