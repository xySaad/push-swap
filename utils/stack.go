package utils

type stack struct {
	slice          []int
	Size, Max, Min int
}

func NewStack(input []int) stack {
	var reversed stack
	for i := len(input) - 1; i >= 0; i-- {
		reversed.push(input[i])
	}

	return reversed
}

func (s *stack) pop() int {
	s.Size--
	defer func() {
		s.slice = s.slice[:s.Size]
	}()

	return s.slice[s.Size]
}

func (s *stack) push(new int) {
	if new > s.Max {
		s.Max = new
	}
	s.Size++
	s.slice = append(s.slice, new)
}

func (s *stack) swap() {
	s.slice[s.Size-1], s.slice[s.Size-2] = s.slice[s.Size-2], s.slice[s.Size-1]
}

func (s *stack) rotate() {
	s.slice = append(s.slice[s.Size-1:s.Size], s.slice[:s.Size-1]...)
}

func (s *stack) reverseRotate() {
	s.slice = append(s.slice[1:], s.slice[0])
}

func (s *stack) Peek() int {
	return s.slice[s.Size-1]
}

func (s *stack) IsEmpty() bool {
	return len(s.slice) == 0
}
