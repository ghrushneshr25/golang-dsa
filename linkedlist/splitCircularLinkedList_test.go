package linkedlist

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type SplitCircularLinkedListTestSuite struct {
	suite.Suite
}

func TestSplitCircularLinkedList(t *testing.T) {
	suite.Run(t, new(SplitCircularLinkedListTestSuite))
}

func (scll *SplitCircularLinkedListTestSuite) TestCase() {
	tests := []struct {
		head    []int
		listOne []int
		listTwo []int
	}{
		{
			head:    []int{1, 2, 3, 4, 5, 6, 7, 8},
			listOne: []int{1, 2, 3, 4},
			listTwo: []int{5, 6, 7, 8},
		},
		{
			head:    []int{1, 2, 3},
			listOne: []int{1, 2},
			listTwo: []int{3},
		},
		{
			head:    []int{1},
			listOne: []int{1},
			listTwo: []int{},
		},
		{
			head:    []int{},
			listOne: []int{},
			listTwo: []int{},
		},
	}

	for index, test := range tests {
		testName := fmt.Sprintf("Test Case :%v", index+1)
		scll.Run(testName, func() {
			mtslll := NewSplitCircularLinkedList(test.head)
			mtslll.Execute()
			scll.Equal(test.listOne, mtslll.ResultOne)
			scll.Equal(test.listTwo, mtslll.ResultTwo)
		})
	}
}
