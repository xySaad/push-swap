package utils

type stack struct {
	slice []int
}

func NewStack(slice []int) stack {
	return stack{slice}
}

func (s *stack) Pop() int {
	defer func() {
		s.slice = s.slice[1:]
	}()

	return s.slice[0]
}

func (s *stack) Push(new int) {
	s.slice = append([]int{new}, s.slice...)
}

func (s *stack) Swap() {
	first := s.slice[0]
	s.slice[0] = s.slice[1]
	s.slice[1] = first
}

func (s *stack) Rotate() {
	s.slice = append(s.slice, s.Pop())
}

func (s *stack) ReverseRotate() {
	lastIndex := len(s.slice) - 1
	s.Push(s.slice[lastIndex])
	s.slice = s.slice[:lastIndex+1]
}
