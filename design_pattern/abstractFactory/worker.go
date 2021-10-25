package abstractFactory

import "fmt"

/**
抽象类接口
*/
type Worker interface {
	Work(task string)
}

/**
抽象工厂类

go里面实现了接口的方法
所以下面
ProgrammerFactory
MeiShuFactory
都是实现了WorkFactory
*/

//抽象工厂
type WorkFactory interface {
	Creator() Worker
}

/**
具体类
*/
//程序员
type Programmer struct {
}

func (p *Programmer) Work(task string) {
	fmt.Println("程序员正在写代码 ", task)
}

/**
具体工厂类
*/
//程序员工厂
type ProgrammerFactory struct {
}

func (p *ProgrammerFactory) Creator() Worker {
	return new(Programmer)
}

/**
具体类
*/
//美术
type MeiShu struct {
}

func (m *MeiShu) Work(task string) {
	fmt.Println("美术正在切图 ", task)
}

/**
具体工厂类
*/
//美术工厂
type MeiShuFactory struct {
}

func (m *MeiShuFactory) Creator() Worker {
	return new(MeiShu)
}
