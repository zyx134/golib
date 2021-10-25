package builder

/**
抽象方法
*/
type Item interface {
	Price() float32   //价钱
	Name() string     //名称
	Category() string //类别
}

type Meal []Item

func (m *Meal) AddItem(item ...Item) {
	*m = append(*m, item...)
}

func (m Meal) GetCost() (cost float32) {
	for _, val := range m {
		cost += val.Price()
	}
	return
}
func (m Meal) ShowItems() (msg string) {
	for _, val := range m {
		msg += "Category：" + val.Category() + " Name:" + val.Name() + "\n"
	}
	return
}

/**
规范实例化
*/
//建造者
type MealBuilder struct {
}

func (MealBuilder) MealOne() (meal *Meal) {
	meal = new(Meal)
	meal.AddItem(new(FriedChicken), new(Beer))
	return
}

/**
复杂对象
*/
//食物
type Food struct {
}

//饮料
type Drink struct {
}

//汉堡
type Hamburger struct {
	Food
}

//炸鸡
type FriedChicken struct {
	Food
}

//可乐
type Cola struct {
	Drink
}

//啤酒
type Beer struct {
	Drink
}

/**
具体实现
*/
func (f Food) Price() float32 {
	return 0.0
}
func (f Food) Name() string {
	return ""
}
func (f Food) Category() string {
	return "food"
}

func (d Drink) Price() float32 {
	return 0.0
}
func (d Drink) Name() string {
	return ""
}
func (d Drink) Category() string {
	return "drink"
}

func (h Hamburger) Price() float32 {
	return 12.00
}
func (h Hamburger) Name() string {
	return "Hamburger"
}

func (f FriedChicken) Price() float32 {
	return 18.00
}
func (f FriedChicken) Name() string {
	return "FriedChicken"
}

func (c Cola) Price() float32 {
	return 3.00
}
func (c Cola) Name() string {
	return "Cola"
}

func (b Beer) Price() float32 {
	return 5.00
}
func (b Beer) Name() string {
	return "Beer"
}
