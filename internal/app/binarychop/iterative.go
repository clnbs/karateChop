package binarychop

// Iterative exists to implement algorithm interface
type Iterative struct {
}

// BinaryChop implement binarychop.Algorithm interface for the iterative algorithm
func (iterative *Iterative) BinaryChop(value int, tab []int) int {
	if len(tab) == 1 {
		if value == tab[0] {
			return 0
		}
		return -1
	}
	// start algorithm in the middle of the slice
	tmpIndex := len(tab) / 2
	highIndex := len(tab) - 1
	lowIndex := 0
	for tab[tmpIndex] != value && lowIndex != highIndex {
		// value is in the upper part of the slice; setting new lower boundaries
		if tab[tmpIndex] < value {
			lowIndex = tmpIndex + 1
		}
		// value is in the lower part of the slice; setting new upper boundaries
		if tab[tmpIndex] > value {
			highIndex = tmpIndex
		}
		// re compute the index within the boundaries
		tmpIndex = ((highIndex - lowIndex) / 2) + lowIndex
		// Ending condition : success, value found
		if tab[tmpIndex] == value {
			return tmpIndex
		}
	}
	//Ending condition : failure, value not found
	if lowIndex == highIndex {
		return -1
	}
	return tmpIndex
}
