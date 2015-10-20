package myinterface

import (
	"fmt"
)

type Shaper interface {
	Area() float32
}

type Square struct {
	side float32
}

type Square2 struct {
	side int
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

func (sq2 *Square2) Area() float32 {
	return float32(sq2.side * sq2.side)
}

func Show(shaper Shaper) {
	if v, ok := shaper.(*Square); ok {
		fmt.Printf("%T\n", v)
		fmt.Println(v.Area())
	} else {
		fmt.Println("check shaper is not called")
	}

}
