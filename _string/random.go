package _string

import (
	"math/rand"
	"time"
)

var letters = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func Random(n int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, n)
	for i := range b {
		bb := r.Intn(26) + 65
		b[i] = byte(bb)
	}
	return string(b)
}
func Random1(n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, n)
	for i := range b {
		bb := rand.Intn(62)
		b[i] = letters[bb]
	}
	return string(b)
}
