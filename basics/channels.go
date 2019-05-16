package main
import (
	"fmt"
	"time"
)

func basicsOfChannels() {
	fmt.Printf("basics of Channels\n")
	// testingSum()
	// testingPutNum()
	// testingChannelWithBuffer1()
	// testingChannelWithBuffer2()
	// testingRangeAndClose()
	// checkSelect()
	checkDefaultSelection()
}

func checkDefaultSelection() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}

func checkSelect() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		// 101 will not be put into quit only until 10 number received from c
		quit <- 101
	}()
	fibonacci(c, quit)
}

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
			fmt.Println("case c")
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func putNumbeAndClose(num int, c chan int) {
	c <- num
	fmt.Printf("put %v to channel\n", num)
	close(c) // not working?
}

// close() does't work?
func testingRangeAndClose() {
	c := make(chan int, 2)

	putNumbeToChannel(1, c)
	go putNumbeToChannel(2, c)

	value, open := <-c
	fmt.Printf("open: %v, value: %v\n", open, value)

	// getAndCheck(c)
	// getAndCheck(c)
}

func getAndCheck(channel chan int) {
	value, notClosed := <-channel
	fmt.Printf("closed: %v, value: %v\n", !notClosed, value)
}

func testingChannelWithBuffer2() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	// ch <- 3
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

func testingChannelWithBuffer1() {
	c := make(chan int, 2)

	go putNumbeToChannel(1, c)
	go putNumbeToChannel(2, c)
	go putNumbeToChannel(3, c)
	go putNumbeToChannel(4, c)
	
	// note: the order is random
	x := <- c
	y := <- c
	z := <- c
	o := <- c
	
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(o)
}

func testingSum() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)

	x := <-c
	y := <-c

	fmt.Println(x)
	fmt.Println(y)
}

func testingPutNum() {
	c := make(chan int)

	go putNumbeToChannel(1, c)
	go putNumbeToChannel(2, c)
	go putNumbeToChannel(3, c)
	go putNumbeToChannel(4, c)
	
	// note: the order is random
	x := <-c
	y := <-c
	z := <-c
	o := <-c
	
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(o)
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func putNumbeToChannel(num int, c chan int) {
	c <- num
	fmt.Printf("put %v to channel\n", num)
}