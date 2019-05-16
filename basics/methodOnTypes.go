package main
import (
	"fmt"
	"math"
)

// defined in map.go
// type Point struct {
// 	x, y float64
// }

func (p Point) distanceToZero() float64 {
	return math.Sqrt(p.x * p.x + p.y * p.y)
}

func (p *Point) Scale(s float64) {
	p.x = p.x * s
	p.y = p.y * s
}


type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func methodOnTypes() {
	fmt.Printf("methodOnTypes\n")
	typeMethod()
	otherTypeMethod()
	pointerMethod()

	//function should take the exact right type, method can take value or pointer: https://tour.golang.org/methods/6
}

func pointerMethod() {
	var point = Point{1,1}
	fmt.Printf("before scale: %v\n", point)
	point.Scale(1.5)
	fmt.Printf("after scale: %v\n", point)
}

func typeMethod(){
	fmt.Printf("{0, 0} distanceTo {0, 0} : %v\n", Point{0,0}.distanceToZero())
	fmt.Printf("{1, 2} distdanceTo {0, 0} : %v\n", Point{1,2}.distanceToZero())
	fmt.Printf("{2, 4} distanceTo {0, 0} : %v\n", Point{2,4}.distanceToZero())
}

func otherTypeMethod() {
	var myFloatA MyFloat = -6 
	fmt.Printf("myFloatA.Abs() %v\n", myFloatA.Abs(),)
}