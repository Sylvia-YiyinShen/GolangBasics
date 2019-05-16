package main
import (
	"fmt"
	"runtime"
)


func basicsOfFlowControler() {
	fmt.Printf("basics of flow control\n")
	// checkLoop()
	checkSwitch()
}

func checkSwitch() {
	fmt.Print("Go runs on ")
	os := runtime.GOOS
	switch os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}

func checkLoop() {
	for i := 0; i<10; i++ {
		fmt.Printf("i: %v\n", i)
	}

	sum := 1
	for ; sum < 1000; {
		sum += sum
	}

	// while loop
	s := 1
	for s < 1000 {
		s += s
	}

	// for ever
	for {

	}
}
