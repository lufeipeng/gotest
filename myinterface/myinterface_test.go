package myinterface

import (
	"fmt"
	"testing"
)

func TestInteface(t *testing.T) {
	sq1 := new(Square)
	sq1.side = 5

	// var areaIntf Shaper
	// areaIntf = sq1
	// shorter,without separate declaration:
	// areaIntf := Shaper(sq1)
	// or even:
	areaIntf := sq1
	fmt.Printf("The square has area: %f\n", areaIntf.Area())

	sq2 := new(Square2)
	sq2.side = 3
	fmt.Printf("The square has area: %f\n", sq2.Area())

	var shape Shaper
	shape = new(Square)
	Show(shape)

	shape = new(Square2)
	shape.Area()
	Show(shape)

	switch t := shape.(type) {
	case *Square2:
		fmt.Printf("type is square2 %v", t)
	case *Square:
		fmt.Printf("type is square %v", t)
	default:
		fmt.Printf("default type")
	}

}
