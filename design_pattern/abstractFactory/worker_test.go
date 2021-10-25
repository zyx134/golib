package abstractFactory

import (
	"fmt"
	"testing"
)

func TestWorker(t *testing.T) {
	var w WorkFactory
	fmt.Println(w)
	var _ WorkFactory = (*ProgrammerFactory)(nil)
	w = new(ProgrammerFactory)
	w.Creator().Work("我在ctrl+c ctrl+v")
	w = new(MeiShuFactory)
	w.Creator().Work("我在抠图")
}

func TestWorker1(t *testing.T) {
	//接口类型初始化只能var不能使用new
	//w := new(WorkFactory) //0xc000096510
	var w WorkFactory            //<nil>
	w1 := new(ProgrammerFactory) //&{}
	fmt.Println(w, w1)
	w = w1
	w.Creator().Work("我在ctrl+c ctrl+v")
}
