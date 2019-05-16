package main
import "fmt"

type Animal interface {
	eat()
}

type Person struct {
	Name string
	Age  int
}

func (p Person) eat() {
	fmt.Println("having meal...")
}

type Cat struct {}
// func (cat Cat) eat() {
// 	fmt.Println("having fish...")
// }

func (cat *Cat) eat() {
	fmt.Println("having fish...")
}


/*
type Stringer interface {
    String() string
}
*/

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)\n", p.Name, p.Age)
}


func basicsOfInterfaces() {
	fmt.Println("basics of Interfaces")

	// basics()

	// checkEmptyInterface()

	checkStringer()
}

func basics() {
	var animalA Animal = Person{"Sylvia",20}
	animalA.eat()
	// interface can be refered as value or type
	fmt.Printf("animalA v: %v, t: %T\n", animalA, animalA)

	var animalB Animal = &Cat{} // requires func (cat *Cat) eat()
	animalB.eat()

	fmt.Printf("animalB v: %v, t: %T\n", animalB, animalB)
}

func checkStringer() {
	var sylvia = Person{"Sylvia",20}
	fmt.Println(sylvia) // will triggerfunc (p Person) String() string
}

func checkEmptyInterface(){
	//every type implements empty interface
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
