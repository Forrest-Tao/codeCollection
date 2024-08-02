package decoratorPattern

import "fmt"

type component interface {
	sendMsg() string
}

type ConcreteComponent struct {
	msg string
}

func (c *ConcreteComponent) sendMsg() string {
	return c.msg
}

//========define Decorator==============

type Decorator struct {
	component component
}

func (d *Decorator) sendMsg() string {
	return d.component.sendMsg()
}

//  ========implement Decorator==============

type ConcreteDecoratorA struct {
	Decorator
}

func (d *ConcreteDecoratorA) SendMessage() string {
	return fmt.Sprintf("[Encrypted] %s", d.component.sendMsg())
}

// ConcreteDecoratorB 实现 Decorator，并添加额外的行为
type ConcreteDecoratorB struct {
	Decorator
}

func (d *ConcreteDecoratorB) SendMessage() string {
	return fmt.Sprintf("[Compressed] %s", d.component.sendMsg())
}
