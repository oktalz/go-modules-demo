package demo

import (
	"math/rand"
	"time"
)

var chars = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

//RandomString returns random string of size n
func RandomString(n int) string {
	if n < 1 {
		return ""
	}
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	size := len(chars)
	for i := range b {
		b[i] = chars[rand.Intn(size)]
	}
	return string(b)
}
