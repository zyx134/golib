package adapter

import "testing"

func TestAdapterTest(t *testing.T) {
	var battery RechargeableBattery

	//不可充电电池A
	//这里必须实例化子类
	battery = AdapterNonToYes{NonRechargeableBattery: NonRechargeableA{}}
	battery.Use()
	battery.Charge()
	//NonRechargeableA using
	//AdapterNonToYes Charging

	//不可充电电池b
	//其实就是方法的提升
	//自己有的方法就用自己的，自己没有的就去子类找
	//但是这里不需要实例化子类
	battery = NonRechargeableB{}
	//battery = NonRechargeableB{RechargeableBatteryAbstract: RechargeableBatteryAbstract{}}
	battery.Use()
	battery.Charge()
	//NonRechargeableB using
	//RechargeableBatteryAbstract Charging
}
