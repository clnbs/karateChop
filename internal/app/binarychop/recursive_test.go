package binarychop

import (
	"fmt"
	"testing"

	"github.com/clnbs/karateChop/internal/pkg/mathtool"
)

func createIntSlice(n int) []int {
	ints := make([]int, n)
	for index := 0; index < n; index++ {
		ints[index] = index
	}
	return ints
}

// TestRecursive_BinaryChop trigger test for recursive algorithm without Algorithm interface
func TestRecursive_BinaryChop(t *testing.T) {
	GeneralTestInterfaceImplementation(&Recursive{}, t)
}

// TestRecursive_BinaryChop2 trigger test for recursive algorithm within Algorithm interface
func TestRecursive_BinaryChop2(t *testing.T) {
	ShuffleGeneralTestInterface(&Recursive{}, t)
}

// TestRecursive_BinaryChop3 trigger tests like asked in the code kata
func TestRecursive_BinaryChop3(t *testing.T) {
	GeneralTestInterface(&Recursive{}, t)
}

func benchmarkRecursiveBinaryChop(tab []int, b *testing.B) {
	recursive := Recursive{}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		value := mathtool.RandomIntBetween(0, len(tab)-1)
		recursive.BinaryChop(value, tab)
	}
}

// BenchmarkRecursive_BinaryChop10 trigger benchmark with a slice of length 10 for recursive algorithm
func BenchmarkRecursive_BinaryChop10(b *testing.B) {
	sliceLength10 := createIntSlice(10)
	benchmarkRecursiveBinaryChop(sliceLength10, b)
}

// BenchmarkRecursive_BinaryChop100 trigger benchmark with a slice of length 100 for recursive algorithm
func BenchmarkRecursive_BinaryChop100(b *testing.B) {
	sliceLength100 := createIntSlice(100)
	benchmarkRecursiveBinaryChop(sliceLength100, b)
}

// BenchmarkRecursive_BinaryChop1000 trigger benchmark with a slice of length 1000 for recursive algorithm
func BenchmarkRecursive_BinaryChop1000(b *testing.B) {
	sliceLength1000 := createIntSlice(1000)
	benchmarkRecursiveBinaryChop(sliceLength1000, b)
}

// BenchmarkRecursive_BinaryChop10000 trigger benchmark with a slice of length 10000 for recursive algorithm
func BenchmarkRecursive_BinaryChop10000(b *testing.B) {
	sliceLength10000 := createIntSlice(10000)
	benchmarkRecursiveBinaryChop(sliceLength10000, b)
}

//ExampleRecursive_BinaryChop canonical example of how to use Recursive Binary chop
func ExampleRecursive_BinaryChop() {
	values := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	valueToFind := 5
	chop := NewBinaryChop(&Recursive{})

	indexFound := chop.Execute(valueToFind, values)
	if indexFound == -1 {
		fmt.Println("value", valueToFind, "not found")
	}
	fmt.Println("value ", valueToFind, "found at index", indexFound)
	// Output :
	// value 5 found at index 5
}
