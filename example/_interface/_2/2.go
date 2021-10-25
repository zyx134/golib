package _2

import "fmt"

//多个类型实现同一接口
//Go语言中不同的类型还可以实现同一接口 首先我们定义一个Mover接口，它要求必须由一个move方法。

// Mover 接口
type Mover interface {
	move()
}

//例如狗可以动，汽车也可以动，可以使用如下代码实现这个关系：
type dog struct {
	name string
}
type car struct {
	name string
}

// dog类型实现Mover接口
func (d dog) move() {
	fmt.Printf("%s会跑\n", d.name)
}

// car类型实现Mover接口
func (c car) move() {
	fmt.Printf("%s速度70迈\n", c.name)
}
