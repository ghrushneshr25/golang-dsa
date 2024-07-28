package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type UnrolledLinkedListTestSuite struct {
	suite.Suite
	list *UnrolledLinkedList
}

func TestUnrolledLinkedList(t *testing.T) {
	suite.Run(t, new(UnrolledLinkedListTestSuite))
}

func (suite *UnrolledLinkedListTestSuite) SetupTest() {
	var err error
	suite.list, err = NewUnrolledLinkedList(4)
	suite.NoError(err)
}

func (suite *UnrolledLinkedListTestSuite) TestNewUnrolledLinkedList() {
	tests := []struct {
		nodeCapacity int
		expectError  bool
	}{
		{nodeCapacity: 4, expectError: false},
		{nodeCapacity: 3, expectError: true},
	}

	for _, test := range tests {
		suite.Run("NewUnrolledLinkedList", func() {
			list, err := NewUnrolledLinkedList(test.nodeCapacity)
			if test.expectError {
				suite.Error(err)
			} else {
				suite.NoError(err)
				suite.NotNil(list)
				suite.Equal(test.nodeCapacity, list.NodeCapacity)
			}
		})
	}
}

func (suite *UnrolledLinkedListTestSuite) TestAdd() {
	tests := []struct {
		values       []int
		expectedSize int
	}{
		{values: []int{1}, expectedSize: 1},
		{values: []int{2, 3, 4, 5}, expectedSize: 5},
	}

	for _, test := range tests {
		suite.Run("Add", func() {
			for _, val := range test.values {
				suite.list.Add(val)
			}
			suite.Equal(test.expectedSize, suite.list.Size())
		})
	}
}

func (suite *UnrolledLinkedListTestSuite) TestRemove() {
	suite.list.Add(1)
	suite.list.Add(2)
	suite.list.Add(3)
	suite.list.Add(5)
	suite.list.Add(6)
	suite.list.Add(7)
	suite.list.Add(8)

	tests := []struct {
		value          int
		expectedSize   int
		expectedResult bool
	}{
		{value: 2, expectedSize: 6, expectedResult: true},
		{value: 4, expectedSize: 6, expectedResult: false},
		{value: 7, expectedSize: 5, expectedResult: true},
	}

	for _, test := range tests {
		suite.Run("Remove", func() {
			result := suite.list.Remove(test.value)
			suite.Equal(test.expectedResult, result)
			suite.Equal(test.expectedSize, suite.list.Size())
		})
	}
}

func (suite *UnrolledLinkedListTestSuite) TestGet() {
	suite.list.Add(1)
	suite.list.Add(2)
	suite.list.Add(3)

	tests := []struct {
		index         int
		expectedValue int
		expectError   bool
	}{
		{index: 0, expectedValue: 1, expectError: false},
		{index: 2, expectedValue: 3, expectError: false},
		{index: 3, expectedValue: 0, expectError: true},
	}

	for _, test := range tests {
		suite.Run("Get", func() {
			value, err := suite.list.Get(test.index)
			if test.expectError {
				suite.Error(err)
			} else {
				suite.NoError(err)
				suite.Equal(test.expectedValue, value)
			}
		})
	}
}

func (suite *UnrolledLinkedListTestSuite) TestSet() {
	suite.list.Add(1)
	suite.list.Add(2)
	suite.list.Add(3)

	tests := []struct {
		index            int
		newValue         int
		expectedOldValue int
		expectError      bool
	}{
		{index: 0, newValue: 10, expectedOldValue: 1, expectError: false},
		{index: 2, newValue: 30, expectedOldValue: 3, expectError: false},
		{index: 3, newValue: 40, expectedOldValue: 0, expectError: true},
	}

	for _, test := range tests {
		suite.Run("Set", func() {
			oldValue, err := suite.list.Set(test.index, test.newValue)
			if test.expectError {
				suite.Error(err)
			} else {
				suite.NoError(err)
				suite.Equal(test.expectedOldValue, oldValue)
			}
		})
	}
}

func (suite *UnrolledLinkedListTestSuite) TestClear() {
	suite.list.Add(1)
	suite.list.Add(2)
	suite.list.Add(3)
	suite.list.Add(4)
	suite.list.Add(5)

	suite.list.Clear()
	suite.Equal(0, suite.list.Size())
	suite.True(suite.list.IsEmpty())
}

func (suite *UnrolledLinkedListTestSuite) TestContains() {
	suite.list.Add(1)
	suite.list.Add(2)
	suite.list.Add(3)

	tests := []struct {
		value          int
		expectedResult bool
	}{
		{value: 2, expectedResult: true},
		{value: 4, expectedResult: false},
	}

	for _, test := range tests {
		suite.Run("Contains", func() {
			result := suite.list.Contains(test.value)
			suite.Equal(test.expectedResult, result)
		})
	}
}

func (suite *UnrolledLinkedListTestSuite) TestAddAt() {
	suite.list.Add(1)
	suite.list.Add(3)

	tests := []struct {
		index        int
		value        int
		expectedSize int
		expectError  bool
	}{
		{index: 1, value: 2, expectedSize: 3, expectError: false},
		{index: 3, value: 4, expectedSize: 4, expectError: false},
		{index: 4, value: 5, expectedSize: 5, expectError: false},
		{index: 5, value: 6, expectedSize: 6, expectError: false},
		{index: 6, value: 7, expectedSize: 7, expectError: false},
		{index: 7, value: 8, expectedSize: 8, expectError: false},
		{index: 8, value: 9, expectedSize: 9, expectError: false},
		{index: 3, value: 20, expectedSize: 10, expectError: false},
		{index: 20, value: 5, expectedSize: 4, expectError: true},
	}

	for _, test := range tests {
		suite.Run("AddAt", func() {
			err := suite.list.AddAt(test.index, test.value)
			if test.expectError {
				suite.Error(err)
			} else {
				suite.NoError(err)
				suite.Equal(test.expectedSize, suite.list.Size())
			}
		})
	}
}

func (suite *UnrolledLinkedListTestSuite) TestRemoveAt() {
	suite.list.Add(1)
	suite.list.Add(2)
	suite.list.Add(3)

	tests := []struct {
		index         int
		expectedValue int
		expectedSize  int
		expectError   bool
	}{
		{index: 1, expectedValue: 2, expectedSize: 2, expectError: false},
		{index: 2, expectedValue: 0, expectedSize: 2, expectError: true},
	}

	for _, test := range tests {
		suite.Run("RemoveAt", func() {
			value, err := suite.list.RemoveAt(test.index)
			if test.expectError {
				suite.Error(err)
			} else {
				suite.NoError(err)
				suite.Equal(test.expectedValue, value)
				suite.Equal(test.expectedSize, suite.list.Size())
			}
		})
	}
}
