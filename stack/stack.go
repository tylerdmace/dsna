package stack

import "errors"

type Stack struct {
	S []int
}

func (s *Stack) Pop() (int, error) {
	if len(s.S) == 0 {
		return 0, errors.New("stack is empty!")
	}

	i := s.S[len(s.S)-1]
	s.S = s.S[:len(s.S)-1]

	return i, nil
}

func (s *Stack) Push(i int) {
	s.S = append(s.S, i)
}

func (s *Stack) IsEmpty() bool {
	return len(s.S) == 0
}

func (s *Stack) Size() int {
	return len(s.S)
}
