package generics

type StackOfInts = Stack[int]
type StackOfStrings = Stack[string]
type Stack[T comparable] struct {
	values []T
}

func (s *Stack[T]) Push(value T) {
	s.values = append(s.values, value)
}
func (s *Stack[T]) IsEmpty() bool {
	return len(s.values) == 0
}
func (s *Stack[T]) Pop() (value T, ok bool) {
	if s.IsEmpty() {
		//value = *new(T)
		//var value T
		ok = false
		return
	}
	index := len(s.values) - 1
	value = s.values[index]
	s.values = s.values[:index]
	ok = true
	return
}

func NewStack[T comparable]() *Stack[T] {
	return new(Stack[T])
}
