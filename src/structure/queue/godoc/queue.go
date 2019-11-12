package godoc

type Element interface{}

type Queue interface {
	Offer(e Element)
	Poll() Element
	Clear() bool
	Size() int
	IsEmpty() bool
}

type sliceEntry struct {
	element []Element
}

func NewQueue() *sliceEntry {
	return &sliceEntry{}
}

func (entry *sliceEntry) Offer(e Element) {
	entry.element = append(entry.element, e)
}

func (entry *sliceEntry) Poll() Element {
	if entry.IsEmpty() {
		return nil
	}
	first := entry.element[0]
	entry.element = entry.element[1:]
	return first
}

func (entry *sliceEntry) Clear() bool {
	if entry.IsEmpty() {
		return false
	}
	for i := 0; i < len(entry.element); i++ {
		entry.element[i] = nil
	}
	entry.element = nil
	return true
}
func (entry *sliceEntry) Size() int {
	return len(entry.element)
}

func (entry *sliceEntry) IsEmpty() bool {
	return entry.Size() == 0
}
