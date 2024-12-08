package utils

type stack struct {
	slice []int
	Size  int
}

func NewStack(slice []int) stack {
	return stack{slice, len(slice)}
}

func (s *stack) Pop() int {
	s.Size--
	defer func() {
		s.slice = s.slice[1:]
	}()

	return s.slice[0]
}

func (s *stack) Push(new int) {
	s.Size++
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

func (s *stack) Peek() int {
	return s.slice[0]
}

func (s *stack) Top() int {
	return s.Peek()
}

func (s *stack) IsEmpty() bool {
	return len(s.slice) == 0
}
