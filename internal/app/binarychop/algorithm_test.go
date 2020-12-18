package binarychop

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

//
type TestCase struct {
	input      []int
	lookingFor int
	expected   int
}

// AlgorithmMock is a dumb mock to test Algorithm interface with dep injection
type AlgorithmMock struct{}

// BinaryChop : interface implementation, return the value passed in argument
func (algoMock *AlgorithmMock) BinaryChop(value int, _ []int) int {
	return value
}

func createIntSlice(n int) []int {
	ints := make([]int, n)
	for index := 0; index < n; index++ {
		ints[index] = index
	}
	return ints
}

// GeneralTestInterfaceImplementation is a test template for Algorithm implementation.
// Every implementation of the Algorithm interface should trigger this test.
// Its purpose is to test the implementation without wrapper
func GeneralTestInterfaceImplementation(algo Algorithm, t *testing.T) {
	testCases := []TestCase{
		{
			input:      []int{-4, -2, -1, 0, 2, 5, 7, 10, 11, 12, 15, 25},
			lookingFor: 5,
			expected:   5,
		},
		{
			input:      []int{-4, -2, -1, 0, 2, 5, 7, 10, 11, 12, 15, 25},
			lookingFor: -4,
			expected:   0,
		},
		{
			input:      []int{-4, -2, -1, 0, 2, 5, 7, 10, 11, 12, 15, 25},
			lookingFor: 12,
			expected:   9,
		},
		{
			input:      []int{-4, -2, -1, 0, 2, 5, 7, 10, 11, 12, 15, 25},
			lookingFor: 100,
			expected:   -1,
		},
	}
	for _, test := range testCases {
		assert.Equal(t, test.expected, algo.BinaryChop(test.lookingFor, test.input))
	}
}

// ShuffleGeneralTestInterface is a test template for Algorithm implementation
// Every implementation of the Algorithm interface should trigger this test.
// Its purpose is to test random value in a shuffled int slice using a wrapper
func ShuffleGeneralTestInterface(algo Algorithm, t *testing.T) {
	testCases := []TestCase{
		{
			input:      []int{0, -2, 7, 11, -4, 10, 15, 12, -1, 2, 25, 5},
			lookingFor: 5,
			expected:   5,
		},
		{
			input:      []int{-4, -2, -1, 0, 2, 5, 7, 10, 11, 12, 15, 25},
			lookingFor: -4,
			expected:   0,
		},
		{
			input:      []int{-4, -2, -1, 0, 2, 5, 7, 10, 11, 12, 15, 25},
			lookingFor: 12,
			expected:   9,
		},
		{
			input:      []int{-4, -2, -1, 0, 2, 5, 7, 10, 11, 12, 15, 25},
			lookingFor: 100,
			expected:   -1,
		},
	}
	chop := NewBinaryChop(algo)
	for _, test := range testCases {
		assert.Equal(t, test.expected, chop.Execute(test.lookingFor, test.input))
	}
}

// GeneralTestInterface is a test template for Algorithm implementation.
// Every implementation of the Algorithm interface should trigger this test.
// It tests implementation as asked in the code kata : http://codekata.com/kata/kata02-karate-chop/
func GeneralTestInterface(algo Algorithm, t *testing.T) {
	chop := NewBinaryChop(algo)
	// ##
	assert.Equal(t, -1, chop.Execute(3, []int{}))
	assert.Equal(t, -1, chop.Execute(3, []int{1}))
	assert.Equal(t, 0, chop.Execute(1, []int{1}))
	// ##
	assert.Equal(t, 0, chop.Execute(1, []int{1, 3, 5}))
	assert.Equal(t, 1, chop.Execute(3, []int{1, 3, 5}))
	assert.Equal(t, 2, chop.Execute(5, []int{1, 3, 5}))
	assert.Equal(t, -1, chop.Execute(0, []int{1, 3, 5}))
	assert.Equal(t, -1, chop.Execute(2, []int{1, 3, 5}))
	assert.Equal(t, -1, chop.Execute(4, []int{1, 3, 5}))
	assert.Equal(t, -1, chop.Execute(6, []int{1, 3, 5}))
	// ##
	assert.Equal(t, 0, chop.Execute(1, []int{1, 3, 5, 7}))
	assert.Equal(t, 1, chop.Execute(3, []int{1, 3, 5, 7}))
	assert.Equal(t, 2, chop.Execute(5, []int{1, 3, 5, 7}))
	assert.Equal(t, 3, chop.Execute(7, []int{1, 3, 5, 7}))
	assert.Equal(t, -1, chop.Execute(0, []int{1, 3, 5, 7}))
	assert.Equal(t, -1, chop.Execute(2, []int{1, 3, 5, 7}))
	assert.Equal(t, -1, chop.Execute(4, []int{1, 3, 5, 7}))
	assert.Equal(t, -1, chop.Execute(6, []int{1, 3, 5, 7}))
	assert.Equal(t, -1, chop.Execute(8, []int{1, 3, 5, 7}))
}

//TestNewBinaryChop test if dependencies injection work
func TestNewBinaryChop(t *testing.T) {
	fakeAlgo := new(AlgorithmMock)
	binaryChop := NewBinaryChop(fakeAlgo)
	assert.Equal(t, binaryChop.algorithm, fakeAlgo, "algorithms should be pointer")
}

//TestBinaryChop_Execute test if dep injection implementation can be trigger, test error value
func TestBinaryChop_Execute(t *testing.T) {
	fakeAlgo := new(AlgorithmMock)
	binaryChop := NewBinaryChop(fakeAlgo)

	// test caching empty slice
	assert.Equal(t, -1, binaryChop.Execute(5, []int{}))

	// test "bad value"
	assert.Equal(t, -1, binaryChop.Execute(-1, []int{1}))

	// test "good value"
	assert.Equal(t, 0, binaryChop.Execute(0, []int{0}))
}

// ExampleNewBinaryChop canonical example of how to use BinaryChop interface
func ExampleNewBinaryChop() {
	values := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	valueToFind := 5
	algo := &AlgorithmMock{}
	chop := NewBinaryChop(algo)

	indexFound := chop.Execute(valueToFind, values)

	if indexFound == -1 {
		fmt.Println("value not found")
		return
	}
	fmt.Println("value ", valueToFind, "found at index", indexFound)
	// Output :
	// value 5 found at index 5
}
