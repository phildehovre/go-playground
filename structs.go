package main

import (
	"fmt"
	"reflect"
)

type Vertex struct {
	X, Y int
}

var array = [5]string{"hello", "world", "How", "are", "you"}

var zeroed = make([]int, 0, 2)
var zeroedFull = make([]int, 2, 2)

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)

func main() {

	pointer := &array
	slice := pointer[0:3]
	slice[0] = "hi"
	// for _, value := range array {
	// 	fmt.Println(value)
	// }
	printArraySlice(zeroed)
	printArraySlice(zeroedFull)
}

func printArraySlice(array []int) {
	fmt.Printf("Array: cap: %v, len: %v, type: %v,  %v\n", cap(array), len(array), reflect.TypeOf(array), array)
}
