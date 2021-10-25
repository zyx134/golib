package builder

import (
	"fmt"
	"testing"
)

func TestMealBuilder_MealOne(t *testing.T) {
	m := MealBuilder{}
	m1 := m.MealOne()
	//for _, v := range *m1 {
	//	fmt.Println(v.Name())
	//	fmt.Println(v.Price())
	//	fmt.Println(v.Category())
	//}
	m1.AddItem(new(Cola))
	fmt.Println(m1.ShowItems())
	fmt.Println(m1.GetCost())
}
