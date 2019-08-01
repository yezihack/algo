package simpleFactory

import "fmt"

//定义一个接口,让产品类实现他
type Factory interface {
	Drive()
}

//定义一个产品,实现Factory工厂
type Car struct {
}

func (f Car) Drive() {
	fmt.Println("我是一个小汽车")
}

type Ship struct {
}

func (f Ship) Drive() {
	fmt.Println("我是一条船")
}

type Aircraft struct {
}

func (f Aircraft) Drive() {
	fmt.Println("我是一架飞机")
}

//实现传入不同参数, 实现不同产品的对象
func New(name string) Factory {
	var f Factory
	switch name {
	case "car":
		f = new(Car)
	case "ship":
		f = new(Ship)
	case "aircraft":
		f = new(Aircraft)
	default:
		f = nil
	}
	return f
}
