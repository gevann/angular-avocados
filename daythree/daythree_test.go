package daythree

import (
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
