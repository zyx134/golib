package abstractFactory

import "fmt"

//抽象工厂
type PenFactory interface {
	Create() Pen //生产笔
}

//铅笔工厂
type PencilFactory struct {
}

func (p PencilFactory) Create() Pen {
	return new(pencil)
}

//毛笔工厂
type BrushPenFactory struct {
}

func (p BrushPenFactory) Create() Pen {
	return new(brushPen)
}

//笔
type Pen interface {
	Write()
}

//铅笔
type pencil struct {
}

func (p *pencil) Write() {
	fmt.Println("铅笔")
}

//毛笔
type brushPen struct {
}

func (p *brushPen) Write() {
	fmt.Println("毛笔")
}
