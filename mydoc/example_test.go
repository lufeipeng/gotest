package mydoc_test

//fmt lib should not used in example_test.go, because go test will invoke error
import (
    "strconv"
    "fmt"
)

func ExampleSparseFileReader() {
	var a int = 0;
	b := a;
	strconv.Itoa(b)
	if(true){
		fmt.Println("ExampleWriter:");
		fmt.Println("This is the source code repository for the Go programming language.");
		// Output:
		// ExampleWriter:
		// This is the source code repository for the Go programming language.
	}else {
		fmt.Println("ExampleWriter:");
		fmt.Println("This is the source code repository for the Go programming language.");
	}
	// Output:
	// ExampleWriter:
	// This is the source code repository for the Go programming language.	
}

func ExampleMyDocFunction() {
	//fmt.Println("ExampleMyDocFunction")
	// Output:
	// ExampleWriter:
	// This is the source code repository for the Go programming language.
}