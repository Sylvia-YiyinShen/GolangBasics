package main
import "fmt"

func basicsOfFunctionParameters() {
	fmt.Printf("basics of function parameters\n")

	// passFunctionAsParameter()
	closureBoundToVariables()
}

func closureBoundToVariables() {
	closureA := adder()
	for i := 0; i < 10; i++ {
		fmt.Println(closureA(i))
	}

	fmt.Println("------closure------")
	closureB := adder()
	for i := 10; i > 0; i-- {
		fmt.Println(closureB(2*i))
	}
}

// adder is bound to sum
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func passFunctionAsParameter() {
	a := float64(10)
	b := float64(12)
	sumFunc := func(a, b float64) float64 {
		return a + b
	}
	multipleFunc := func(a, b float64) float64 {
		return a * b
	}

	fmt.Printf("compute with sumFunc: %v\n", compute(sumFunc, a, b))
	fmt.Printf("compute with multipleFunc: %v\n", compute(multipleFunc, a, b))
}


func compute(computeFunc func(float64, float64) float64, valueA float64, valueB float64) float64 {
	return computeFunc(valueA, valueB)
}
