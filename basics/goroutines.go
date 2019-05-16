package main
import (
	"fmt"
	"time"
)

func basicsOfGoroutines() {
	fmt.Printf("basics of Goroutines\n")
	go count("A")
	count("B")
}

func count(id string) {
	fmt.Printf("count")
	for i:= 0; i<10; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("id-%v    %v\n",id,i)
	}
}

