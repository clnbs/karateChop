package binarychop

import "testing"

func TestIterative_BinaryChop(t *testing.T) {
	GeneralTestInterfaceImplementation(&Iterative{}, t)
}

func TestIterative_BinaryChop2(t *testing.T) {
	ShuffleGeneralTestInterface(&Iterative{}, t)
}

func TestIterative_BinaryChop3(t *testing.T) {
	GeneralTestInterface(&Iterative{}, t)
}
