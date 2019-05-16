package main
import (
	"fmt"
	"strings"
)

func basicsOfString() {
	fmt.Printf("basics of string\n")
	stringsJoin()
}

func stringsJoin() {
	var list = []string{"a","b","c","d"}
	fmt.Printf("%s\n", strings.Join(list, ":"))
}



