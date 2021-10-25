package builder

//复杂对象
//创建复杂对象
type Character struct {
	name string
	arm  string
}

func (c *Character) SetName(name string) {
	c.name = name
}
func (c *Character) SetArm(arm string) {
	c.arm = arm
}
func (c Character) GetName() string {
	return c.name
}
func (c Character) GetArm() string {
	return c.arm
}

//抽象方法
//实现抽象建造者接口
type Builder interface {
	SetName(string) Builder
	SetArm(string) Builder
	Build() *Character
}

//具体实现
//创建具体的创建者类
type CharacterBuilder struct {
	character *Character
}

func (c *CharacterBuilder) SetName(name string) Builder {
	if c.character == nil {
		c.character = &Character{}
	}
	c.character.SetName(name)
	return c
}
func (c *CharacterBuilder) SetArm(arm string) Builder {
	if c.character == nil {
		c.character = &Character{}
	}
	c.character.SetArm(arm)
	return c
}
func (c *CharacterBuilder) Build() *Character {
	return c.character
}

//规范实例化
//Director的实现
type Director struct {
	builder Builder
}

func (d Director) Create(name, arm string) *Character {
	return d.builder.SetName(name).SetArm(arm).Build()
}
