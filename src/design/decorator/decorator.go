package decorator

import "fmt"

type Component interface {
	Operation()
}

type ConcreteComponent struct{}

func (component *ConcreteComponent) Operation() {
	fmt.Println("this is concrete component")
}

type Decorator struct {
	component Component
}

func NewDecorator(component Component) *Decorator {
	return &Decorator{component: component}
}

func (decorator *Decorator) Operation() {
	decorator.component.Operation()
}

type ConcreteDecorator struct {
	Decorator
}

func NewConcreteDecorator(component Component) *ConcreteDecorator {
	return &ConcreteDecorator{Decorator{component: component}}
}
func (concreteDecorator *ConcreteDecorator) Operation() {
	fmt.Println("this is concrete decorator , ready?")
	concreteDecorator.component.Operation()
}
