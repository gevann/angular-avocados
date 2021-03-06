package tree

import (
	"testing"
)

func TestTree_ToString(t *testing.T) {
	type fields struct {
		rightChild *Tree
		leftChild  *Tree
		parent     *Tree
		value      string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "It prints a string representation of the node",
			fields: fields{
				value: "01010",
			},
			want: "[data: 01010 | leaves: 0 | height: 0]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := &Tree{
				rightChild: tt.fields.rightChild,
				leftChild:  tt.fields.leftChild,
				parent:     tt.fields.parent,
				Value:      tt.fields.value,
			}
			if got := tree.ToString(); got != tt.want {
				t.Errorf("Tree.ToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTree_Insert(t *testing.T) {
	tree := &Tree{}

	inputs := []string{
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
	}

	for _, input := range inputs {
		tree.Insert(input)
	}

	wantedMaxByLeafCount := "10111"
	wantedMinByLeafCount := "01010"

	if tree.MaxByLeafCount().Value != wantedMaxByLeafCount {
		t.Errorf("Tree.MaxByLeafCount() = %v, want %v", tree.MaxByLeafCount().Value, wantedMaxByLeafCount)
	}

	if tree.MinByLeafCount().Value != wantedMinByLeafCount {
		t.Errorf("Tree.MinByLeafCount() = %v, want %v", tree.MinByLeafCount().Value, wantedMinByLeafCount)
	}

	wantedLeafCount := 12
	if tree.leafCount != wantedLeafCount {
		t.Errorf("Tree.height = %d, want %d", tree.leafCount, wantedLeafCount)
	}
}
