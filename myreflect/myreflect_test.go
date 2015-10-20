package myreflect

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMyReflect(t *testing.T) {
	a := 0
	aType := reflect.TypeOf(a)
	aValue := reflect.ValueOf(a)
	aKind := aValue.Kind()
	fmt.Println("TestMyReflect a type is", aType)
	fmt.Println("TestMyReflect a value is", aValue)
	fmt.Println("TestMyReflect a kind is", aKind)

	b := 0.0
	bType := reflect.TypeOf(b)
	bValue := reflect.ValueOf(b)
	bKind := bValue.Kind()
	fmt.Println("TestMyReflect b type is", bType)
	fmt.Println("TestMyReflect b value is", bValue)
	fmt.Println("TestMyReflect b kind is", bKind)
}
