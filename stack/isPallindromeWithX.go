/*
	Problem-8
	Given an array of characters formed with a’s and b’s.
	The string is marked with special character X which represents the
	middle of the list (for example: ababa…ababXbabab…..baaa).
	Check whether the string is palindrome.


	Algorithm
	- 	Traverse the list till we encounter X as input element.
	- 	During the traversal push all the elements (until X) on to the stack.
	- 	For the second half of the list, compare each element’s content with top of the stack.
		If they are the same then pop the stack and go to the next element in the input list.
	- 	If they are not the same then the given string is not a palindrome.
	- 	Continue this process until the stack is empty or the string is not a palindrome.

*/

package stack

type IsPalindromeStack struct {
	stack  *LinkedListStack
	input  string
	Result bool
}

func NewIsPalindromeStack(input string) *IsPalindromeStack {
	return &IsPalindromeStack{
		stack:  NewLinkedListStack(),
		input:  input,
		Result: false,
	}
}

func (obj *IsPalindromeStack) IsPalindrome() bool {
	i := 0
	for index, char := range obj.input {
		i = index
		if char == 'X' {
			break
		}
		obj.stack.Push(char)
	}
	i++
	inputRune := []rune(obj.input)
	for ; i < len(obj.input); i++ {
		ch, err := obj.stack.Pop()
		if err != nil {
			return false
		}
		if ch.(rune) != inputRune[i] {
			return false
		}
	}
	return true
}

func (obj *IsPalindromeStack) Execute() {
	obj.Result = obj.IsPalindrome()
}
