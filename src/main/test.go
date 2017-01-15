package main

import (
	"fmt"
	"kernel/channel/socket"
	"kernel/channel"
)

type Base struct {
	Name string
}



func (b *Base) SetName(name string) {
	b.Name = name
}

func (b *Base) GetName() string {
	return b.Name
}

// 组合，实现继承
type Child struct {
	base Base  // 这里保存的是Base类型
}

// 重写GetName方法
func (c *Child) GetName() string {
	c.base.SetName("modify...")
	return c.base.GetName()
}

// 实现继承，但需要外部提供一个Base的实例
type Child2 struct {
	base *Base  // 这里是指针
}

func (c *Child2) GetName() string {
	c.base.SetName("canuser?")
	return c.base.GetName()
}



type _Base interface {
	Test()
}

type Base1 struct {


}

func (base1 Base1)Test()  {
	fmt.Println("base1")
}

type Base2 struct {
	Base1
}

func (base2 Base2)Test()  {
	fmt.Println("base2")
}

func main() {

	var channel channel.IChannel =&socket.TCPServerSocketChannel{}

	channel.(socket.IServerSocketChannel).DoBindAndAccept(90)








var _base Base2=Base2{}
	_base.Base1.Test()








	c := new(Child)
	c.base.SetName("world")
	fmt.Println(c.GetName())

	c2 := new(Child2)
	c2.base = new(Base)  // 因为Child2里面的Base是指针类型，所以必须提供一个Base的实例
	fmt.Println(c2.GetName())
}


func test() Base{
    return Base{Name:"kf"}
}