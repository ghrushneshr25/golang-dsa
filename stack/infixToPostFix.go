/*

Problem-2 Infix to postfix conversion algorithm using stack.

Solution: Before discussing the algorithm, first let us see the definitions of infix, prefix and postfix expressions.
	Infix: An infix expression is a single letter, or an operator, proceeded by one infix string and followed by another Infix string.
		A
		A+B
		(A+B)+ (C-D)

	Prefix:  A prefix expression is a single letter, or an operator, followed by two prefix strings.
			Every prefix string longer than a single variable contains an operator, first operand and second operand.
		A
		+AB
		++AB-CD

	Postfix: A postfix expression (also called Reverse Polish Notation) is a single letter or an operator, preceded by two postfix strings.
	Every postfix string longer than a single variable contains first and second operands followed by an operator.

		A
 		AB+
 		AB+CD-+

	Algorithm:
		a) Create a stack
		b) for each character t in the input stream{
				if(t is an operand)
					output t
				else if(t is a right parenthesis){
					Pop and output tokens until a left parenthesis is popped (but not output)
				}
				else // t is an operator or left parenthesis{
					pop and output tokens until one of lower priority than t is encountered or a left  parenthesis is encountered or the
					stack is empty
					Push t
				}
			}
		c) pop and output tokens until the stack is empty
*/

package stack

type InfixToPostFix struct {
	stack  *LinkedListStack
	Input  string
	Result string
}

func NewInfixToPostFix(input string) *InfixToPostFix {
	return &InfixToPostFix{
		stack: NewLinkedListStack(),
		Input: input,
	}
}

func (itp *InfixToPostFix) Convert() string {
	if len(itp.Input) == 0 {
		return ""
	}

	var result string
	for _, c := range itp.Input {
		if itp.isOperator(c) {
			for !itp.stack.IsEmpty() {
				top, _ := itp.stack.Peek()
				if itp.precedence(top.(rune)) < itp.precedence(c) {
					break
				}
				result += string(top.(rune))
				itp.stack.Pop()
			}
			itp.stack.Push(c)
		} else if c == '(' {
			itp.stack.Push(c)
		} else if c == ')' {
			for !itp.stack.IsEmpty() {
				top, _ := itp.stack.Peek()
				if top.(rune) == '(' {
					itp.stack.Pop()
					break
				}
				result += string(top.(rune))
				itp.stack.Pop()
			}
		} else if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') {
			result += string(c)
		}
	}

	for !itp.stack.IsEmpty() {
		top, _ := itp.stack.Pop()
		result += string(top.(rune))
	}
	
	return result
}

func (itp *InfixToPostFix) precedence(op rune) int {
	switch op {
	case '+', '-':
		return 1
	case '*', '/':
		return 2
	case '^':
		return 3
	}
	return 0
}

func (itp *InfixToPostFix) isOperator(c rune) bool {
	switch c {
	case '+', '-', '*', '/', '^':
		return true
	}
	return false
}
