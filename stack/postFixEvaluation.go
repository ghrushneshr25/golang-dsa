/*
	Problem-4: Discuss postfix evaluation using stacks? 
	Algorithm: 
		1  	Scan the Postfix string from left to right. 
		2  	Initialize an empty stack. 
		3  	Repeat steps 4 and 5 till all the characters are scanned. 
		4  	If the scanned character is an operand, push it onto the stack.  
		5  	If the scanned character is an operator, and if the operator is a unary operator, then pop an element from the stack. 
			If the operator is a binary operator, then pop two elements from the stack. After popping the elements, 
			apply the operator to those popped elements. Let the result of this operation be retVal onto the stack.  
		6  	After all characters are scanned, we will have only one element in the stack. 
		7  	Return top of the stack as result. 

*/

package stack

type PostFixEvaluation struct {
	stack      *LinkedListStack
	expression string
	Result     int
}

func NewPostFixEvaluation(expression string) *PostFixEvaluation {
	return &PostFixEvaluation{
		stack:      NewLinkedListStack(),
		expression: expression,
	}
}

func (pfe *PostFixEvaluation) Evaluate() int {
	for _, c := range pfe.expression {
		if c == '+' {
			op1, _ := pfe.stack.Pop()
			op2, _ := pfe.stack.Pop()
			res := op1.(int) + op2.(int)
			pfe.stack.Push(res)
		} else if c == '-' {
			op1, _ := pfe.stack.Pop()
			op2, _ := pfe.stack.Pop()
			res := op2.(int) - op1.(int)
			pfe.stack.Push(res)
		} else if c == '*' {
			op1, _ := pfe.stack.Pop()
			op2, _ := pfe.stack.Pop()
			res := op1.(int) * op2.(int)
			pfe.stack.Push(res)
		} else if c == '/' {
			op1, _ := pfe.stack.Pop()
			op2, _ := pfe.stack.Pop()
			res := op2.(int) / op1.(int)
			pfe.stack.Push(res)
		} else {
			pfe.stack.Push(int(c - '0'))
		}
	}

	result, _ := pfe.stack.Pop()
	return result.(int)
}

func (pfe *PostFixEvaluation) Execute() {
	pfe.Result = pfe.Evaluate()
}
