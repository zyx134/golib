package chain_of_responsobolity

import "testing"

func TestChainOfResponsibilityTest(t *testing.T) {
	var i IScreenEventHandler
	i = &AbsScreenEventHandler{}
	var home = &HomeScreenEventHandler{}
	var user = &UserScreenEventHandler{}
	home.SetNextHandler(user)
	i.SetNextHandler(home)

	a := &ScreenEvent{Type: "HomeClick"}
	i.Handle(a)

	a = &ScreenEvent{Type: "Null"}
	i.Handle(a)
}
