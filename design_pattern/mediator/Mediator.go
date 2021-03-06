package mediator

import "fmt"

// 中介者模式（mediator Pattern）
//     是用来降低多个对象和类之间的通信复杂性。
//  这种模式提供了一个中介类，该类通常处理不同类之间的通信，并支持松耦合，使代码易于维护

type ChatRoom struct {
	name string
}

func (cr *ChatRoom) SendMsg(msg string) {
	fmt.Println(cr.name + " : " + msg)
}
func (cr *ChatRoom) RegisterUser(u *User) {
	u.cr = cr
}

type User struct {
	name string
	cr   *ChatRoom
}

func (u *User) SendMsg(msg string) {
	if u.cr != nil {
		//这里调用的是ChatRoom->SendMsg
		u.cr.SendMsg(u.name + " : " + msg)
	}
}
