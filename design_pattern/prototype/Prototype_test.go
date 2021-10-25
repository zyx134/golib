package prototype

import (
	"fmt"
	"testing"
)

func TestPrototype(t *testing.T) {
	p := make(ContextPool)
	p.AddContext("a", &Context{})
	p.GetContext("a")
}
func TestPrototype1(t *testing.T) {
	c := &Context1{
		Uri:     "a",
		Context: Context{},
	}
	fmt.Println(c.String())
}
