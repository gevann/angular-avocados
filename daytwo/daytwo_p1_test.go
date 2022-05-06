package daytwo

import (
	"reflect"
	"testing"
)

func Test_parseOpCode(t *testing.T) {
	type args struct {
		s string
	}
	var zeroValue OpCode
	tests := []struct {
		name    string
		args    args
		want    OpCode
		wantErr bool
	}{
		{
			"converts a string starting with f to Forward",
			args{"foo"},
			forward,
			false,
		},
		{
			"converts a string starting with d to Down",
			args{"dart"},
			down,
			false,
		},
		{
			"converts a string starting with u to Up",
			args{"unicorn"},
			up,
			false,
		},
		{
			"ignore case",
			args{"Foo"},
			forward,
			false,
		},
		{
			"returns an error in all other cases",
			args{"bar"},
			zeroValue,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseOpCode(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseOpCode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("parseOpCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parse(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    Instruction
		wantErr bool
	}{
		{
			"converts a string to an Instruction",
			args{"f 5"},
			Instruction{forward, 5},
			false,
		},
		{
			"panics if the string cannot be converted into an OpCode",
			args{" invalid 5"},
			Instruction{},
			true,
		},
		{
			"panics if the string's OpCode is not followed by a uint8 number",
			args{"f"},
			Instruction{},
			true,
		},
		{
			"panics if the string's OpCode is not followed by a uint8 number",
			args{"f -1"},
			Instruction{},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parse(tt.args.s)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parse() = %v, want %v", got, tt.want)
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("parse() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_channelInstructions(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		want    []Result[Instruction]
		wantErr bool
	}{
		{
			"converts a file to a channel of Instructions",
			args{"d2p1_input.partial.txt"},
			[]Result[Instruction]{
				{Data: Instruction{forward, 5}},
				{Data: Instruction{down, 5}},
				{Data: Instruction{forward, 8}},
				{Data: Instruction{up, 3}},
				{Data: Instruction{down, 8}},
				{Data: Instruction{forward, 2}},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			channel, err := channelInstructions(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("channelInstructions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			for _, want := range tt.want {
				got := <-channel
				if !reflect.DeepEqual(got.Data, want.Data) {
					t.Errorf("channelInstructions() = %v, want %v", got, want)
				}
			}
		})
	}
}

func TestPartOne(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			"solves the puzzle",
			args{"d2p1_input.partial.txt"},
			150,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PartOne(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("PartOne() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("PartOne() = %v, want %v", got, tt.want)
			}
		})
	}
}
