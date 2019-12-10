package numbers

import (
	"math/rand"
	"time"
)

func RandomFloat() (float64, error) {
	rand.Seed(time.Now().UnixNano())
	return rand.Float64(), nil
}
