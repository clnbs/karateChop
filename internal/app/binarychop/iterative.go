package binarychop

type Iterative struct {
}

func (iterative *Iterative) BinaryChop(value int, tab []int) int {
	if len(tab) == 1 {
		if value == tab[0] {
			return 0
		}
		return -1
	}
	tmpIndex := len(tab) / 2
	highIndex := len(tab) - 1
	lowIndex := 0
	for tab[tmpIndex] != value && lowIndex != highIndex {
		if tab[tmpIndex] < value {
			lowIndex = tmpIndex + 1
		}
		if tab[tmpIndex] > value {
			highIndex = tmpIndex
		}
		tmpIndex = ((highIndex - lowIndex) / 2) + lowIndex
		if tab[tmpIndex] == value {
			return tmpIndex
		}
	}
	if lowIndex == highIndex {
		return -1
	}
	return tmpIndex
}
