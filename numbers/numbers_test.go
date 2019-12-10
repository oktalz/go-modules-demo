package numbers

import (
	"fmt"
	"testing"
)

func TestSplitInt(t *testing.T) {
	tests := map[string]struct {
		input    int
		result   int
		hasError bool
	}{
		"negative": {input: -1, result: 1, hasError: true},
		"zero":     {input: 0, result: 1, hasError: true},
		"10":       {input: 10, result: 10},
		"1000":     {input: 1000, result: 1000},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			rnd, err := RandomInt(tc.input)
			if tc.hasError && err == nil {
				t.Errorf("error: did not throw error")
				return
			}
			if rnd >= tc.result {
				t.Errorf(fmt.Sprintf("error: expects [%d] to be less than [%d]", rnd, tc.result))
			}
		})
	}
}

func TestSplitFloat(t *testing.T) {
	rnd, err := RandomFloat()
	if err != nil {
		t.Errorf("error: did not throw error")
		return
	}
	if float64(0) > rnd || rnd >= float64(1) {
		t.Errorf(fmt.Sprintf("error: expects [%f] to be less than [%f]", rnd, 1.0))
	}

}
