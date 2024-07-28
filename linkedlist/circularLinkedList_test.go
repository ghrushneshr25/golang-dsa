package linkedlist

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type CircularLinkedListTestSuite struct {
	suite.Suite
	head *CircularLinkedList
}

func TestCircularLinkedList(t *testing.T) {
	suite.Run(t, &CircularLinkedListTestSuite{})
}

func (suite *CircularLinkedListTestSuite) SetupTest() {
	suite.head = nil
}

func (suite *CircularLinkedListTestSuite) TestInsertAtEndOfCircularLinkedList() {
	tests := []struct {
		values       []int
		expectedData []int
	}{
		{
			values:       []int{1},
			expectedData: []int{1},
		},
		{
			values:       []int{1, 2},
			expectedData: []int{1, 2},
		},
		{
			values:       []int{1, 2, 3},
			expectedData: []int{1, 2, 3},
		},
	}

	for index, test := range tests {
		testName := fmt.Sprintf("InsertAtEnd %v", index+1)
		suite.Run(testName, func() {
			for _, value := range test.values {
				InsertAtEndOfCircularLinkedList(&suite.head, value)
			}
			current := suite.head
			for _, expected := range test.expectedData {
				suite.Equal(expected, current.GetData())
				current = current.GetNext()
			}
			suite.Equal(suite.head, current)
			suite.head = nil
		})
	}
}

func (suite *CircularLinkedListTestSuite) TestInsertAtFrontOfCircularLinkedList() {
	tests := []struct {
		values       []int
		expectedData []int
	}{
		{
			values:       []int{1},
			expectedData: []int{1},
		},
		{
			values:       []int{2, 1},
			expectedData: []int{1, 2},
		},
		{
			values:       []int{3, 2, 1},
			expectedData: []int{1, 2, 3},
		},
	}

	for index, test := range tests {
		testName := fmt.Sprintf("InsertAtFront %v", index+1)
		suite.Run(testName, func() {
			for _, value := range test.values {
				InsertAtFrontOfCircularLinkedList(&suite.head, value)
			}
			current := suite.head
			for _, expected := range test.expectedData {
				suite.Equal(expected, current.GetData())
				current = current.GetNext()
			}
			suite.Equal(suite.head, current)
			suite.head = nil
		})
	}
}

func (suite *CircularLinkedListTestSuite) TestDeleteFirstNodeOfCircularLinkedList() {
	tests := []struct {
		initialValues []int
		expectedData  []int
	}{
		{
			initialValues: []int{1, 2, 3},
			expectedData:  []int{2, 3},
		},
		{
			initialValues: []int{2, 3},
			expectedData:  []int{3},
		},
		{
			initialValues: []int{3},
			expectedData:  []int{},
		},
		{
			initialValues: []int{},
			expectedData:  []int{},
		},
	}

	for index, test := range tests {
		testName := fmt.Sprintf("DeleteFirstNode %v", index+1)
		suite.Run(testName, func() {
			for _, value := range test.initialValues {
				InsertAtEndOfCircularLinkedList(&suite.head, value)
			}
			DeleteFirstNodeOfCircularLinkedList(&suite.head)
			current := suite.head
			for _, expected := range test.expectedData {
				suite.Equal(expected, current.GetData())
				current = current.GetNext()
			}
			if len(test.expectedData) == 0 {
				suite.Nil(current)
			} else {
				suite.Equal(suite.head, current)
			}
			suite.head = nil
		})
	}
}

func (suite *CircularLinkedListTestSuite) TestDeleteLastNodeOfCircularLinkedList() {
	tests := []struct {
		initialValues []int
		expectedData  []int
	}{
		{
			initialValues: []int{1, 2, 3},
			expectedData:  []int{1, 2},
		},
		{
			initialValues: []int{1, 2},
			expectedData:  []int{1},
		},
		{
			initialValues: []int{1},
			expectedData:  []int{},
		},
		{
			initialValues: []int{},
			expectedData:  []int{},
		},
	}

	for _, test := range tests {
		suite.Run("DeleteLastNode", func() {
			for _, value := range test.initialValues {
				InsertAtEndOfCircularLinkedList(&suite.head, value)
			}
			DeleteLastNodeOfCircularLinkedList(&suite.head)
			current := suite.head
			for _, expected := range test.expectedData {
				suite.Equal(expected, current.GetData())
				current = current.GetNext()
			}
			if len(test.expectedData) == 0 {
				suite.Nil(current)
			} else {
				suite.Equal(suite.head, current)
			}
			suite.head = nil
		})
	}
}

func (suite *CircularLinkedListTestSuite) TestGetCircularLinkedListLength() {
	tests := []struct {
		values         []int
		expectedLength int
	}{
		{
			values:         []int{},
			expectedLength: 0,
		},
		{
			values:         []int{1},
			expectedLength: 1,
		},
		{
			values:         []int{1, 2, 3},
			expectedLength: 3,
		},
	}

	for _, test := range tests {
		suite.Run("GetLength", func() {
			for _, value := range test.values {
				InsertAtEndOfCircularLinkedList(&suite.head, value)
			}
			length := GetCircularLinkedListLength(suite.head)
			suite.Equal(test.expectedLength, length)
			suite.head = nil
		})
	}
}

func (suite *CircularLinkedListTestSuite) TestPrintContentOfCircularLinkedList() {
	tests := []struct {
		values       []int
		expectedData string
	}{
		{
			values:       []int{1},
			expectedData: "1",
		},
		{
			values:       []int{1, 2},
			expectedData: "1 2",
		},
		{
			values:       []int{1, 2, 3},
			expectedData: "1 2 3",
		},
		{
			values:       []int{},
			expectedData: "",
		},
	}

	for index, test := range tests {
		testName := fmt.Sprintf("PrintContent %v", index+1)
		suite.Run(testName, func() {
			for _, value := range test.values {
				InsertAtEndOfCircularLinkedList(&suite.head, value)
			}
			result := PrintContentOfCircularLinkedList(suite.head)
			suite.Equal(test.expectedData, result)
			suite.head = nil
		})
	}
}

func (suite *CircularLinkedListTestSuite) TestGettersAndSetters() {
	// Test SetData and GetData
	node := &CircularLinkedList{}
	node.SetData(10)
	suite.Equal(10, node.GetData(), "GetData should return the value set by SetData")

	// Test SetNext and GetNext
	nextNode := &CircularLinkedList{}
	node.SetNext(nextNode)
	suite.Equal(nextNode, node.GetNext(), "GetNext should return the node set by SetNext")
}
