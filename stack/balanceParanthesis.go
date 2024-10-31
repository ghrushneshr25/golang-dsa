/*
Problem-1 Discuss how stacks can be used for checking balancing of symbols/parentheses.

Solution: Stacks can be used to check whether the given expression has balanced symbols. This algorithm is very useful in compilers. Each
time the parser reads one character at a time. If the character is an opening delimiter such as (, {, or [- then it is written to the stack. When a
closing delimiter is encountered like ), }, or ]- the stack is popped.

The opening and closing delimiters are then compared. If they match, the parsing of the string continues. If they do not match, the parser
indicates that there is an error on the line.

A linear-time O(n) algorithm based on stack can be given as:
	a) Create a stack.
	b) while (end of input is not reached) {
		1) If the character read is not a symbol to be balanced, ignore it.
		2) If the character is an opening symbol like (, [, {, push it onto the stack
		3) If it is a closing symbol like ),],}, then if the stack is empty report an error. Otherwise pop the stack.
		4) If the symbol popped is not the corresponding opening symbol, report an error.
	}
	c) At end of input, if the stack is not empty report an error

*/

package stack

type BalanceParanthesis struct {
	stack  *LinkedListStack
	Input  string
	Result bool
}

func NewBalanceParanthesis(input string) *BalanceParanthesis {
	return &BalanceParanthesis{
		stack: NewLinkedListStack(),
		Input: input,
	}
}

func (bp *BalanceParanthesis) IsBalanced() bool {

	pairs := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	if len(bp.Input) == 0 {
		return true
	}

	for _, c := range bp.Input {
		if c == '(' || c == '{' || c == '[' {
			bp.stack.Push(c)
		} else if c == ')' || c == '}' || c == ']' {
			pair, err := bp.stack.Peek()
			if err != nil {
				return false
			}
			if pair.(rune) != pairs[c] {
				return false
			}
			bp.stack.Pop()
		}
	}

	return bp.stack.IsEmpty()
}

func (bp *BalanceParanthesis) Execute() {
	bp.Result = bp.IsBalanced()
}
