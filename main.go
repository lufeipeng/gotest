package main

import (
	"fmt"
	"github.com/lufeipeng/gotest/mystrings"
)

func main() {
	fmt.Println("main is called \n")
	var a int
	fmt.Printf("a type is %T \n", a)
	fmt.Println(mystrings.MyStringCompare("1", "2"))
}
