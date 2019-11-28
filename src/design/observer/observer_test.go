package observer

import (
	"fmt"
	"testing"
)

func TestConcreteSubject_Notify(t *testing.T) {
	subject := &ConcreteSubject{}
	first := NewFirstObserver(subject)
	second := NewSecondObserver(subject)

	subject.SetState(8)

	fmt.Println("after notify")
	first.Update()
	second.Update()
}
