package _4

import "fmt"

//接口嵌套
//接口与接口间可以通过嵌套创造出新的接口。

//Sayer接口
type Sayer interface {
	say()
}

//Mover接口
type Mover interface {
	move()
}

// 接口嵌套
type animal interface {
	Sayer
	Mover
}

//嵌套得到的接口的使用与普通接口一样，这里我们让cat实现animal接口：
type cat struct {
	name string
}

func (c cat) say() {
	fmt.Println("喵喵喵")
}
func (c cat) move() {
	fmt.Println("速度10码")
}
