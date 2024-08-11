package stack

type stackItem struct {
	value int
	next  *stackItem
}
type Stack struct {
	head *stackItem
}

func NewStack() Stack {
	return Stack{}
}

func (s *Stack) Push(value int) {
	s.head = &stackItem{
		value: value,
		next:  s.head,
	}
}

func (s *Stack) Pop() (value int, Ok bool) {
	if s.head == nil {
		return 0, false
	}

	value = s.head.value
	s.head = s.head.next

	return value, true
}

func (s *Stack) Top() (value int, Ok bool) {
	if s.head == nil {
		return 0, false
	}
	return s.head.value, true
}
