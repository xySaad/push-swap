package utils

type stack struct {
	track          bool
	slice          []int
	Size, Min, Max int
}

func NewStack(input []int, track bool) stack {
	var reversed stack
	reversed.track = track

	for i := len(input) - 1; i >= 0; i-- {
		reversed.push(input[i])
	}
	return reversed
}

func (s *stack) pop() int {
	s.Size--
	top := s.slice[s.Size]
	s.slice = s.slice[:s.Size]

	if top < s.Min && s.track {
		oldMin := s.Min
		s.Min = 2*s.Min - top
		return oldMin
	} else if top > s.Max && s.track {
		oldMax := s.Max
		s.Min = 2*s.Max - top
		return oldMax
	}

	return top
}

func (s *stack) push(x int) {
	if s.Size == 0 {
		s.Min = x
		s.Max = x
		s.slice = append(s.slice, x)
	} else {
		if x < s.Min && s.track {
			s.slice = append(s.slice, 2*x-s.Min)
			s.Min = x
		} else if x > s.Max && s.track {
			s.slice = append(s.slice, 2*x-s.Max)
			s.Max = x
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

func (s *stack) PeekValue() int {
	top := s.slice[s.Size-1]
	if top < s.Min {
		return s.Min
	} else if top > s.Max {
		return s.Max
	}
	return top
}

func (s *stack) IsEmpty() bool {
	return len(s.slice) == 0
}
