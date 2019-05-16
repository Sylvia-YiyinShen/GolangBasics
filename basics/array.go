package main
import "fmt"

// https://blog.golang.org/go-slices-usage-and-internals

func basicsOfArray() {
	fmt.Printf("basics of array\n")
	// arrayAndSlice()
	// appendingElement()
	iterateOverSlice()
}

func iterateOverSlice() {
	var list = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range list {
		fmt.Printf("i: %d, v: %d\n", i, v)
		// fmt.Printf("2**%d = %d\n", i, v)
	}
}

func appendingElement() {
	var list = []int{0}
	fmt.Printf("before appending %v\n",list)

	list = append(list, 1)
	list = append(list, 2, 3)
	fmt.Printf("after appending %v\n",list)
}

func arrayAndSlice() {
	a := make([]int, 5)
	printSlice("a", a)

	numList := []int{0, 1, 2, 3, 4}
	printSlice("numList", numList)

	//range [a:b] a...b-1

	slice :=  numList[:3] // len = 3 (0,1,2) cap = 5
	printSlice("slice numList[:3]", slice)

	slice2 :=  numList[2:] // len = 3 (index starts at 2) cap = 3
	printSlice("slice2 numList[2:]", slice2)
}

func printSlice(sliceName string, slice []int) {
	fmt.Printf("%s:\nlen = %d, cap = %d %v\n\n", sliceName, len(slice), cap(slice), slice)
}