/*
	Problem-21
	Given a string containing D D′D and D D′D where D indicates a push operation and D indicates
	a pop operation, and with the stack initially empty, formulate a rule to check whether a given
	string D of operations is admissible or not?
*/

package stack

type AdmissibleString struct {
	stack  *LinkedListStack
	input  string
	Result bool
}

func NewAdmissibleString(input string) *AdmissibleString {
	return &AdmissibleString{
		stack: NewLinkedListStack(),
		input: input,
	}
}

func (s *AdmissibleString) IsAdmissible() bool {
	for _, char := range s.input {
		if char == 'S' {
			s.stack.Push(char)
		} else {
			_, err := s.stack.Pop()
			if err != nil {
				return false
			}
		}
	}
	return true
}

func (s *AdmissibleString) Execute() {
	s.Result = s.IsAdmissible()
}
