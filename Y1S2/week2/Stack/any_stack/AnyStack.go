package any_stack

type AnyStack[T any] struct {
	elements []T
	length   int
}

// InitAnyStack 初始化新栈
func InitAnyStack[T any]() *AnyStack[T] {
	return &AnyStack[T]{
		elements: make([]T, 0),
		length:   0,
	}
}

// Push 新元素入栈
func (s *AnyStack[T]) Push(value T) {
	s.elements = append(s.elements, value)
	s.length++
}

// Pop 取出栈顶元素
func (s *AnyStack[T]) Pop() (T, bool) {
	if len(s.elements) == 0 {
		var zero T
		return zero, false
	}
	var zero T
	value := s.elements[s.length-1]
	s.elements[s.length-1] = zero
	s.length--
	return value, true
}

// Peek 查看栈顶元素
func (s *AnyStack[T]) Peek() (T, bool) {
	if len(s.elements) == 0 {
		var zero T
		return zero, false
	}
	return s.elements[s.length-1], true
}

// IsEmpty 检查栈是否为空，空为真，非空为假
func (s *AnyStack[T]) IsEmpty() bool {
	return s.length == 0
}

// CleanStack 清空栈中元素
func (s *AnyStack[T]) CleanStack() {
	s.elements = make([]T, 0)
	s.length = 0
}
