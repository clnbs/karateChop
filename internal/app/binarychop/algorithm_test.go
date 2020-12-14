package binarychop

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	// Values is a int slice used for testing purpose
	Values = []int{-4, -2, -1, 0, 2, 5, 7, 10, 11, 12, 15, 25}
	// ShuffledValues is a copy of Values but shuffled for testing purpose
	ShuffledValues = make([]int, len(Values))
	// ExpectedResult map a value contains in Values to its index in order to find result easily
	ExpectedResult = map[int]int{}
)

// AlgorithmMock is a dumb mock to test Algorithm interface with dep injection
type AlgorithmMock struct{}

// BinaryChop : interface implementation, return the value passed in argument
func (algoMock *AlgorithmMock) BinaryChop(value int, _ []int) int {
	return value
}

func init() {
	sort.Ints(Values)
	copy(ShuffledValues, Values)
	for index, value := range Values {
		ExpectedResult[value] = index
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(ShuffledValues), func(i, j int) { ShuffledValues[i], ShuffledValues[j] = ShuffledValues[j], ShuffledValues[i] })
}

func findExpectedResult(value int) int {
	if _, ok := ExpectedResult[value]; !ok {
		return -1
	}
	return ExpectedResult[value]
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
	lowestValue, highestValue := Values[0], Values[len(Values)-1]
	for testingValue := lowestValue; testingValue <= highestValue; testingValue++ {
		expectedResult := findExpectedResult(testingValue)
		assert.Equal(t, expectedResult, algo.BinaryChop(testingValue, Values))
	}
}

// ShuffleGeneralTestInterface is a test template for Algorithm implementation
// Every implementation of the Algorithm interface should trigger this test.
// Its purpose is to test random value in a shuffled int slice using a wrapper
func ShuffleGeneralTestInterface(algo Algorithm, t *testing.T) {
	chop := NewBinaryChop(algo)
	testingValue := 12
	expectedValue := findExpectedResult(testingValue)
	assert.Equal(t, expectedValue, chop.Execute(testingValue, ShuffledValues))

	testingValue = -3
	expectedValue = findExpectedResult(testingValue)
	assert.Equal(t, expectedValue, chop.Execute(testingValue, ShuffledValues))
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
