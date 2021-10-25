package mediator

import "testing"

func TestMediator(t *testing.T) {
	AUser := &User{name: "AUser"}
	BUser := &User{name: "BUser"}

	chatRoom := &ChatRoom{name: "chatRoom123456"}
	chatRoom.RegisterUser(AUser)
	chatRoom.RegisterUser(BUser)

	AUser.SendMsg("hello AUser")
	BUser.SendMsg("hello BUser")
}
