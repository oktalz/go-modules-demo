package demo

import (
	"fmt"
	"testing"
)

func TestSplit(t *testing.T) {
	tests := map[string]struct {
		input  int
		result int
	}{
		"negative": {input: -1, result: 0},
		"zero":     {input: 0, result: 0},
		"10":       {input: 10, result: 10},
		"1000":     {input: 1000, result: 1000},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			rnd := RandomString(tc.input)
			if len(rnd) != tc.result {
				t.Errorf(fmt.Sprintf("error: has len [%d] expects [%d]", len(rnd), tc.result))
			}
		})
	}
}
