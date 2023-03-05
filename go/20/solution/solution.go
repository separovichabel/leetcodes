package solution

var pairs map[rune]rune = map[rune]rune{
	'[': ']',
	'(': ')',
	'{': '}',
}

func IsValid(s string) bool {
	if len(s) == 0 {
		return true
	}

	stack := Stack{}

	for _, char := range s {
		if stack.IsEmpty() {
			stack.Push(char)
		} else if pairs[stack.Tail()] == char {
			stack.Pop()
		} else {
			stack.Push(char)
		}
	}
	return stack.IsEmpty()
}

type Stack struct {
	items []rune
}

func (s *Stack) Pop() {
	if s.IsEmpty() {
		return
	}

	s.items = s.items[:len(s.items)-1]
}

func (s *Stack) Push(item rune) {
	s.items = append(s.items, item)
}

func (s *Stack) Tail() rune {
	return s.items[len(s.items)-1]
}

func (s *Stack) IsEmpty() bool {
	if len(s.items) == 0 {
		return true
	}
	return false
}
