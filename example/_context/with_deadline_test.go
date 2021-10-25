package _context

import (
	"testing"
	"time"
)

func TestWithDeadline(t *testing.T) {

	d := time.Now().Add(time.Millisecond * 50)
	WithDeadline(d)
}
func TestWithDeadline2(t *testing.T) {

	d := time.Now().Add(time.Second * 2)
	WithDeadline(d)
}
