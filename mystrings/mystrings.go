package mystrings

import (
	"fmt"
	"strings"
	//"strconv"
)

var MyString string;


func init() {
	fmt.Println("mystring package is init")	
}

// This is My String Compare2.
// str1 and str2 is type string.
// Test For godoc generatge.
func MyStringCompare(str1 string, str2 string) int {
	return strings.Compare(str1, str2);
}