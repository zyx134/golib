package memento

import (
	"container/list"
	"fmt"
	"testing"
)

func TestMemento(t *testing.T) {
	storage := &Storage{list.New()}
	text := &Text{"hello world"}
	fmt.Println(text.Read())
	storage.PushBack(text.SaveToMemento())
	text.Write("nihao")
	fmt.Println(text.Read())
	storage.PushBack(text.SaveToMemento())
	text.Write("i know")
	fmt.Println(text.Read())

	//后退回滚
	mediator := storage.RPop()
	if mediator != nil {
		text.RestoreFromMemento(mediator.Value.(*Memento))
	}
	fmt.Println(text.Read())

	//后退回滚
	mediator = storage.RPop()
	if mediator != nil {
		text.RestoreFromMemento(mediator.Value.(*Memento))
	}
	fmt.Println(text.Read())

	//后退 已没有
	mediator = storage.RPop()
	if mediator != nil {
		text.RestoreFromMemento(mediator.Value.(*Memento))
	}
	fmt.Println(text.Read())
}
