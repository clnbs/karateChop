package binarychop

import (
	"sort"
)

// Algorithm can be satisfied to create a new binary chop algorithm with dependencies injection
type Algorithm interface {
	BinaryChop(int, []int) int
}

// BinaryChop is a wrapper of Algorithm interface
type BinaryChop struct {
	algorithm Algorithm
}

// NewBinaryChop is the BinaryChop constructor
func NewBinaryChop(algo Algorithm) *BinaryChop {
	binaryChop := new(BinaryChop)
	binaryChop.algorithm = algo
	return binaryChop
}

// Execute is a wrapper for algorithm injected as a dependency.
// Only this method should trigger ErrorNoDataInSlice and ErrorNoMatchingValue errors
func (binaryChop *BinaryChop) Execute(value int, tab []int) int {
	if len(tab) == 0 {
		return -1
	}
	sort.Ints(tab)
	return binaryChop.algorithm.BinaryChop(value, tab)
}
