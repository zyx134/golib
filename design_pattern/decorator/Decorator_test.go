package decorator

import (
	"fmt"
	"testing"
)

func TestDecorator(t *testing.T) {
	var mb MessageBuilder
	mb = &BaseMessageBuilder{}
	fmt.Println(mb.Build("hello", "world"))
	mb = &QuoteMessageBuilderDecorator{Builder: mb}
}
func TestDecorator1(t *testing.T) {
	var mb MessageBuilder
	mb = &BaseMessageBuilder{}
	mb = &QuoteMessageBuilderDecorator{Builder: mb}
	fmt.Println(mb.Build("hello", "world"))
}
func TestDecorator2(t *testing.T) {
	var mb MessageBuilder
	mb = &BaseMessageBuilder{}
	mb = &BraceMessageBuilderDecorator{Builder: mb}
	fmt.Println(mb.Build("hello", "world"))
}
func TestDecorator3(t *testing.T) {
	f := LogDecorate(Double)
	f(5)
}
