package observer

import "fmt"

type Observer interface {
	Update()
}

type Subject interface {
	GetState() int
	SetState(int)
	Attach(Observer)
	Notify()
}

type FirstObserver struct {
	subject Subject
}

func NewFirstObserver(subject Subject) *FirstObserver {
	observer := &FirstObserver{subject: subject}
	subject.Attach(observer)
	return observer
}
func (observer *FirstObserver) Update() {
	fmt.Printf("this is the first observer , state from Subject is %d\n", observer.subject.GetState())
}

type SecondObserver struct {
	subject Subject
}

func NewSecondObserver(subject Subject) *SecondObserver {
	observer := &SecondObserver{subject: subject}
	subject.Attach(observer)
	return observer
}
func (observer *SecondObserver) Update() {
	fmt.Printf("this is the second observer , state from Subject is %d\n", observer.subject.GetState())
}

type ConcreteSubject struct {
	observers []Observer
	state     int
}

func (subject *ConcreteSubject) GetState() int {
	return subject.state
}

func (subject *ConcreteSubject) SetState(state int) {
	subject.state = state
	subject.Notify()
}

func (subject *ConcreteSubject) Attach(observer Observer) {
	subject.observers = append(subject.observers, observer)
}

func (subject *ConcreteSubject) Notify() {
	for _, observer := range subject.observers {
		observer.Update()
	}
}
