package abstractFactory

import "testing"

//判断接口是否实现
//var _ 接口 = (*实现接口的结构)(nil)
func TestAbstractFactory(t *testing.T) {
	var a PenFactory
	var _ PenFactory = (*PencilFactory)(nil)
	a = new(PencilFactory)
	a.Create().Write()
	a = new(BrushPenFactory)
	a.Create().Write()
}
