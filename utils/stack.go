package utils

type stack struct {
	slice []int
	Size  int
}

func NewStack(input []int) stack {
	var reversed []int
	for i := len(input) - 1; i >= 0; i-- {
		reversed = append(reversed, input[i])
	}

	return stack{reversed, len(input)}
}

func (s *stack) pop() int {
	s.Size--
	defer func() {
		s.slice = s.slice[:s.Size]
	}()

	return s.slice[s.Size]
}

func (s *stack) push(new int) {
	s.Size++
	s.slice = append(s.slice, new)
}

func (s *stack) swap() {
	first := s.slice[s.Size-1]
	s.slice[s.Size-1] = s.slice[s.Size-2]
	s.slice[s.Size-2] = first
}

func (s *stack) rotate() {
	s.slice = append(s.slice, s.pop())
}

func (s *stack) reverseRotate() {
	lastIndex := len(s.slice) - 1
	s.push(s.slice[lastIndex])
	s.slice = s.slice[:lastIndex+1]
}

func (s *stack) Peek() int {
	return s.slice[0]
}

func (s *stack) IsEmpty() bool {
	return len(s.slice) == 0
}
