package main
import "fmt"

func bascisOfMap() {
	fmt.Printf("basics of map\n")

	// var pointsMap = make(map[string]Point)
	// pointsMap["start"] = Point{0,0}
	// pointsMap["end"] = Point{3,5}

	var pointsMap = map[string]Point{
		"start": Point{0,0},
		"end": Point{3,5},
	}

	delete(pointsMap, "start")
	elem, found := pointsMap["start"]
	fmt.Printf("found 'start': %v  ", found)
	if found { 
		fmt.Printf("pointsMap['start']: %v\n", elem)
	} else {
		fmt.Printf("default value: %v\n", elem)
	}

	fmt.Printf("pointsMap: %v\n", pointsMap)
	fmt.Printf("start point: %v\n", pointsMap["start"])
	fmt.Printf("? point: %v\n", pointsMap["?"])
}

type Point struct{
	x, y float64
}
