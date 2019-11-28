package decorator

import "testing"

func TestConcreteDecorator_Operation(t *testing.T) {
	decorator := NewConcreteDecorator(&ConcreteComponent{})
	decorator.Operation()
}
