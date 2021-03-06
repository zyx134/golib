package memento

import (
	"container/list"
)

// 备忘录模式（Memento pattern）
//   在不破坏封装性的前提下，捕获一个对象的内部状态，并在该对象之外保存这个状态。
//   这样以后就可将该对象恢复到原先保存的状态。

// 相关模式：原型模式

// 文本编辑
type Text struct {
	Value string
}

// 写
func (t *Text) Write(value string) {
	t.Value = value
}

// 读取
func (t *Text) Read() string {
	return t.Value
}

// 备忘
func (t *Text) SaveToMemento() *Memento {
	return &Memento{Value: t.Value}
}

// 从备忘恢复
func (t *Text) RestoreFromMemento(m *Memento) {
	if m != nil {
		t.Value = m.Value
	}
	return
}

// 备忘结构
type Memento struct {
	Value string
}

// 管理备忘记录
type Storage struct {
	//内嵌一个列表，这样列表的方法会提升
	*list.List
}

// Back returns the last element of list l or nil.
// and remove form list
func (s *Storage) RPop() *list.Element {
	ele := s.Back()
	if ele != nil {
		s.Remove(ele)
	}
	return ele
}
