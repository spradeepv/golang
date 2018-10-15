package ds

import "errors"

type Stack []interface{}

func (s *Stack) Push(value interface{}) {
	*s = append(*s, value)
}

func (s *Stack) Pop() (interface{}, error) {
	temp, err := s.Peek()
	if err != nil {
		return 0.0, err
	}

	*s = (*s)[0 : len(*s)-1]

	return temp, nil
}

func (s *Stack) Peek() (interface{}, error) {
	if len(*s) == 0 {
		return 0.0, errors.New("stack is empty")
	}

	return (*s)[len(*s)-1], nil
}

func (s *Stack) Len() int {
	return len(*s)
}
