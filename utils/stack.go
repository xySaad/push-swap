package utils

type stack struct {
	slice     []int
	Size, Min int
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
	top := s.slice[s.Size]
	if top < s.Min {
		oldMin := s.Min
		s.Min = 2*s.Min - top
		return oldMin
	}
	return top
}

func (s *stack) push(x int) {
	if s.Size == 0 {
		s.Min = x
		s.slice = append(s.slice, x)
	} else {
		if x < s.Min {
			s.slice = append(s.slice, 2*x-s.Min)
			s.Min = x
		} else {
			s.slice = append(s.slice, x)
		}
	}
	s.Size++
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
