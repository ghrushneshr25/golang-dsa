/*
	Problem 4: Generate all the strings of length n drawn from 0...k−1

	Using Subtraction and Conquer Master theorem, we get: T(n)= O(݇k^n)
*/

package backtracking

import (
	"fmt"
	"strings"
)

type NCharStringFromK struct {
	PoolSize     int // refers to K
	Length       int
	Combinations []string
	Current      []int
}

func NewNCharStringFromK(poolSize, length int) *NCharStringFromK {
	field := &NCharStringFromK{}
	field.PoolSize = poolSize
	field.Length = length
	field.Current = make([]int, length)
	field.Combinations = []string{}
	return field
}

func (nchar *NCharStringFromK) Execute() {
	nchar.Recursion(nchar.Length - 1) // since indexing starts from zero
}

func (nchar *NCharStringFromK) Recursion(index int) {
	if index < 0 {
		nchar.generate()
		return
	}

	for j := 0; j < nchar.PoolSize; j++ {
		nchar.Current[index] = j
		nchar.Recursion(index - 1)
	}
}

func (nchar *NCharStringFromK) generate() {
	strs := make([]string, nchar.Length)
	for i, v := range nchar.Current {
		strs[i] = fmt.Sprintf("%d", v)
	}
	result := strings.Join(strs, "")
	if len(result) > 0 {
		nchar.Combinations = append(nchar.Combinations, result)
	}
}
