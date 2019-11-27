package strategy

import "fmt"

type Strategy interface {
	Operate()
}

type ImplOne struct{}

func (impl *ImplOne) Operate() {
	fmt.Println("this is the first implementation")
}

type ImplTwo struct{}

func (impl *ImplTwo) Operate() {
	fmt.Println("this is the second implementation")
}

type Context struct {
	strategy Strategy
}

func NewContext(strategy Strategy) *Context {
	return &Context{strategy: strategy}
}

func (context *Context) SetStrategy(strategy Strategy) {
	context.strategy = strategy
}

func (context *Context) ExecuteStrategy() {
	context.strategy.Operate()
}
