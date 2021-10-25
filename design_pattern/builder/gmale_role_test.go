package builder

import (
	"fmt"
	"testing"
)

func TestDirector_Create(t *testing.T) {
	var builder Builder = &CharacterBuilder{}
	var director = &Director{builder: builder}
	var character = director.Create("baba", "AK47")
	fmt.Println(character.GetName(), character.GetArm())
}
func TestDirector_Create1(t *testing.T) {
	var director = &Director{builder: &CharacterBuilder{}}
	var character = director.Create("baba", "AK47")
	fmt.Println(character.GetName(), character.GetArm())
	//需要调用Create方法，需要一个Director实例
	//Director实例需要一个实现了Builder接口的数据类型
	//在这里就是CharacterBuilder
}
