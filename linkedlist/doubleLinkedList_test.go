package linkedlist

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type DoubleLinkedListTestSuite struct {
	suite.Suite
	head *DoubleLinkedList
}

func TestDoubleLinkedList(t *testing.T) {
	suite.Run(t, &DoubleLinkedListTestSuite{})
}

func (suite *DoubleLinkedListTestSuite) SetupTest() {
	suite.head = nil
}

func (suite *DoubleLinkedListTestSuite) TestNewDoubleLinkedList() {
	tests := []struct {
		data     int
		next     *DoubleLinkedList
		prev     *DoubleLinkedList
		expected *DoubleLinkedList
	}{
		{
			data: 1,
			next: nil,
			prev: nil,
			expected: &DoubleLinkedList{
				data: 1,
				next: nil,
				prev: nil,
			},
		},
		{
			data: 2,
			next: &DoubleLinkedList{data: 3},
			prev: &DoubleLinkedList{data: 1},
			expected: &DoubleLinkedList{
				data: 2,
				next: &DoubleLinkedList{data: 3},
				prev: &DoubleLinkedList{data: 1},
			},
		},
	}

	for _, test := range tests {
		suite.Run("NewDoubleLinkedList", func() {
			result := NewDoubleLinkedList(test.data, test.next, test.prev)
			suite.Equal(test.expected.data, result.GetData())
			suite.Equal(test.expected.next, result.GetNext())
			suite.Equal(test.expected.prev, result.GetPrev())
		})
	}
}

func (suite *DoubleLinkedListTestSuite) TestInsertAtStartDoubleLL() {
	tests := []struct {
		insertValue  int
		expectedData []int
	}{
		{
			insertValue:  1,
			expectedData: []int{1},
		},
		{
			insertValue:  2,
			expectedData: []int{2, 1},
		},
		{
			insertValue:  3,
			expectedData: []int{3, 2, 1},
		},
	}

	for _, test := range tests {
		suite.Run("InsertAtStartDoubleLL", func() {
			InsertAtStartDoubleLL(&suite.head, test.insertValue)
			current := suite.head
			for _, expected := range test.expectedData {
				suite.Equal(expected, current.GetData())
				current = current.GetNext()
			}
			tail := suite.head
			for tail.GetNext() != nil {
				tail = tail.GetNext()
			}

			for i := len(test.expectedData) - 1; i >= 0; i-- {
				suite.Equal(test.expectedData[i], tail.GetData())
				tail = tail.GetPrev()
			}
		})
	}
}

func (suite *DoubleLinkedListTestSuite) TestInsertAtEndDoubleLL() {
	tests := []struct {
		insertValue  int
		expectedData []int
	}{
		{
			insertValue:  1,
			expectedData: []int{1},
		},
		{
			insertValue:  2,
			expectedData: []int{1, 2},
		},
		{
			insertValue:  3,
			expectedData: []int{1, 2, 3},
		},
	}

	for _, test := range tests {
		suite.Run("InsertAtEndDoubleLL", func() {
			InsertAtEndDoubleLL(&suite.head, test.insertValue)
			current := suite.head
			for _, expected := range test.expectedData {
				suite.Equal(expected, current.GetData())
				current = current.GetNext()
			}

			tail := suite.head
			for tail.GetNext() != nil {
				tail = tail.GetNext()
			}

			// Traverse from tail to head
			for i := len(test.expectedData) - 1; i >= 0; i-- {
				suite.Equal(test.expectedData[i], tail.GetData())
				tail = tail.GetPrev()
			}
		})
	}
}

func (suite *DoubleLinkedListTestSuite) TestInsertAtPositionOfDoubleLL() {
	tests := []struct {
		initialValues []int
		insertValue   int
		position      int
		expectedData  []int
	}{
		{
			initialValues: []int{1, 2, 4},
			insertValue:   3,
			position:      2,
			expectedData:  []int{1, 2, 3, 4},
		},
		{
			initialValues: []int{1, 3, 4},
			insertValue:   2,
			position:      1,
			expectedData:  []int{1, 2, 3, 4},
		},
		{
			initialValues: []int{2, 3, 4},
			insertValue:   1,
			position:      0,
			expectedData:  []int{1, 2, 3, 4},
		},
		{
			initialValues: []int{1, 2, 3},
			insertValue:   4,
			position:      3,
			expectedData:  []int{1, 2, 3, 4},
		},
		{
			initialValues: []int{},
			insertValue:   1,
			position:      0,
			expectedData:  []int{1},
		},
		{
			initialValues: []int{1, 2, 3},
			insertValue:   4,
			position:      8,
			expectedData:  []int{1, 2, 3},
		},
	}

	for index, test := range tests {
		testName := fmt.Sprintf("InsertAtPositionOfDoubleLL %v", index+1)
		suite.Run(testName, func() {
			for _, value := range test.initialValues {
				InsertAtEndDoubleLL(&suite.head, value)
			}
			InsertAtPositionOfDoubleLL(&suite.head, test.insertValue, test.position)
			current := suite.head
			for _, expected := range test.expectedData {
				suite.Equal(expected, current.GetData())
				current = current.GetNext()
			}
			tail := suite.head
			for tail.GetNext() != nil {
				tail = tail.GetNext()
			}
			for i := len(test.expectedData) - 1; i >= 0; i-- {
				suite.Equal(test.expectedData[i], tail.GetData())
				tail = tail.GetPrev()
			}
			suite.head = nil
		})
	}
}

func (suite *DoubleLinkedListTestSuite) TestGetDoubleLinkedListLength() {
	tests := []struct {
		initialValues  []int
		expectedLength int
	}{
		{
			initialValues:  []int{1, 2, 3},
			expectedLength: 3,
		},
		{
			initialValues:  []int{1, 2},
			expectedLength: 2,
		},
		{
			initialValues:  []int{1},
			expectedLength: 1,
		},
		{
			initialValues:  []int{},
			expectedLength: 0,
		},
	}

	for _, test := range tests {
		suite.Run("GetDoubleLinkedListLength", func() {
			for _, value := range test.initialValues {
				InsertAtEndDoubleLL(&suite.head, value)
			}
			length := GetDoubleLinkedListLength(suite.head)
			suite.Equal(test.expectedLength, length)
			suite.head = nil
		})
	}
}

func (suite *DoubleLinkedListTestSuite) TestDeleteAtStartOfDoubleLL() {
	tests := []struct {
		initialValues []int
		expectedData  []int
	}{
		{
			initialValues: []int{1},
			expectedData:  []int{},
		},
		{
			initialValues: []int{1, 2},
			expectedData:  []int{2},
		},
		{
			initialValues: []int{1, 2, 3},
			expectedData:  []int{2, 3},
		},
		{
			initialValues: []int{},
			expectedData:  []int{},
		},
	}

	for _, test := range tests {
		suite.Run(fmt.Sprintf("Delete at start from %v", test.initialValues), func() {
			for _, value := range test.initialValues {
				InsertAtEndDoubleLL(&suite.head, value)
			}

			DeleteAtStartOfDoubleLL(&suite.head)

			current := suite.head
			for _, expected := range test.expectedData {
				suite.Equal(expected, current.GetData())
				current = current.GetNext()
			}

			if suite.head != nil {
				tail := suite.head
				for tail.GetNext() != nil {
					tail = tail.GetNext()
				}

				for i := len(test.expectedData) - 1; i >= 0; i-- {
					suite.Equal(test.expectedData[i], tail.GetData())
					tail = tail.GetPrev()
				}
			}

			// Reset the list
			suite.head = nil
		})
	}
}

func (suite *DoubleLinkedListTestSuite) TestDeleteAtEndOfDoubleLL() {
	tests := []struct {
		initialValues []int
		expectedData  []int
	}{
		{
			initialValues: []int{1},
			expectedData:  []int{},
		},
		{
			initialValues: []int{1, 2},
			expectedData:  []int{1},
		},
		{
			initialValues: []int{1, 2, 3},
			expectedData:  []int{1, 2},
		},
		{
			initialValues: []int{},
			expectedData:  []int{},
		},
	}

	for _, test := range tests {
		suite.Run(fmt.Sprintf("Delete at end from %v", test.initialValues), func() {
			for _, value := range test.initialValues {
				InsertAtEndDoubleLL(&suite.head, value)
			}

			DeleteAtEndOfDoubleLL(&suite.head)

			current := suite.head
			for _, expected := range test.expectedData {
				suite.Equal(expected, current.GetData())
				current = current.GetNext()
			}

			if suite.head != nil {
				tail := suite.head
				for tail.GetNext() != nil {
					tail = tail.GetNext()
				}

				for i := len(test.expectedData) - 1; i >= 0; i-- {
					suite.Equal(test.expectedData[i], tail.GetData())
					tail = tail.GetPrev()
				}
			}

			suite.head = nil
		})
	}
}

func (suite *DoubleLinkedListTestSuite) TestDeleteAtPositionOfDoubleLL() {
	tests := []struct {
		initialValues []int
		position      int
		expectedData  []int
	}{
		{
			initialValues: []int{1},
			position:      0,
			expectedData:  []int{},
		},
		{
			initialValues: []int{1, 2},
			position:      1,
			expectedData:  []int{1},
		},
		{
			initialValues: []int{1, 2, 3},
			position:      1,
			expectedData:  []int{1, 3},
		},
		{
			initialValues: []int{1, 2, 3, 4},
			position:      2,
			expectedData:  []int{1, 2, 4},
		},
		{
			initialValues: []int{1, 2, 3, 4},
			position:      7,
			expectedData:  []int{1, 2, 3, 4},
		},
	}

	for _, test := range tests {
		suite.Run(fmt.Sprintf("Delete at position %d from %v", test.position, test.initialValues), func() {
			for _, value := range test.initialValues {
				InsertAtEndDoubleLL(&suite.head, value)
			}

			DeleteAtPositionOfDoubleLL(&suite.head, test.position)

			current := suite.head
			for _, expected := range test.expectedData {
				suite.Equal(expected, current.GetData())
				current = current.GetNext()
			}
			if suite.head != nil {
				tail := suite.head
				for tail.GetNext() != nil {
					tail = tail.GetNext()
				}

				for i := len(test.expectedData) - 1; i >= 0; i-- {
					suite.Equal(test.expectedData[i], tail.GetData())
					tail = tail.GetPrev()
				}
			}
			suite.head = nil
		})
	}
}

func (suite *DoubleLinkedListTestSuite) TestGettersAndSetters() {
	// Test cases for GetData, SetData, GetNext, SetNext, GetPrev, SetPrev
	tests := []struct {
		initialData int
		newData     int
		nextNode    *DoubleLinkedList
		prevNode    *DoubleLinkedList
	}{
		{
			initialData: 1,
			newData:     2,
			nextNode:    &DoubleLinkedList{data: 3},
			prevNode:    &DoubleLinkedList{data: 0},
		},
		{
			initialData: 4,
			newData:     5,
			nextNode:    &DoubleLinkedList{data: 6},
			prevNode:    &DoubleLinkedList{data: 3},
		},
	}

	for _, test := range tests {
		suite.Run("GettersAndSetters", func() {
			node := NewDoubleLinkedList(test.initialData, nil, nil)

			// Test GetData and SetData
			suite.Equal(test.initialData, node.GetData())
			node.SetData(test.newData)
			suite.Equal(test.newData, node.GetData())

			// Test GetNext and SetNext
			suite.Nil(node.GetNext())
			node.SetNext(test.nextNode)
			suite.Equal(test.nextNode, node.GetNext())

			// Test GetPrev and SetPrev
			suite.Nil(node.GetPrev())
			node.SetPrev(test.prevNode)
			suite.Equal(test.prevNode, node.GetPrev())
		})
	}
}
