package window

import (
	"reflect"
	"testing"
)

func TestIntWindow_Len(t *testing.T) {
	type fields struct {
		data []int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "returns the correct length",
			fields: fields{
				data: []int{1, 2, 3},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := New(tt.fields.data...)
			if got := w.Len(); got != tt.want {
				t.Errorf("IntWindow.len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntWindow_String(t *testing.T) {
	type fields struct {
		data []int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "prints the correct string",
			fields: fields{
				data: []int{1, 2, 3},
			},
			want: "[1 2 3]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := New(tt.fields.data...)
			if got := w.String(); got != tt.want {
				t.Errorf("IntWindow.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntWindow_Append(t *testing.T) {
	type fields struct {
		data []int
	}
	type args struct {
		i int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "appends the correct value",
			fields: fields{
				data: []int{1, 2},
			},
			args: args{i: 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := New(tt.fields.data...)
			w.Append(tt.args.i)
			if len(w.Data()) != 3 {
				t.Errorf("IntWindow.append() data = %v, want %v", w.Data(), []int{1, 2, 3})
			}
			last := w.Data()[len(w.Data())-1]
			if last != tt.args.i {
				t.Errorf("IntWindow.append() last = %v, want %v", last, tt.args.i)
			}

		})
	}
}

func TestIntWindow_Remove(t *testing.T) {
	type fields struct {
		data []int
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "removes the first element",
			fields: fields{
				data: []int{1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := New(tt.fields.data...)
			w.Remove()
			if len(w.Data()) != 0 {
				t.Errorf("IntWindow.remove() data = %v, want %v", w.Data(), []int{})
			}
		})
	}
}

func TestIntWindow_Get(t *testing.T) {
	type fields struct {
		length int
		sum    int
		data   []int
	}
	type args struct {
		i uint
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "returns the correct value",
			fields: fields{
				data: []int{1, 2, 3},
			},
			args:    args{i: 1},
			want:    2,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := New(tt.fields.data...)
			got, err := w.Get(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("IntWindow.get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IntWindow.get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want *Window[int]
	}{
		{
			name: "creates a new window with default values",
			args: []int{},
			want: &Window[int]{
				data: []int{},
			},
		},
		{
			name: "creates a new window with the given values",
			args: []int{1, 2, 3},
			want: &Window[int]{
				data: []int{1, 2, 3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntWindow_Last(t *testing.T) {
	type fields struct {
		data []int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
		wantOk bool
	}{
		{
			name: "returns the last element",
			fields: fields{
				data: []int{1, 2, 3, 4, 5},
			},
			want:   5,
			wantOk: true,
		},
		{
			name: "returns zero if the window is empty",
			fields: fields{
				data: []int{},
			},
			want:   0,
			wantOk: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := New(tt.fields.data...)
			got, gotOk := w.Last()
			if got != tt.want {
				t.Errorf("IntWindow.Last() got = %v, want %v", got, tt.want)
			}
			if gotOk != tt.wantOk {
				t.Errorf("IntWindow.Last() got1 = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func TestIntWindow_Reset(t *testing.T) {
	type fields struct {
		data []int
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "resets the window",
			fields: fields{
				data: []int{1, 2, 3, 4, 5},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := New(tt.fields.data...)
			w.Reset()
			if len(w.Data()) != 0 {
				t.Errorf("IntWindow.reset() data = %v, want %v", w.Data(), []int{})
			}
		})
	}
}
