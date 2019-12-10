package strings

import (
	"fmt"
	"math/rand"
	"time"
)

var chars = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

//RandomStr returns random string of size n
func RandomStr(n int) (string, error) {
	if n < 1 {
		return "", fmt.Errorf("index must be > 0")
	}
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	size := len(chars)
	for i := range b {
		b[i] = chars[rand.Intn(size)]
	}
	return string(b), nil
}
