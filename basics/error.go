package main

import (
	"fmt"
	"time"
)

func bascisOfError() {
	fmt.Println("bascis of error")

	/*
	error type should implements 
	
	type error interface {
    	Error() string
	}
	*/

	anError := DefaultError{"Not working", time.Now()}
	fmt.Println(anError)
}

type DefaultError struct {
	errorMessage string
	errorTime time.Time
}

func (e DefaultError) Error() string {
	return fmt.Sprintf("Error description: %v\nError occurred at: %v\n", e.errorMessage, e.errorTime)
}