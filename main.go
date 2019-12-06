package demo

import (
	"fmt"
	"math/rand"
	"time"
)

var chars = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

//RandomString returns random string of size n
func RandomString(n int) (string, error) {
	if n < 1 {
		return "", fmt.Errorf("index must be positive")
	}
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	size := len(chars)
	for i := range b {
		b[i] = chars[rand.Intn(size)]
	}
	return string(b), nil
}
