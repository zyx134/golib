package singleton

import "testing"

func TestSingleton1(t *testing.T) {
	s := NewSingletonInstance()
	s.SaySomething()
}
func TestSingleton2(t *testing.T) {
	Singleton2SaySomething()
}
