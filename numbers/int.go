package numbers

import (
	"fmt"
	"math/rand"
	"time"
)

func RandomInt(n int) (int, error) {
	if n < 1 {
		return 0, fmt.Errorf("index must be > 0")
	}
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(n), nil
}
