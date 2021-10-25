package bridge

import "fmt"

//Implementor实现类接口
type Color interface {
	Use()
}

//实现类实现
type Red struct {
}

func (r Red) Use() {
	fmt.Printf("用沾有红色颜料")
}

type Green struct {
}

func (g Green) Use() {
	fmt.Printf("用沾有绿色颜料")
}

type Yellow struct {
}

func (y Yellow) Use() {
	fmt.Printf("用沾有黄色颜料")
}

//Abstraction抽象类
type Brush interface {
	Draw()
}

//抽象类实现
type BigBrush struct {
	Color
}

func (b BigBrush) Draw() {
	b.Color.Use()
	fmt.Println("的大毛笔画画")
}

type SmallBrush struct {
	Color
}

func (s SmallBrush) Draw() {
	s.Color.Use()
	fmt.Println("的小毛笔画画")
}

//定义工厂方法生产具体的毛笔
func NewBrush(t string, c Color) Brush {
	switch t {
	case "big":
		return &BigBrush{Color: c}
	case "small":
		return &SmallBrush{Color: c}
	default:
		return nil
	}
}
