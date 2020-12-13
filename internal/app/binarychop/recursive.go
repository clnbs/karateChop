package binarychop

// Recursive exists to implement algorithm interface
type Recursive struct {
}

// BinaryChop implement binarychop.Algorithm interface
func (recursive *Recursive) BinaryChop(value int, tab []int) int {
	return recursiveBinaryChop(value, tab, 0, len(tab)-1)
}

// recursiveBinaryChop looks for a value in a int slice using binary chop.
// This binary chop algorithm use recursion
func recursiveBinaryChop(value int, tab []int, lowIndex, highIndex int) int {
	middleIndex := ((highIndex - lowIndex) / 2) + lowIndex
	// Ending condition : success
	if tab[middleIndex] == value {
		return middleIndex
	}
	// Ending condition : failure - no value found
	if lowIndex == highIndex {
		return -1
	}
	// Use recursion to explore the upper slice's indexes
	if tab[middleIndex] < value {
		return recursiveBinaryChop(value, tab, middleIndex+1, highIndex)
	}
	// Use recursion to explore the lower slice's indexes
	if tab[middleIndex] > value {
		return recursiveBinaryChop(value, tab, lowIndex, middleIndex)
	}
	// Ending condition : failure
	// unreachable code but golang rise an error if it is not here
	return -1
}
