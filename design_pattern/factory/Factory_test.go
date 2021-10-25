package factory

import (
	"fmt"
	"testing"
)

func TestFactory(t *testing.T) {
	var f PenFactory
	var p pen
	p = f.ProducePencil()
	p.Write()
	p = f.ProduceBrushPen()
	p.Write()
	p = f.Produce("aaa")
	if p == nil {
		fmt.Println("没有这种笔")
	}
}
