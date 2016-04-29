package main

import (
	"fmt"
	"testing"
)

func TestFind(t *testing.T) {

	tests := []struct {
		ii   map[string]docList
		term string
		want []int
	}{
		{
			ii: map[string]docList{
				"ink": docList{1, 2},
			},
			term: "ink",
			want: docList{1, 2},
		}, {
			ii: map[string]docList{
				"ink": docList{1, 2},
			},
			term: "pink",
			want: nil,
		}, {
			ii: map[string]docList{
				"ink": docList{1, 2},
			},
			term: "pink",
			want: docList{},
		},
	}

	for _, test := range tests {
		got := find(test.ii, test.term)
		if !deepCompare(got, test.want) {
			t.Errorf("want: %v, got: %v\n", test.want, got)
		}
	}

}

func deepCompare(x []int, y []int) bool {

	if len(x) != len(y) {
		return false
	}

	for k := range x {
		if x[k] != y[k] {
			return false
		}
	}

	return true
}

func TestNil(t *testing.T) {
	var k []int
	k = nil
	fmt.Println(len(k))
}
