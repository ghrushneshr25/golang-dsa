package linkedlist

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type SingleLinkedListTestSuite struct {
	suite.Suite
	head *SingleLinkedList
}

func TestSingleLinkedList(t *testing.T) {
	suite.Run(t, new(SingleLinkedListTestSuite))
}

func (suite *SingleLinkedListTestSuite) SetupTest() {
	suite.head = nil
}

func (suite *SingleLinkedListTestSuite) TestInsertAtStartOfSingleLL() {
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
		testName := fmt.Sprintf("InsertAtStart %v", index+1)
		suite.Run(testName, func() {
			for _, value := range test.values {
				InsertAtStartofSingleLL(&suite.head, value)
			}
			current := suite.head
			for _, expected := range test.expectedData {
				suite.Equal(expected, current.GetData())
				current = current.GetNext()
			}
			suite.head = nil
		})
	}
}

func (suite *SingleLinkedListTestSuite) TestInsertAtEndOfSingleLL() {
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
				InsertAtEndofSingleLL(&suite.head, value)
			}
			current := suite.head
			for _, expected := range test.expectedData {
				suite.Equal(expected, current.GetData())
				current = current.GetNext()
			}
			suite.head = nil
		})
	}
}

func (suite *SingleLinkedListTestSuite) TestDeleteAtStartOfSingleLL() {
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
		testName := fmt.Sprintf("DeleteAtStart %v", index+1)
		suite.Run(testName, func() {
			for _, value := range test.initialValues {
				InsertAtEndofSingleLL(&suite.head, value)
			}
			DeleteAtStartOfSingleLL(&suite.head)
			current := suite.head
			for _, expected := range test.expectedData {
				suite.Equal(expected, current.GetData())
				current = current.GetNext()
			}
			suite.head = nil
		})
	}
}

func (suite *SingleLinkedListTestSuite) TestDeleteAtEndOfSingleLL() {
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

	for index, test := range tests {
		testName := fmt.Sprintf("DeleteAtEnd %v", index+1)
		suite.Run(testName, func() {
			for _, value := range test.initialValues {
				InsertAtEndofSingleLL(&suite.head, value)
			}
			DeleteAtEndSingleLL(&suite.head)
			current := suite.head
			for _, expected := range test.expectedData {
				suite.Equal(expected, current.GetData())
				current = current.GetNext()
			}
			suite.head = nil
		})
	}
}

func (suite *SingleLinkedListTestSuite) TestGetSingleLinkedListLength() {
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

	for index, test := range tests {
		testName := fmt.Sprintf("GetLength %v", index+1)
		suite.Run(testName, func() {
			for _, value := range test.values {
				InsertAtEndofSingleLL(&suite.head, value)
			}
			length := GetSingleLinkedListLength(suite.head)
			suite.Equal(test.expectedLength, length)
			suite.head = nil
		})
	}
}

func (suite *SingleLinkedListTestSuite) TestDeleteAtPositionOfSingleLL() {
	tests := []struct {
		initialValues []int
		position      int
		expectedData  []int
	}{
		{
			initialValues: []int{1, 2, 3, 4, 5},
			position:      0,
			expectedData:  []int{2, 3, 4, 5},
		},
		{
			initialValues: []int{1, 2, 3, 4, 5},
			position:      2,
			expectedData:  []int{1, 2, 4, 5},
		},
		{
			initialValues: []int{1, 2, 3, 4, 5},
			position:      4,
			expectedData:  []int{1, 2, 3, 4},
		},
		{
			initialValues: []int{1, 2, 3, 4, 5},
			position:      5,
			expectedData:  []int{1, 2, 3, 4, 5}, // position out of bounds
		},
		{
			initialValues: []int{1},
			position:      0,
			expectedData:  []int{},
		},
		{
			initialValues: []int{},
			position:      0,
			expectedData:  []int{}, // empty list
		},
	}

	for index, test := range tests {
		testName := fmt.Sprintf("DeleteAtPosition %v", index+1)
		suite.Run(testName, func() {
			for _, value := range test.initialValues {
				InsertAtEndofSingleLL(&suite.head, value)
			}
			DeleteAtPositionOfSingleLL(&suite.head, test.position)
			current := suite.head
			for _, expected := range test.expectedData {
				suite.Equal(expected, current.GetData())
				current = current.GetNext()
			}
			suite.head = nil
		})
	}
}

func (suite *SingleLinkedListTestSuite) TestInsertAtPositionOfSingleLL() {
	tests := []struct {
		initialValues []int
		data          int
		position      int
		expectedData  []int
	}{
		{
			initialValues: []int{1, 2, 3, 4, 5},
			data:          10,
			position:      0,
			expectedData:  []int{10, 1, 2, 3, 4, 5},
		},
		{
			initialValues: []int{1, 2, 3, 4, 5},
			data:          10,
			position:      2,
			expectedData:  []int{1, 2, 10, 3, 4, 5},
		},
		{
			initialValues: []int{1, 2, 3, 4, 5},
			data:          10,
			position:      5,
			expectedData:  []int{1, 2, 3, 4, 5, 10},
		},
		{
			initialValues: []int{1, 2, 3, 4, 5},
			data:          10,
			position:      8,
			expectedData:  []int{1, 2, 3, 4, 5}, // position out of bounds
		},
		{
			initialValues: []int{},
			data:          10,
			position:      0,
			expectedData:  []int{10},
		},
	}

	for index, test := range tests {
		testName := fmt.Sprintf("InsertAtPosition %v", index+1)
		suite.Run(testName, func() {
			for _, value := range test.initialValues {
				InsertAtEndofSingleLL(&suite.head, value)
			}
			InsertAtPositionOfSingleLL(&suite.head, test.data, test.position)
			current := suite.head
			for _, expected := range test.expectedData {
				suite.Equal(expected, current.GetData())
				current = current.GetNext()
			}
			suite.head = nil
		})
	}
}

func (suite *SingleLinkedListTestSuite) TestSingleLinkedListGettersSetters() {
	tests := []struct {
		data     int
		nextData int
	}{
		{
			data:     1,
			nextData: 2,
		},
		{
			data:     10,
			nextData: 20,
		},
		{
			data:     100,
			nextData: 200,
		},
	}

	for index, test := range tests {
		testName := fmt.Sprintf("GettersSetters %v", index+1)
		suite.Run(testName, func() {
			node := &SingleLinkedList{}
			node.SetData(test.data)
			suite.Equal(test.data, node.GetData())

			nextNode := &SingleLinkedList{}
			nextNode.SetData(test.nextData)
			node.SetNext(nextNode)
			suite.Equal(nextNode, node.GetNext())
			suite.Equal(test.nextData, node.GetNext().GetData())
		})
	}
}

func (suite *SingleLinkedListTestSuite) TestSearchInSingleLL() {
	intPtr := func(i int) *int { return &i }
	tests := []struct {
		initialValues []int
		searchData    int
		expectedData  *int
	}{
		{
			initialValues: []int{1, 2, 3, 4, 5},
			searchData:    3,
			expectedData:  intPtr(3),
		},
		{
			initialValues: []int{1, 2, 3, 4, 5},
			searchData:    6,
			expectedData:  nil,
		},
		{
			initialValues: []int{},
			searchData:    1,
			expectedData:  nil,
		},
	}

	for index, test := range tests {
		testName := fmt.Sprintf("SearchInSingleLL %v", index+1)
		suite.Run(testName, func() {
			for _, value := range test.initialValues {
				InsertAtEndofSingleLL(&suite.head, value)
			}
			result := SearchInSingleLL(suite.head, test.searchData)
			if test.expectedData == nil {
				suite.Nil(result)
			} else {
				suite.NotNil(result)
				suite.Equal(*test.expectedData, result.GetData())
			}
			suite.head = nil
		})
	}
}
