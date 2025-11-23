package main

import "fmt"

func main() {
	fmt.Println(isValid("()"))     // true
	fmt.Println(isValid("()[]{}")) // true
	fmt.Println(isValid("(]"))     // false
	fmt.Println(isValid("([])"))   // true
	fmt.Println(isValid("([)]"))   // false
}

func isValid(s string) bool {
	stack := newStack()
	pairs := map[rune]rune{
		CloseParenthess: OpenParenthess,
		CloseBracket:    OpenBracket,
		CloseBrace:      OpenBrace,
	}

	for _, char := range s {
		if char == OpenParenthess || char == OpenBracket || char == OpenBrace {
			stack.push(char)
		} else {
			if stack.isEmpty() {
				return false
			}

			top := stack.pop()
			if pairs[char] != *top {
				return false
			}
		}
	}

	return stack.isEmpty()
}

const (
	OpenParenthess  = '('
	CloseParenthess = ')'
	OpenBracket     = '['
	CloseBracket    = ']'
	OpenBrace       = '{'
	CloseBrace      = '}'
)

type Stack struct {
	Container []rune
}

func newStack() *Stack {
	return &Stack{
		Container: []rune{},
	}
}

func (s *Stack) push(char rune) {
	s.Container = append(s.Container, char)
}

func (s *Stack) pop() *rune {
	volume := len(s.Container)
	if volume > 0 {
		top := s.Container[volume-1]
		s.Container = s.Container[:volume-1]
		return &top
	}
	return nil
}

func (s *Stack) isEmpty() bool {
	return len(s.Container) == 0
}

// func (s *Stack) peek() (rune, bool) { // 新增：查看 top 但不移除
// 	if len(s.Container) > 0 {
// 		return s.Container[len(s.Container)-1], true
// 	}
// 	return 0, false
// }
