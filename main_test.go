package demo

import (
	"fmt"
	"testing"
)

func TestSplit(t *testing.T) {
	tests := map[string]struct {
		input    int
		result   int
		hasError bool
	}{
		"negative": {input: -1, result: 0, hasError: true},
		"zero":     {input: 0, result: 0},
		"10":       {input: 10, result: 10},
		"1000":     {input: 1000, result: 1000},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			rnd, err := RandomString(tc.input)
			if tc.hasError && err == nil {
				t.Errorf("error: did not throw error")
				return
			}
			if len(rnd) != tc.result {
				t.Errorf(fmt.Sprintf("error: has len [%d] expects [%d]", len(rnd), tc.result))
			}
		})
	}
}
